package main

import (
	"log"
	"net/http"
	"os"

	utility "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/utils"
	handlers "github.com/NoirAbhinav/rss-aggregator/internal/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("SERVICE_PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	get_db, err := utility.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	api_router := handlers.ApiConfig{
		DBPointer: get_db,
	}
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.ReadinessHandler)
	v1Router.Get("/err", handlers.ErrHandler)
	v1Router.Get("/user", api_router.UserCreateHandler)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server listening on port %s\n", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
