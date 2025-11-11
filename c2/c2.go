package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("module: c2")
	fmt.Println("version: 0.0.1")

	urls := []string{
		"https://github.com",
		"https://www.google.co",
		"https://www.example.net",
		"https://google.com",
		"https://linkedin.com",
		"https://coderpad.io",
	}

	tlds := []string{".com", ".net", ".io"}

	for _, tld := range tlds {
		count := countURLsWithTLD(urls, tld)
		fmt.Printf("Number of URLs with TLD %s: %d\n", tld, count)
	}
}

func countURLsWithTLD(urls []string, tld string) int {

	if len(urls) == 0 || tld == "" {
		return 0
	}

	count := 0

	for _, url := range urls {
		if strings.HasSuffix(url, tld) {
			count++
		}
	}

	return count
}
