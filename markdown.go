package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("testing.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	f.WriteString("# Testing some output")

	if err != nil {
		log.Println(err)
	}

}
