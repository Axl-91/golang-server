package structs

import (
	"math/rand"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func InitMovies() {
	movies = append(movies, Movie{ID: "1", Isbn: "52345", Title: "Last Action Hero", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "51231", Title: "Predator", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "3", Isbn: "12314", Title: "Action Man", Director: &Director{Firstname: "John", Lastname: "Doe"}})
}

func GetMovies() []Movie {
	return movies
}

func CreateMovie(new_movie Movie) {
	if new_movie.ID == "" {
		new_movie.ID = strconv.Itoa(rand.Intn(1000000))
	}
	movies = append(movies, new_movie)
}

func UpdateMovie(id string, updated_values Movie) {
	RemoveMovie(id)
	CreateMovie(updated_values)
}

func RemoveMovie(id string) {
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
