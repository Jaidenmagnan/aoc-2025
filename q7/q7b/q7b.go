package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numberOfSplits(lines [][]string) int {
	splits := 0
	for i:= 0; i < len(lines); i++ {

		for j:= 0; j < len(lines[0]); j++ {

			if lines[i][j] == "S" {
				translateDown := i+1
				for (translateDown < len(lines) && lines[translateDown][j] == ".") {
					lines[translateDown][j] = fmt.Sprintf("%d", 1)
					translateDown += 1
				}
			}

			if lines[i][j] == "^" && i != 0 && (lines[i-1][j] != "." && lines[i-1][j] != "S" && lines[i-1][j] != "^") {
				splits += 1
				numberOfWays, err := strconv.Atoi(lines[i-1][j])
				if err != nil {
					panic(err)
				}

				// split left
				translateDown := i
				left := j - 1

				for (left >= 0 && translateDown < len(lines) && lines[translateDown][left] != "^") {
					leftWays := numberOfWays
					if (lines[translateDown][left] != "." && lines[translateDown][left] != "S" && lines[translateDown][left] != "^") {
						existingWays, err := strconv.Atoi(lines[translateDown][left])
						if err != nil {
							panic(err)
						}
						leftWays = existingWays + numberOfWays
					}
					lines[translateDown][left] = fmt.Sprintf("%d", leftWays)
					translateDown += 1
				}

				// split right
				translateDown = i
				right := j + 1
				for (right < len(lines[0]) && translateDown < len(lines) && lines[translateDown][right] != "^") {
					rightWays := numberOfWays
					if (lines[translateDown][right] != "." && lines[translateDown][right] != "S" && lines[translateDown][right] != "^") {
						existingWays, err := strconv.Atoi(lines[translateDown][right])
						if err != nil {
							panic(err)
						}
						rightWays = existingWays + numberOfWays
					}
			
					lines[translateDown][right] = fmt.Sprintf("%d", rightWays)
					translateDown += 1
				}
			}
		}
	}

	res := 0
	for _, valStr := range lines[len(lines)-1] {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			continue
		}
		res += val
	}

	return res
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