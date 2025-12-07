package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func trashCompactor(operations []string, worksheet [][]int) int {
	fmt.Println(worksheet)
	res := 0
	for i, operation := range operations {
		total := 0
		for _, row := range worksheet {
			val := row[i]

			switch operation {
				case "*":
					if (total == 0) {
						total = 1
					}
					total *= val
				case "+":
					total += val
			}
		}

		fmt.Println(total)
		res += total
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

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	operationsRaw := strings.Split(lines[len(lines)-1], " ")
	operations := make([]string, 0)

	for _, op := range operationsRaw {
		if (op != "") {
			operations = append(operations, op)
		}
	}

	lines = lines[:len(lines)-1]

	worksheet := make([][]int, 0)

	for _, line := range lines {
		worksheetLine := strings.Split(line, " ")
		intList := make([]int, 0)

		for _, number := range worksheetLine {
			intVal, err := strconv.Atoi(strings.TrimSpace(number))

			if (err != nil) {
				continue
			}

			intList = append(intList, intVal)
		}
		worksheet = append(worksheet, intList)
	}

	fmt.Println(trashCompactor(operations, worksheet))
}