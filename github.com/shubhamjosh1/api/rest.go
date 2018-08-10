package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBooks() {

}
func InitiliseRouter() {

	// init Routers
	r := mux.NewRouter()

	//Route Handlers/ Endpoints
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", delBook).Methods("DELETE")

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
