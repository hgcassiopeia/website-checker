package checker

import (
	"bufio"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type storer interface {
	Upload(interface{}) error
}

type Context interface {
	FormFile(string) (*multipart.FileHeader, error)
	SaveUploadedFile(*multipart.FileHeader, string) error
	Bind(interface{}) error
	JSON(int, interface{})
}

type Handler struct {
	store storer
}

func NewCheckerHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Upload(c Context) {
	start := time.Now()

	count := &Counter{}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if hasCsv := strings.HasSuffix(file.Filename, ".csv"); !hasCsv {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid file type. Please upload a CSV file."})
		return
	}

	err = c.SaveUploadedFile(file, "saved/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
		return
	}

	f, err := os.Open("saved/" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	links := scanFile(f)
	var sitesStatus map[string]bool
	sitesStatus = CheckWebsites(links, count)

	responseCtx := &CheckerResponse{
		StatusCode:   http.StatusOK,
		FileName:     file.Filename,
		SitesStatus:  sitesStatus,
		CounterUp:    count.Up,
		CounterDown:  count.Down,
		TotalCounter: len(sitesStatus),
		ProcessTime:  time.Since(start).Milliseconds(),
	}
	c.JSON(http.StatusOK, responseCtx)
}

func scanFile(f *os.File) []string {
	s := bufio.NewScanner(f)
	links := []string{}

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		link := lineArray[0]
		links = append(links, link)
	}
	return links
}

type result struct {
	string
	bool
}

func CheckWebsites(urls []string, count *Counter) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, CheckAvailableStatus(u, count)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

func CheckAvailableStatus(link string, count *Counter) bool {
	if link == "" {
		return false
	}

	_, err := http.Get(link)
	if err != nil {
		if link != "" {
			fmt.Errorf("%s is down and got error %s", link, err)

			count.LinksDown = append(count.LinksDown, link)
			count.Down++
		}
		return false
	}
	count.LinksUp = append(count.LinksUp, link)
	count.Up++
	return true
}
