// main_test.go
package main

import (
	"go-lang-assignments/assignment10/movies/repository"
	"go-lang-assignments/assignment10/movies/service"
	"testing"
)

func TestMovieRepository(t *testing.T) {
	repo := repository.NewInMemoryMovieRepository()

	movie1 := repository.Movie{Name: "Movie ABC", Year: 2021}
	movie2 := repository.Movie{Name: "Movie XYZ", Year: 2022}

	repo.SaveMovie(movie1)
	repo.SaveMovie(movie2)

	if len(repo.GetMoviesByYear(2021)) != 1 {
		t.Errorf("Expected 1 movie for year 2021, got %d", len(repo.GetMoviesByYear(2021)))
	}

	if len(repo.GetMoviesByYear(2022)) != 1 {
		t.Errorf("Expected 1 movie for year 2022, got %d", len(repo.GetMoviesByYear(2022)))
	}
}

func TestMovieService(t *testing.T) {
	repo := repository.NewInMemoryMovieRepository()
	service := service.NewMovieService(repo)

	err := service.ReleaseMovie("Movie A", 2021)
	if err != nil {
		t.Errorf("Error releasing movie: %v", err)
	}

	movies := service.GetMoviesReleasedInYear(2021)
	if len(movies) != 1 || movies[0].Name != "Movie A" {
		t.Errorf("Expected Movie A for year 2021, got %v", movies)
	}
}
