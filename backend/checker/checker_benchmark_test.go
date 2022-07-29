package checker

import (
	"testing"
)

func BenchmarkCheckWebsites(b *testing.B) {
	mockCount := &Counter{}
	urls := make([]string, 1000)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(urls, mockCount)
	}
}
