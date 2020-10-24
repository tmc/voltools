// Command volunteerd is the server for the volunteerd project.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/HF-RapidResponse/geotools/tigerline/models"
	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	flagVerbose = flag.Bool("v", false, "be vergose")
	flagZipCode = flag.String("zip", "9411%", "zip code query")
)

func main() {
	flag.Parse()

	if err := run(Config{
		Verbose:     *flagVerbose,
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Server struct {
	config Config

	db *sql.DB
}

// Config describes the configuration for the server.
type Config struct {
	Verbose     bool
	DatabaseURL string
}

func NewServer(ctx context.Context, config Config) (*Server, error) {
	connStr, err := pq.ParseURL(config.DatabaseURL)
	if err != nil {
		return nil, err
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
	server, err := NewServer(ctx, cfg)

	zipCodeCount, err := models.ZipCodes().Count(ctx, server.db)
	if err != nil {
		return err
	}

	fmt.Println("zip codes:", zipCodeCount)
	fmt.Println("zip code query:", *flagZipCode)
	zipCodes, err := models.ZipCodes(
		qm.Where(`zcta5ce10 like ?`, *flagZipCode),
	).All(ctx, server.db)
	if err != nil {
		return err
	}
	for _, zipCode := range zipCodes {
		fmt.Println("zip code:", zipCode)
	}
	return nil
}
