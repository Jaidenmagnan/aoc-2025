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

		maxes := []int{}
		maxIndexes := []int{}


		for i := 0; i < 12; i++ {
		maxes = append(maxes, -1)
		maxIndexes = append(maxIndexes, -1)

			for j, character := range line {
				length := len(line)
			
				if j < length - (12-i - 1) && (i == 0 || j > maxIndexes[i-1]) {
					if int(character) > maxes[i] {
						maxes[i] = int(character)
						maxIndexes[i] = j
					}
				}
			}
		}

		batteryString := ""
		for _, max := range maxes {
			batteryString += string(max)
		}

		battery, err  := strconv.Atoi(batteryString)
		if err != nil {
			panic(err)
		}

		fmt.Println(battery)
		res += battery
	}
	return res
}

func main() {
	path := "../q3.input.txt"
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

