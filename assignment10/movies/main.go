// main.go
package main

import (
	"fmt"
	"go-lang-assignments/assignment10/movies/repository"
	"go-lang-assignments/assignment10/movies/service"
)

func main() {
	repo := repository.NewInMemoryMovieRepository()
	service := service.NewMovieService(repo)

	err := service.ReleaseMovie("Movie A", 2021)
	if err != nil {
		fmt.Println("Error releasing movie:", err)
	} else {
		fmt.Println("Movie released successfully")
	}

	movies := service.GetMoviesReleasedInYear(2021)
	fmt.Println("Movies released in 2021:", movies)
}
