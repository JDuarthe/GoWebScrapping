package movies

import (
	"log"
	"strings"
	"webScrapping/pkg/models"

	"github.com/gocolly/colly"
)

func GetTopMovies() ([]models.Movie, error) {
	maxMovies := 10
	var movies []models.Movie
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	c.OnHTML(".ipc-metadata-list-summary-item", func(h *colly.HTMLElement) {
		if maxMovies > 0 {
			ranking := strings.Split(h.ChildText(".meter-const-ranking"), " ")[0]
			title := h.ChildText(".ipc-title__text")
			movies = append(movies, models.Movie{Title: title, Ranking: ranking})
			maxMovies--
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something get wrong:", err)
	})

	err := c.Visit("https://www.imdb.com/chart/moviemeter/")
	if err != nil {
		return nil, err
	}
	return movies, nil
}