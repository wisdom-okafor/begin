package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/beginner/restfulapi/crud/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
