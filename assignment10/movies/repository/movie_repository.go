package repository

type Movie struct {
	Name string
	Year int
}

type MovieRepository interface {
	SaveMovie(movie Movie) error
	GetMoviesByYear(year int) []Movie
}

type SaveMovieRepository struct {
	movies map[int][]Movie
}

func NewInMemoryMovieRepository() *SaveMovieRepository {
	return &SaveMovieRepository{
		movies: make(map[int][]Movie),
	}
}

func (repo *SaveMovieRepository) SaveMovie(movie Movie) error {
	repo.movies[movie.Year] = append(repo.movies[movie.Year], movie)
	return nil
}

func (repo *SaveMovieRepository) GetMoviesByYear(year int) []Movie {
	return repo.movies[year]
}
