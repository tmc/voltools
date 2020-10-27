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
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		TemplateRoot: "./templates",
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Server implements the server for volunteerd.
type Server struct {
	config Config

	db *sql.DB
}

// Config describes the configuration for the server.
type Config struct {
	Verbose bool
	// If true, the process should run an http server.
	Listen      bool
	DatabaseURL string

	TemplateRoot string
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	connStr, err := pq.ParseURL(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("issue parsing DATABASE_URL: %w", err)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Server{
		config: config,
		db:     db,
	}, nil
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
	// TODO: Parse on startup (and HUP?), not every time.
	t, err := template.New(templateName).Funcs(sprig.FuncMap()).Parse(s.getTemplateFromDisk(templateName))
	if err != nil {
		return err
	}
	return t.Execute(w, ctx)
}

func (s *Server) templateContext(r *http.Request) interface{} {
	app, err := renderReactApp("volunteer-ui/dist/parcel-manifest.json", "index.js", "renderServerSide()")
	return struct {
		Env         string
		Request     *http.Request
		Prerendered string
		Error       error
	}{
		Env:         "dev",
		Request:     r,
		Prerendered: app,
		Error:       err,
	}
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root req")
	path := r.URL.Path
	if path == "/" {
		path = "index.html"
	}
	fmt.Println("rendering", path)
	if err := s.renderTemplate(w, path, s.templateContext(r)); err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Println("done rendering index.html")
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
