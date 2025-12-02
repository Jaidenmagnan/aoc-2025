package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mod(a int) int {
	return ((a % 100) + 100) % 100
}

func a(s []string) int {
	current := 50
	res := 0
	for _, v := range s {
		previous := current
		if strings.Contains(v, "L") {
			change, err := strconv.Atoi((strings.Replace(v, "L", "", -1)))
			check(err)

			current -= change
		} else if strings.Contains(v, "R") {
			change, err := strconv.Atoi((strings.Replace(v, "R", "", -1)))
			check(err)

			current += change
		}
		adding := 0
		if current < 0 {

			adding = 1 + (-1*current)/100

			if previous == 0 {
				adding -= 1
			}

		} else if mod(current) == 0 {
			if current == 0 {
				adding = 1

			} else {
				adding = current / 100
			}

		} else if (current > 0) && mod(current) != current {
			adding = current / 100
		}

		res += adding

		current = mod(current)

	}
	return res
}

func main() {
	path := "q1.input.txt"
	file, err := os.Open(path)

	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	fmt.Println(a(text))

	file.Close()
}
