package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/m-d-nabeel/exploding-kittens/config"
	"github.com/m-d-nabeel/exploding-kittens/handler"
	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://:@localhost:6379/0")
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)
	apiCfg := config.ApiConfig{ // Use correct casing for ApiConfig
		DB: rdb,
	}

	// Create an apiConfigHandler
	apiHandler := handler.NewApiConfigHandler(&apiCfg)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	userRouter := chi.NewRouter()
	gameRouter := chi.NewRouter()
	router.Mount("/v1", v1Router)

	// API ROUTES
	v1Router.Mount("/user", userRouter)
	v1Router.Mount("/game", gameRouter)
	v1Router.Get("/healthz", apiHandler.HandlerReadiness)
	v1Router.Get("/data", apiHandler.GetAllData)

	// USER ROUTES
	userRouter.Post("/create", apiHandler.HandlerCreateUser)
	userRouter.Get("/details", apiHandler.MiddlewareAuth(apiHandler.GetUserDetails))

	// GAME ROUTES
	gameRouter.Get("/details", apiHandler.MiddlewareAuth(apiHandler.GetAllGameDetails))

	// START SERVER
	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}
	log.Printf("Server is running on port %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
