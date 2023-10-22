package main

import (
	"expvar"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"

	"projector-test-app/app/handlers"
)

func NewRouter(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) http.Handler {
	rootRouter := chi.NewRouter()

	rootRouter.Route("/", func(router chi.Router) {
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		userHandler := handlers.NewUser(mongoClient, elasticClient)
		router.Route("/", func(r chi.Router) {
			r.Get("/users", userHandler.HandleGetUsers)
			r.Post("/users", userHandler.HandleCreateUser)
		})

		router.Method("GET", "/debug/vars", expvar.Handler())
	})

	return rootRouter
}
