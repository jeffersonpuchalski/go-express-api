package utils

import (
	"encoding/json"
	"net/http"
)

type ApiRatings struct {
	Source string `json:Source`
	Value  string `json:Value`
}
type ApiMovie struct {
	Title      string       `json:Title`
	Year       string       `json:Year`
	Rated      string       `json:Rated`
	Released   string       `json:Released`
	Runtime    string       `json:Runtime`
	Genre      string       `json:Genre`
	Director   string       `json:Director`
	Writer     string       `json:Writer`
	Actors     string       `json:Actors`
	Plot       string       `json:Plot`
	Language   string       `json:Language`
	Country    string       `json:Country`
	Awards     string       `json:Awards`
	Poster     string       `json:Poster`
	Ratings    []ApiRatings `json:Ratings`
	Metascore  string       `json:Metascore`
	imdbRating string       `json:imdbRating`
	imdbVotes  string       `json:imdbVotes`
}

func get(w http.ResponseWriter, r *http.Request) {
	movie := ApiMovie{

		Title:    "Guardians of the Galaxy Vol. 2",
		Year:     "2017",
		Rated:    "PG-13",
		Released: "05 May 2017",
		Runtime:  "136 min",
		Genre:    "Action, Adventure, Comedy, Sci-Fi",
		Director: "James Gunn",
		Writer:   "James Gunn, Dan Abnett (based on the Marvel comics by), Andy Lanning (based on the Marvel comics by), Steve Englehart (Star-Lord created by), Steve Gan (Star-Lord created by), Jim Starlin (Gamora and Drax created by), Stan Lee (Groot created by), Larry Lieber (Groot created by), Jack Kirby (Groot created by), Bill Mantlo (Rocket Raccoon created by), Keith Giffen (Rocket Raccoon created by), Steve Gerber (Howard the Duck created by), Val Mayerik (Howard the Duck created by)",
		Actors:   "Chris Pratt, Zoe Saldana, Dave Bautista, Vin Diesel",
		Plot:     "The Guardians struggle to keep together as a team while dealing with their personal family issues, notably Star-Lord's encounter with his father the ambitious celestial being Ego.",
		Language: "English",
		Country:  "USA",
		Awards:   "Nominated for 1 Oscar. Another 15 wins & 57 nominations.",
		Poster:   "https://m.media-amazon.com/images/M/MV5BNjM0NTc0NzItM2FlYS00YzEwLWE0YmUtNTA2ZWIzODc2OTgxXkEyXkFqcGdeQXVyNTgwNzIyNzg@._V1_SX300.jpg",
		Ratings: []ApiRatings{
			{
				Source: "Internet Movie Database",
				Value:  "7.6/10",
			},
			{
				Source: "Rotten Tomatoes",
				Value:  "85%",
			},
			{
				Source: "Metacritic",
				Value:  "67/100",
			},
		},
		Metascore:  "67",
		imdbRating: "7.6",
		imdbVotes:  "562,796"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(movie)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
