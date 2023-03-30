package db

import "strconv"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func init() {
	movies = []Movie{
		{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}},
		{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}},
		{ID: "3", Isbn: "448745", Title: "Movie Three", Director: &Director{FirstName: "Karen", LastName: "Williams"}},
	}
}

func GetMovies() []Movie {
	return movies
}
func GetMovie(id string) *Movie {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie
		}
	}
	return nil
}
func AddMovie(movie Movie) {
	movies = append(movies, movie)
}

func DeleteMovie(id string) {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func GetNextID() string {
	return strconv.Itoa(len(movies))
}

func UpdateMovie(id string, movie *Movie) {
	for index, m := range movies {
		if m.ID == id {
			(*movie).ID = id
			movies[index] = *movie
			break
		}
	}
}
