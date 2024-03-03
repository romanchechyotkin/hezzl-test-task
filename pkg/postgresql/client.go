package postgresql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	pool *pgxpool.Pool
}

func (c *Client) CreateProject(ctx context.Context, name string) error {
	exec, err := c.pool.Exec(ctx, "insert into projects (name) values ($1)", name)
	if err != nil {
		return err
	}

	log.Println("rows affected", exec.RowsAffected())

	return nil
}

func newClient(dsn string) *Client {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")

	return &Client{
		pool: pool,
	}
}
