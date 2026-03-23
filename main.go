package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("/path/to/filename")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	fmt.Println("number of lines:", lineCount)
}
