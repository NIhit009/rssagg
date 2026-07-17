package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("Port could not be found in the env file")
	}
	fmt.Printf("The Server is running on port %v", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)
	router.Mount("/v1", v1Router);

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}	