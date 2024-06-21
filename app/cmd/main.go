package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jirlon/digitalbank/app/api"
)

func main() {

	ctx := context.Background()
	connStr := "postgresql://postgres:postgres@localhost:5432/digital_bank_db"

	pool, err := pgxpool.New(ctx, connStr)

	if err != nil {
		panic(err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}
	r := api.NewRouter(pool)

	log.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
