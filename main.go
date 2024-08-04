package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		fmt.Println("Port is not defined")
		portString = "8080" 
	}

	router := chi.NewRouter()

	// CORS options configuration
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "https://*"}, // list of allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},      // list of allowed HTTP methods
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // list of allowed headers
		ExposedHeaders:   []string{"Link"},                              // list of headers exposed to the browser
		AllowCredentials: false,                                          // whether to allow credentials
		MaxAge:           300,                                           // maximum value for Access-Control-Max-Age header in seconds
	}

	// Apply CORS middleware to router
	router.Use(cors.Handler(corsOptions))

	v1Router := chi.NewRouter()


	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server starting on port:", portString)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
