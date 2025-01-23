package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-server/controllers"
	"go-server/structs"
	"log"
	"net/http"
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
