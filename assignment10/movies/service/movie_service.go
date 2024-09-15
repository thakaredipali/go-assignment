package service

import (
	"go-lang-assignments/assignment10/movies/repository"
)

type MovieService struct {
	repository repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{repository: repo}
}

func (service *MovieService) ReleaseMovie(name string, year int) error {
	movie := repository.Movie{Name: name, Year: year}
	return service.repository.SaveMovie(movie)
}

func (service *MovieService) GetMoviesReleasedInYear(year int) []repository.Movie {
	return service.repository.GetMoviesByYear(year)
}
