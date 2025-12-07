package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func trashCompactor(lines [][]string) int {
	fmt.Println(lines)
	numbers := make([]int, 0)
	operation := ""
	res := 0


	for columnNum := len(lines[0]) - 1; columnNum >= 0; columnNum-- {
		columnString := make([]string, 0)

		for _, row := range lines {
			columnString = append(columnString, row[columnNum])
		}

		if columnString[len(columnString)-1] == "*" || columnString[len(columnString)-1] == "+" {
			operation = columnString[len(columnString)-1]
			columnString = columnString[:len(columnString)-1]
		}

		columnStringFormatted := strings.Trim(strings.Join(columnString, ""), " ")
		columnInt, err := strconv.Atoi(columnStringFormatted)
		if err != nil {
			continue
		}

		numbers = append(numbers, columnInt)

		if (operation != "") {
			if operation == "*" {
				product := 1
				for _, num := range numbers {
					product *= num
				}
				res += product
			}
			if operation == "+" {
				sum := 0
				for _, num := range numbers {
					sum += num
				}
				res += sum
			}
			numbers = make([]int, 0)
			operation = ""
		}
	}
	return res
}

func main() {
	path := "../q6.input.txt"
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(trashCompactor(lines))
}