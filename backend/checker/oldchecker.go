package checker

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"mime/multipart"
// 	"net/http"
// 	"strings"
// 	"time"
// )

// // type Counter struct {
// // 	LinksUp   []string
// // 	LinksDown []string
// // 	Up        int
// // 	Down      int
// // }

// type CheckerResponse struct {
// 	StatusCode  int      `json:"status_code"`
// 	FileName    string   `json:"file_name"`
// 	LinksUp     []string `json:"links_up"`
// 	LinksDown   []string `json:"links_down"`
// 	CounterUp   int      `json:"counter_up"`
// 	CounterDown int      `json:"counter_down"`
// 	ProcessTime string   `json:"process_time"`
// }

// type storer interface {
// 	Upload(interface{}) error
// }

// type Context interface {
// 	FormFile(string) (*multipart.FileHeader, error)
// 	SaveUploadedFile(*multipart.FileHeader, string) error
// 	Bind(interface{}) error
// 	JSON(int, interface{})
// }

// type Handler struct {
// 	store storer
// }

// func NewCheckerHandler() *Handler {
// 	return &Handler{}
// }

// func (h *Handler) Upload2(c Context) {
// 	start := time.Now()

// 	count := &Counter{}

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
// 		return
// 	}

// 	if fileName := ValidateFile(file, c); fileName != "" {
// 		links := ReadFile("saved/" + fileName)
// 		ch := make(chan string)

// 		for _, link := range links {
// 			go CheckLink(link, count, ch)
// 		}

// 		time.Sleep(time.Second)
// 	}

// 	responseCtx := &CheckerResponse{
// 		StatusCode:  http.StatusOK,
// 		FileName:    file.Filename,
// 		LinksUp:     count.LinksUp,
// 		LinksDown:   count.LinksDown,
// 		CounterUp:   count.Up,
// 		CounterDown: count.Down,
// 		ProcessTime: fmt.Sprintf("(Used %.f minute and %.f second)", time.Since(start).Minutes(), time.Since(start).Seconds()),
// 	}

// 	c.JSON(http.StatusOK, responseCtx)
// }

// func ValidateFile(file *multipart.FileHeader, c Context) string {
// 	if hasCsv := strings.HasSuffix(file.Filename, ".csv"); !hasCsv {
// 		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid file type. Please upload a CSV file."})
// 		return ""
// 	}

// 	err := c.SaveUploadedFile(file, "saved/"+file.Filename)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
// 		return ""
// 	}

// 	return file.Filename
// }

// func ReadFile(fileName string) []string {
// 	contents, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return []string{}
// 	}

// 	links := strings.Split(string(contents), ",")
// 	for i := range links {
// 		links[i] = strings.TrimSpace(links[i])
// 	}
// 	return links
// }

// func CheckLink(link string, count *Counter, ch chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		if link != "" {
// 			fmt.Errorf("%s is down and got error %s", link, err)

// 			count.LinksDown = append(count.LinksDown, link)
// 			count.Down++
// 			ch <- link
// 		}
// 		return
// 	}
// 	if link != "" {
// 		count.LinksUp = append(count.LinksUp, link)
// 		count.Up++
// 		ch <- link
// 	}
// }
