package checker

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestScanFileNotFound(t *testing.T) {
	// Arrange
	mockFileName := "test.csv"

	// Act
	_, err := scanFile(mockFileName)

	//Assert
	assert.NotNil(t, err)
}

func TestCheckWebsites(t *testing.T) {
	// Arrange
	mockCount := &Counter{}
	websites := []string{
		"http://google.com",
		"https://www.wikipedia.org",
		"https://www.instag777ram.com",
	}

	// Act
	actualResults := CheckWebsites(websites, mockCount)

	//Assert
	want := len(websites)
	got := len(actualResults)
	assert.Equal(t, want, got, "they length should be equal")
}
