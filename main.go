package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"strings"
	"slices"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load godotenv")
	}

	var path string = os.Getenv("savePATH")
	var url string = os.Getenv("saveURL")
	var crawl bool = os.Getenv("CRAWL") == "y"
	var history []string

	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	ex, err := os.Getwd()
	if err != nil {
		log.Fatal("Error in getting executable dir")
	}

	if crawl {
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			absoluteURL := e.Request.AbsoluteURL(link)

			if !slices.Contains(history, absoluteURL) {
				history = append(history, absoluteURL)
				c.Visit(absoluteURL)
			}
		})
	}


	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")

		if strings.HasSuffix(strings.ToLower(link), ".gif") {
			var rawPath []string = strings.Split(link, "/")
			var filePath string = ex + path + rawPath[len(rawPath)-1]
			
			fmt.Printf("Image SRC found: %s\n", link)

			if strings.Contains(link, "https://") {
				downloadFile(filePath, link)
			} else {
				downloadFile(filePath, "https://"+url+link)
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) { log.Fatal("ERROR...", err) })
	c.OnResponse(func(r *colly.Response) { fmt.Println("Loading...", r.Request.URL) })
	c.OnRequest(func(r *colly.Request) { fmt.Println("Visiting", r.URL.String()) })

	c.Visit("https://" + url)

	fmt.Println("Collected all gifs to './gifs/'!")
}

func downloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil  {
	  return err
	}
	defer out.Close()
  
	resp, err := http.Get(url)
	if err != nil {
	  return err
	}
	defer resp.Body.Close()
  
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
	  return err
	}
  
	return nil
  }
