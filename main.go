package main

import (
	"fmt"
	"os"
	"log"
	"net/http"


	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hi :D")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not Found in the Environment")
	}

	router := chi.NewRouter()

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router) // /v1/ready

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	fmt.Println(portString)
}