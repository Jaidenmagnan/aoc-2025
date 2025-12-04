package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func accessibleForklifts(graph [][]string) int {
	accessibleForklifts := 0

	dy := []int{-1, 1, 0, 0, 1, -1, 1, -1}
	dx := []int{0, 0, -1, 1, 1, -1, -1, 1}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j] == "@" {
				adjacentForklifts := 0

				for k := 0; k < 8; k++ {
					ny := i + dy[k]
					nx := j + dx[k]

					if ny >= 0 && ny < len(graph) && nx >= 0 && nx < len(graph[0]) {
						if graph[ny][nx] == "@" {
							adjacentForklifts++
						}
					}
				}

				if adjacentForklifts < 4 {
					accessibleForklifts++
				}
			}
		}
	}

	return accessibleForklifts
}

func main() {
	path := "../q4.input.txt"
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var graph [][]string
	for _, line := range lines {
		graph = append(graph, strings.Split(line, ""))
	}

	fmt.Println(accessibleForklifts(graph))
}