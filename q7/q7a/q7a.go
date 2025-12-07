package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGrid(lines [][]string) {
	for _, row := range lines {
		fmt.Println(row)
	}
}

func numberOfSplits(lines [][]string) int {
	splits := 0
	for i:= 0; i < len(lines); i++ {

		for j:= 0; j < len(lines[0]); j++ {

			if lines[i][j] == "S" {
				translateDown := i+1
				for (translateDown < len(lines) && lines[translateDown][j] == ".") {
					lines[translateDown][j] = "|"
					translateDown += 1
				}
			}

			if lines[i][j] == "^" && i != 0 && lines[i-1][j] == "|" {
				splits += 1

				// split left
				translateDown := i + 1
				left := j - 1

				for (left >= 0 && translateDown < len(lines) && lines[translateDown][left] == ".") {
					lines[translateDown][left] = "|"
					translateDown += 1
				}

				// split right
				translateDown = i + 1
				right := j + 1
				for (right < len(lines[0]) && translateDown < len(lines) && lines[translateDown][right] == ".") {
					lines[translateDown][right] = "|"
					translateDown += 1
				}
			}
		}
	}

	printGrid(lines)

	return splits
}

func main() {

	path := "../q7.input.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		lines = append(lines, splitLine)
	}

	result := numberOfSplits(lines)

	fmt.Println(result)
}