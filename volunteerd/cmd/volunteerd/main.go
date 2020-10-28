// Command volunteerd is the server for the volunteerd project.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/Masterminds/sprig"
	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var (
	flagVerbose = flag.Bool("v", false, "be vergose")
	flagZipCode = flag.String("zip", "94110", "zip code query")
	flagListen  = flag.Bool("server", true, "if true, enables server mode")
)

func main() {
	flag.Parse()

	if err := run(Config{
		Verbose:      *flagVerbose,
		Listen:       *flagListen,
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://localhost/db"),
		TemplateRoot: getEnv("VOLUNTEERD_TEMPLATE_ROOT", "./volunteer-ui/dist"),
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getEnv(envVar, defaultValue string) string {
	if s := os.Getenv(envVar); s != "" {
		return s
	}
	return defaultValue
}

// Server implements the server for volunteerd.
type Server struct {
	// holds configuration for the server.
	config Config

	// assetRegexp identifies patterns that should be served directly, without template interpolation.
	assetRegexp *regexp.Regexp

	db *sql.DB // connection to a SQL database.

	mu           sync.RWMutex // protects the following
	cachedAppCSS template.CSS
}

// Config describes the configuration for the server.
type Config struct {
	Verbose bool
	// If true, the process should run an http server.
	Listen      bool
	DatabaseURL string

	// A/the pattern to pass files through without template interpolation.
	AssetRegexp  string
	TemplateRoot string
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	connStr, err := pq.ParseURL(config.DatabaseURL)
	fmt.Println("connstr:", connStr)
	if err != nil {
		return nil, fmt.Errorf("issue parsing DATABASE_URL: %w", err)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if config.AssetRegexp == "" {
		config.AssetRegexp = `\.(css|js|svg)$`
	}
	re, err := regexp.Compile(config.AssetRegexp)
	return &Server{
		config:      config,
		db:          db,
		assetRegexp: re,
	}, err
}

func run(cfg Config) error {
	ctx := context.Background()
	if *flagVerbose {
		boil.DebugMode = true
		ctx = boil.WithDebug(ctx, true)
	}
	server, err := NewServer(ctx, cfg)
	if err != nil {
		return err
	}
	if cfg.Listen {
		return server.ListenAndServe()
	} else {
		// test call
		return server.zipCodeTestCalls(ctx)
	}
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleRoot)
	mux.HandleFunc("/index.css", s.handleCSS)
	mux.HandleFunc("/zip/", s.handleZip)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Println("listening on", port)
	return http.ListenAndServe(":"+port, mux)
}

func (s *Server) getTemplateFromDisk(templateName string) string {
	// TODO: add (aggressive) caching.
	c, err := ioutil.ReadFile(filepath.Join(s.config.TemplateRoot, templateName))
	if err != nil {
		return fmt.Sprintf(`issue finding template: %v`, err)
	}
	return string(c)
}

func (s *Server) renderTemplate(w http.ResponseWriter, templateName string, ctx interface{}) error {
	contents := s.getTemplateFromDisk(templateName)
	if s.assetRegexp.MatchString(templateName) {
		// TODO: implement a more efficient/caching mechanism. Also consider file embedding.
		fmt.Fprint(w, contents)
		return nil
	}
	// TODO: Parse on startup (and HUP?), not every time.
	t, err := template.New(templateName).Funcs(sprig.FuncMap()).Parse(contents)
	if err != nil {
		return err
	}
	return t.Execute(w, ctx)
}

func (s *Server) templateContext(r *http.Request) interface{} {
	appHTML, appCSS, err := renderReactApp(filepath.Join(s.config.TemplateRoot, "parcel-manifest.json"), "index.js", "renderServerSide()")

	s.mu.Lock()
	s.cachedAppCSS = appCSS
	s.mu.Unlock()
	return struct {
		Env             string
		Request         *http.Request
		IsPrerendered   string
		PrerenderedHTML template.HTML
		PrerenderedCSS  template.CSS
		Error           error
	}{
		Env:             "dev",
		Request:         r,
		PrerenderedHTML: appHTML,
		PrerenderedCSS:  appCSS,
		Error:           err,
	}
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "index.html"
	}
	// don't parse as template
	if err := s.renderTemplate(w, path, s.templateContext(r)); err != nil {
		fmt.Fprintln(w, err)
	}
}

func (s *Server) handleZip(w http.ResponseWriter, r *http.Request) {
	zip := r.URL.Path[len("/zip/"):]
	fmt.Println("zip", zip)
	fmt.Fprintln(w, "zip:", zip)
	zi, err := s.lookupZipInfo(r.Context(), zip)
	if err != nil {
		panic(err)
	}
	s.renderTemplate(w, "zip_detail.html", zi)
}

func (s *Server) handleCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Fprintln(w, s.cachedAppCSS)
}
