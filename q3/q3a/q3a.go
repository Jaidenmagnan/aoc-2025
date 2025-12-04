package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxBattery(lines []string) int {
	res := 0
	for _, line := range lines {

		leftMaxValue := -1
		leftMaxIndex := -1

		rightMax := 0

		for i, character := range line {
			length := len(line)

			if i < length - 1 {
				if int(character) > leftMaxValue {
					leftMaxValue = int(character)
					leftMaxIndex = i
				}
			}
		}

		for i, character := range line {

			if i > leftMaxIndex {
				if int(character) >= rightMax {
					rightMax = int(character)
				}
			}
		}

		battery, err  := strconv.Atoi(string(leftMaxValue) + (string(rightMax)))
		if err != nil {
			panic(err)
		}

		fmt.Println(battery)
		res += battery
	}
	return res
}

func main() {
	fmt.Println("Hello, World!")

	path := "q3.input.txt"
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(maxBattery(lines))


}

