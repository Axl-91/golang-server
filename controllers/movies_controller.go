package controllers

import (
	"encoding/json"
	"go-server/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(writer http.ResponseWriter, request *http.Request) {
	movies := structs.GetMovies()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func GetMovie(writer http.ResponseWriter, request *http.Request) {
	movies := structs.GetMovies()

	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func CreateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie structs.Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)

	structs.CreateMovie(movie)

	movies := structs.GetMovies()
	json.NewEncoder(writer).Encode(movies)
}

func UpdateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie structs.Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)

	structs.UpdateMovie(movie.ID, movie)

	movies := structs.GetMovies()
	json.NewEncoder(writer).Encode(movies)
}

func DeleteMovie(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	structs.RemoveMovie(params["id"])

	movies := structs.GetMovies()
	json.NewEncoder(writer).Encode(movies)
}
