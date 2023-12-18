package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/FMinister/chrono_paste/internal/models"
	_ "github.com/lib/pq"
)

type application struct {
	logger  *slog.Logger
	chronos *models.ChronoModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "postgres://<username>:<password>@<url>:<port>/<db-name>?sslmode=disable", "PostgreSql data source name")
	flag.Parse()

	// logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:  logger,
		chronos: &models.ChronoModel{DB: db},
	}

	logger.Info("starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Check that the connection is working correctly.
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Logging to a file with flag
// go run ./cmd/web >>/tmp/web.log
// Note: Using the double arrow >> will append to an existing file, instead of truncating it when starting the application.
