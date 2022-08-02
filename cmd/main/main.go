package main

import (
	"log"
	"net/http"

	"github.com/alireza/bookstore/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	if err := http.ListenAndServe(":9090", router); err != nil {
		log.Fatal(err)
	}
}
