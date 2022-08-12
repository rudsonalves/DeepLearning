package main

func scrape(url string, selector string, ch chan []string) {
	scraper := Scraper.NewScraper(url)
	selection := scraper.Find(selector)
	ch <- selection
}

func main() {
	years := []string{"2009", "2010", "2011", "2012", "2013"}
	channels := []chan []string{
		make(chan []string),
		make(chan []string),
		make(chan []string),
		make(chan []string),
		make(chan []string),
		make(chan []string),
	}

	for idx, year := range years {
		ch := channels[idx]
		go scrape("http://en.wikipedia.org/wiki/List_of_Bollywood_films_of_"+year,
			"table.wikitable i a", ch)
	}

	for i := 0; i < 5; i++ {
		select {
		case movies2009 := <-channels[0]:
			printMovies(movies2009)
		case movies2010 := <-channels[1]:
			printMovies(movies2010)
		case movies2011 := <-channels[2]:
			printMovies(movies2011)
		case movies2012 := <-channels[3]:
			printMovies(movies2012)
		case movies2013 := <-channels[4]:
			printMovies(movies2013)
		}
	}
}
