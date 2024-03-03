package postgresql_test

import (
	"context"
	"testing"

	"github.com/romanchechyotkin/hezzl-test-task/pkg/postgresql"
)

func TestPostgres(t *testing.T) {
	cfg := postgresql.Config{
		User:     "postgres",
		Password: "5432",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
	}

	pgClient := postgresql.New(&cfg)

	t.Run("create project success", func(t *testing.T) {
		err := pgClient.CreateProject(context.Background(), "test 1")
		if err != nil {
			t.Fatal()
		}
	})

}
