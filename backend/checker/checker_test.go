package checker

import (
	"reflect"
	"testing"
)

// func mockWebsiteChecker(url string) bool {
// 	if url == "waat://furhurterwe.geds" {
// 		return false
// 	}
// 	return true
// }

func TestCheckWebsites(t *testing.T) {
	mockCount := &Counter{}
	websites := []string{
		"http://google.com",
		"https://www.wikipedia.org",
		"https://www.instag777ram.com",
	}

	actualResults := CheckWebsites(websites, mockCount)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("Lenght of input and output expected is %v, but got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":            true,
		"https://www.wikipedia.org":    true,
		"https://www.instag777ram.com": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}
