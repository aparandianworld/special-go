package main

import (
	"fmt"
	"strings"
	"log"
)

func main() {
	fmt.Println("module: c3")
	fmt.Println("version: 0.0.1")

	sentence := "I can't do this. I'm not good at this. This is hard."
	negativeWords := []string{"can't", "not", "hard"}

	result := improveSentence(sentence, negativeWords)

	fmt.Println(result)
}

func improveSentence(sentence string, negativeWords []string) string {

	if len(sentence) == 0 || len(negativeWords) == 0 {
		return sentence
	}

	result := sentence

	for _, word := range negativeWords {

		asterisks := strings.Repeat("*", len(word))
		result = strings.Replace(result, word, asterisks, -1)
	}

	log.Println("DEBUG improveSentence result: ", result)

	return result
}
