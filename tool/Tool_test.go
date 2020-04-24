package tool

import (
	"strings"
	"testing"
)

func TestUrls(t *testing.T) {
	tool := new(AdjustTool)
	urls := []string{"http://google.com"}
	results := tool.Run(urls, false, 0)
	for i, result := range results {
		if result.url != urls[i] {
			t.Error("Doesn't contain the url")
		}
		if result.hash == "" {
			t.Error("Doesn't contain hash")
		}
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
func TestUrlsParallel(t *testing.T) {
	tool := new(AdjustTool)
	urls := []string{"http://google.com","http://yahoo.com", "http://adjust.com"}
	results := tool.Run(urls, true, 3)
	for _, result := range results {
		if !contains(urls, result.url) {
			t.Error("Doesn't contain the url",result.url)
		}
		if result.hash == "" {
			t.Error("Doesn't contain hash")
		}
	}
}

func TestUrlsWithoutHttp(t *testing.T) {
	tool := new(AdjustTool)
	urls := []string{"google.com"}
	results := tool.Run(urls, false, 0)
	for i, result := range results {
		if result.url != "http://" + urls[i] {
			t.Error("Doesn't contain the url")
		}
		if result.hash == "" {
			t.Error("Doesn't contain hash")
		}
	}
}

func TestUrlsWithoutHttpParallel(t *testing.T) {
	tool := new(AdjustTool)
	urls := []string{"google.com","yahoo.com", "adjust.com"}
	results := tool.Run(urls, true, 3)
	for _, result := range results {
		if !contains(urls, strings.Replace(result.url, "http://", "", 1)) {
			t.Error("Doesn't contain the url", result.url)
		}
		if result.hash == "" {
			t.Error("Doesn't contain hash")
		}
	}
}