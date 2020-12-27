package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var helpMsg = `
Usage: CLI Template [OPTIONS]
Options:
	-row	       		Display number or rows.  
	-words             Display number or words.            
	-char               Display number or characters.	
	-help         		Print all instructions		
    `

// All Options in CLI
func showHelp() {
	fmt.Println(helpMsg)
}
func setFlage(flag *flag.FlagSet) {
	flag.Usage = func() {
		showHelp()
	}
}
func help() {

	fmt.Printf(helpMsg)
}

func main() {
	file, err := os.Open("gehad.txt")
	if err != nil {
		fmt.Println("Err ", err)
	}
	// Read the content of the file
	scanner := bufio.NewScanner(file)
	lines, words, characters := 0, 0, 0
	// itrate each line in the file
	// coun the number of lines, words, and characters
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}
	var sHelp bool
	flag.BoolVar(&sHelp, "help", false, "")

	var row bool
	flag.BoolVar(&row, "row", false, "")

	var word bool
	flag.BoolVar(&word, "word", false, "")

	var char bool
	flag.BoolVar(&char, "char", false, "")

	flag.Parse()

	setFlage(flag.CommandLine)

	showHelp()

	if row {
		fmt.Println("\nnumber or rows: ", lines)
	}
	if word {
		fmt.Println("\nnumber or words: ", words)
	}
	if char {
		fmt.Println("\nnumber or characters: ", characters)

	}

}
