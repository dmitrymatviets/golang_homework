package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCounts := make(map[string]int)

	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCounts[strings.ToLower(scanner.Text())]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	for key, val := range wordCounts {
		fmt.Println(key, val)
	}

}
