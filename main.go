package main

import (
	"github/go_web-server-with-echo/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func main() {
	// scrapper.Scrape("python")

	// Echo instance
	e := echo.New()
	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}
