package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	sourceName      = "source.txt"
	destinationName = "destination.txt"
	message         = "This is the content of the source file"
)

func main() {

	sourceFile, err := os.Create(sourceName)

	if err != nil {
		log.Fatalf("Error creating source file: %v", err)
	}

	defer sourceFile.Close()

	_, err = sourceFile.WriteString(message)

	if err != nil {
		log.Fatalf("Error writing to source file: %v", err)
	}

	fmt.Println("Source file created successfully...")

	if err != nil {
		log.Fatalf("Error closing source file: %v", err)
	}

	fmt.Println("Source file written successfully...")

	src, err := os.Open(sourceName)

	if err != nil {
		log.Fatalf("Error opening source file for reading: %v", err)
	}

	defer src.Close()

	dest, err := os.Create(destinationName)

	if err != nil {
		log.Fatalf("Error creating destination file: %v", err)
	}

	defer dest.Close()

	destLength, err := io.Copy(dest, src)

	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}

	log.Printf("Copied %d bytes from %s to %s\n", destLength, sourceName, destinationName)
}
