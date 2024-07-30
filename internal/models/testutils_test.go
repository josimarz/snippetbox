package models

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func newTestDB(t *testing.T) *sql.DB {
	ctx := context.Background()
	mysqlContainer, err := mysql.Run(ctx,
		"mysql:5.7",
		mysql.WithDatabase("snippetbox"),
		mysql.WithUsername("admin"),
		mysql.WithPassword("vQ1kN7uJ1nN4dL2kG1aN3bE9eQ6vM4nH"),
		mysql.WithScripts(filepath.Join("testdata", "setup.sql")),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	connStr, err := mysqlContainer.ConnectionString(ctx, "parseTime=true", "multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		defer db.Close()
		if err := mysqlContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	})
	return db
}
