package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main()  {
	maxMovies := 10
	// Crear un nuevo collector
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	// Definir lo que queremos hacer al encontrar los elementos deseados
	c.OnHTML(".ipc-metadata-list-summary-item", func(e *colly.HTMLElement)  {
		if maxMovies > 0 {
			ranking := strings.Split(e.ChildText(".meter-const-ranking"), " ")[0]
			title := e.ChildText(".ipc-title__text")
			fmt.Println(fmt.Sprintf("%s. %s", ranking, title))
			maxMovies--
		} 
	})

	// Manejar errores en caso de que ocurra alguno
	c.OnError(func(_ *colly.Response, err error)  {
		log.Println("Algo salió mal:", err)
	})

	// Visitar la página objetivo
	err := c.Visit("https://www.imdb.com/chart/moviemeter/")
	if err != nil {
		log.Fatal((err))
	}
}