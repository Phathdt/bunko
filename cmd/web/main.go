package main

import (
	"bunko/db/postgres"
	"bunko/web"
	"log"
	"net/http"
)

func main() {
	store, err := postgres.NewStore("postgresql://postgres:123123123@localhost:5432/bunko?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	h := web.NewHandler(store)

	http.ListenAndServe(":3000", h)
}
