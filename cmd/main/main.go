package main

import (
	"fmt"
	"log"
	"webScrapping/internal/movies"
)

func main()  {
	topMovies, err := movies.GetTopMovies()

	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range topMovies {
		fmt.Println(fmt.Sprintf("%s. %s", movie.Ranking, movie.Title))
	}
}