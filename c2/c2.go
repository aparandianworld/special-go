package main

import (
	"fmt"
	"regexp"
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

	result1 := countURLsWithTLD(urls, ".com")
	fmt.Println("Result for TLD .com: ", result1)

	result2 := countURLsWithTLD(urls, ".net")
	fmt.Println("Result for TLD .net: ", result2)

	result3 := countURLsWithTLD(urls, ".io")
	fmt.Println("Result for TLD .io: ", result3)
}

func countURLsWithTLD(urls []string, tld string) int {
	count := 0
	pattern := tld + "$"

	re := regexp.MustCompile(pattern)

	for _, url := range urls {
		if re.MatchString(url) {
			count++
		}
	}

	return count
}
