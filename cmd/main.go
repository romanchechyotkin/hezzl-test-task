package main

import "github.com/romanchechyotkin/hezzl-test-task/pkg/postgresql"

func main() {
	cfg := postgresql.Config {
		User: "postgres",
		Password: "5432",
		Host: "localhost",
		Port: "5432",
		Database: "hezzl",
	}

	pgClient := postgresql.New(&cfg)	
	_ = pgClient
}