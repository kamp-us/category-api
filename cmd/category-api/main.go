package main

import (
	"github.com/kamp-us/category-api/internal/backend"
	"log"
	"net/http"
	"os"

	"github.com/kamp-us/category-api/internal/db"
	"github.com/kamp-us/category-api/internal/models"
	categoryapi "github.com/kamp-us/category-api/rpc/category-api"
	"github.com/kamp-us/category-api/server"
)

func main() {
	dbClient, err := db.NewPostgreSQLConnection(db.PostgreSQLConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	})

	if err != nil {
		log.Fatal("error while creating a db connection pool", err)
	}

	err = models.AutoMigrate(dbClient)
	if err != nil {
		return
	}

	postgreSQLBackend := backend.NewPostgreSQLBackend(dbClient)

	s := server.NewCategoryAPIServer(postgreSQLBackend)
	twirpHandler := categoryapi.NewCategoryAPIServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK"))
	}))

	http.ListenAndServe(":80", mux)
}
