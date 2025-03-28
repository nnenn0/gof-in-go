package facade_test

import (
	"testing"

	. "gof-in-go/structural/facade"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	testData := "bob@example.com=Bob\ncharlie@example.com=Charlie"
	db := NewDatabase(testData)

	properties, err := db.GetProperties()

	assert.NoError(t, err)
	assert.Contains(t, properties, "bob@example.com")
	assert.Equal(t, "Bob", properties["bob@example.com"])
	assert.Contains(t, properties, "charlie@example.com")
	assert.Equal(t, "Charlie", properties["charlie@example.com"])
}

func TestGetProperties(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput map[string]string
	}{
		{
			name:  "Valid properties",
			input: "bob@example.com=Bob\ncharlie@example.com=Charlie",
			expectedOutput: map[string]string{
				"bob@example.com":     "Bob",
				"charlie@example.com": "Charlie",
			},
		},
		{
			name:           "Empty input",
			input:          "",
			expectedOutput: map[string]string{},
		},
		{
			name:           "Input with empty lines",
			input:          "\n\nbob@example.com=Bob\n\n",
			expectedOutput: map[string]string{"bob@example.com": "Bob"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := NewDatabase(tc.input)
			properties, err := db.GetProperties()

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, properties)
		})
	}
}

func TestHtmlWriter(t *testing.T) {
	writer := &HtmlWriter{}

	writer.Title("Test Title")
	writer.Paragraph("Test Paragraph")
	writer.Link("https://example.com", "Example Link")
	writer.Mailto("alice@example.com", "Alice")

	html := writer.Close()

	assert.Contains(t, html, "<!DOCTYPE html>")
	assert.Contains(t, html, "<title>Test Title</title>")
	assert.Contains(t, html, "<h1>Test Title</h1>")
	assert.Contains(t, html, "<p>Test Paragraph</p>")
	assert.Contains(t, html, "<a href=\"https://example.com\">Example Link</a>")
	assert.Contains(t, html, "<a href=\"mailto:alice@example.com\">Alice</a>")
	assert.Contains(t, html, "</html>")
}

func TestMakeWelcomePage(t *testing.T) {
	testData := "bob@example.com=Bob\ncharlie@example.com=Charlie"
	db := NewDatabase(testData)
	pm := NewPageMaker(db)

	html, err := pm.MakeWelcomePage("bob@example.com")

	assert.NoError(t, err)
	assert.Contains(t, html, "<title>Bob's web page</title>")
	assert.Contains(t, html, "<p>Welcome to Bob's web page!</p>")
	assert.Contains(t, html, "<a href=\"mailto:bob@example.com\">Bob</a>")
}

func TestMakeWelcomePage_InvalidEmail(t *testing.T) {
	testData := "bob@example.com=Bob\ncharlie@example.com=Charlie"
	db := NewDatabase(testData)
	pm := NewPageMaker(db)

	_, err := pm.MakeWelcomePage("unknown@example.com")

	assert.Error(t, err)
	assert.EqualError(t, err, "email address not found")
}
