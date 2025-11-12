package main

import (
	"fmt"
	"strings"
	"log"
)

func main() {
	fmt.Println("module: c3")
	fmt.Println("version: 0.0.1")

	sentence := "I can't do this activity due to time and resources."
	negativeWords := []string{"can't", "not"}

	result := improveSentence(sentence, negativeWords)

	fmt.Println(result)
}

func improveSentence(sentence string, negativeWords []string) string {

	if len(sentence) == 0 || len(negativeWords) == 0 {
		return sentence
	}

	negativeWordsString := strings.Join(negativeWords, " ")

	words := strings.Split(sentence, " ")

	log.Println("DEBUG: words:", words)
	log.Println("DEBUG: negativeWordsString:", negativeWordsString)

	for index, word := range words {

		for _, negWord := range negativeWords {

			if strings.EqualFold(word, negWord) {
				words[index] = strings.Repeat("*", len(negWord))
				break
			}

		}
	}

	log.Println("DEBUG: words:", words)

	return strings.Join(words, " ")
}
