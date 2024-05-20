package main

import (
	"fmt"
	"go-server/controllers"
	"go-server/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	structs.InitMovies()

	route.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	route.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", route))

}
