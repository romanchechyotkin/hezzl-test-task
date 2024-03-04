package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	base             *http.Server
	router           *gin.Engine
	pgClient         *pgxpool.Pool
	redisClient      *redis.Client
	clickhouseClient *sql.DB
}

func New(pg *pgxpool.Pool, rc *redis.Client, cc *sql.DB) *Server {
	srv := &Server{
		router:           gin.Default(),
		pgClient:         pg,
		redisClient:      rc,
		clickhouseClient: cc,
	}

	srv.router.GET("/status", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "status")
	})

	srv.base = &http.Server{
		Addr:    ":8080",
		Handler: srv.router,
	}

	return srv
}

func (s *Server) Run() {
	log.Fatal(s.base.ListenAndServe())
}
