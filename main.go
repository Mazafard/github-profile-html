package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	file, err := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// i := r.Perm(5)

	c := colly.NewCollector(
	// colly.AllowedDomains("github.com"),
	)

	c.OnHTML("main", func(h *colly.HTMLElement) {
		e, err := h.DOM.Html()
		fmt.Fprintf(w, h.Text)
		// fmt.Println(h.Text)
		fmt.Println(e, err)
		//h = "This is a Heading This is a paragraph."

	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://github.com/mazafard")
	w.Flush()
}
