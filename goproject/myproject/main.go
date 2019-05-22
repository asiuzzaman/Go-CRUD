package main

import (
	"fmt"
	"goproject/myproject/driver"
	"net/http"
	"os"

	ph "goproject/myproject/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbName := "products"
	dbPass := "docker"
	dbHost := "localhost"
	dbPort := "5432"
	dbUserName := "postgres"

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUserName, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(pHandler))
	})

	err = connection.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection Successfully")

	fmt.Println("Server Started...")
	http.ListenAndServe(":9000", r)
}

func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
