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
	
	

	f.WriteString("# Testing some output - \n")
	f.WriteString("## sdga \n")
	f.WriteString("| a | b | c |\n")
	f.WriteString("| - | - | - |\n")
	f.WriteString("| a | b | c |\n")
		  

	if err != nil {
		log.Println(err)
	}

}
