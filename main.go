package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/wasay1567/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello World")
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database")
	}

	apiCfg := ApiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Post("/users", apiCfg.handlerUserCreate)
	v1router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1router.Get("/users", apiCfg.handlerGetFeeds)

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port: %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
