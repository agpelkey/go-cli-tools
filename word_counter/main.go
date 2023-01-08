package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)


func main() {

	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Defining a boolean flag -b to count bytes instead of words or lines
	bites := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided
	flag.Parse()


	// calling the count function to count the number of words (or lines) received
	// from the standard input and print it out
	fmt.Println(count(os.Stdin, *lines, *bites))
}

func count (r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files) 
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if !countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// for every work scanned, incremenet the counter
	for scanner.Scan() {
		wc ++
	}

	// return the total
	return wc
}
