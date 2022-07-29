package checker

import (
	"bufio"
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

	err = c.SaveUploadedFile(file, os.Getenv("PATH_UPLOAD")+"/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
		return
	}

	links, err := scanFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
		return
	}

	var sitesStatus []WebsiteStatus
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

func scanFile(Filename string) ([]string, error) {
	f, err := os.Open(os.Getenv("PATH_UPLOAD") + "/" + Filename)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)
	links := []string{}

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		link := lineArray[0]
		links = append(links, link)
	}
	return links, nil
}

func CheckWebsites(urls []string, count *Counter) []WebsiteStatus {
	results := []WebsiteStatus{}
	resultChannel := make(chan WebsiteStatus)

	for index, url := range urls {
		go func(u string, i int) {
			i++
			resultChannel <- WebsiteStatus{
				ID:     i,
				Link:   u,
				Status: CheckAvailableStatus(u, count),
			}
		}(url, index)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		r := WebsiteStatus{
			ID:     result.ID,
			Link:   result.Link,
			Status: result.Status,
		}
		results = append(results, r)
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
			count.LinksDown = append(count.LinksDown, link)
			count.Down++
		}
		return false
	}
	count.LinksUp = append(count.LinksUp, link)
	count.Up++
	return true
}
