package models

import (
	"bytes"
	"database/sql"
	"os"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	_ "github.com/lib/pq"
)

func createTestDB(t *testing.T) *embeddedpostgres.EmbeddedPostgres {
	logger := &bytes.Buffer{}
	database := embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Username("web_test").
			Password("web_test").
			Database("chronopaste_test").
			Logger(logger).
			Port(9876),
	)

	return database
}

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=9876 user=web_test password=web_test dbname=chronopaste_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	return db
}
