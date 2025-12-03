package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidId(id string) bool {
	length := len(id)

	// this is for part 1
	//if (length % 2) != 0 {
	//	return false
	//}

	for i := 0; i < len(id)-1; i++ {
		for j:= i + 1; j < length; j++ {

			//if (id[0: length / 2] == id[length / 2: length]) {
			if (strings.Repeat(id[i:j], length/(j-i)) == id) && (length%(j-i) == 0) {
				return false
			}
		}	
	}
	return true
}

func a(ids []string) int {
	invalidCount := 0

	for _, idRange := range ids {
		start, end := strings.Split(idRange, "-")[0], strings.Split(idRange, "-")[1]	

		startInt, err := strconv.Atoi(start)
		if err != nil {
			panic(err)
		}

		endInt, err := strconv.Atoi(end)
		if err != nil {
			panic(err)
		}

		for i:= startInt; i <= endInt; i++ {
			if !isValidId(fmt.Sprintf("%d", i)) {
				fmt.Println(i)
				invalidCount += i
			}
		}
	}
	return invalidCount
}

func main() {
	fmt.Println("Hello, World!")

	path := "q2.example.txt"
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	data := strings.Split(text[0], ",")
	fmt.Println(a(data))
	fmt.Println(data)
	file.Close()
}