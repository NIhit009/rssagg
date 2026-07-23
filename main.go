package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NIhit009/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("Port could not be found in the env file")
	}

	Db_url := os.Getenv("DB_URL")
	if Db_url == ""{
		log.Fatal("Database url could not be found in the env file")
	}

	conn, err := sql.Open("postgres", Db_url)
	if err != nil{
		log.Fatal("Error while connecting to the database..", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	fmt.Printf("The Server is running on port %v \n", port)

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
	v1Router.Post("/users", apiCfg.handlerUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.HandlerGetUserByApiKey))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	v1Router.Post("/feedFollow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollow))
	v1Router.Get("/feedFollows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Delete("/feedFollow/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	err = server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}