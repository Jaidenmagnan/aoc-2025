package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getAreas(points [][]int, current int) []int {
	areas := make([]int, 0)

	for i, point := range points {
		if i <= current {
			continue
		}

		x2, y2 := points[current][0], points[current][1]

		area := math.Abs(float64(x2-point[0]+1)) * math.Abs(float64(y2-point[1]+1))

		fmt.Println("Point:", point, "Current:", points[current], "Area:", area)

		areas = append(areas, int(area))
	}
	return areas
}

func getLargestArea(points [][]int) int {
	largestArea := 0

	for i := 0; i < len(points); i++ {
		areas := getAreas(points, i)

		for _, area := range areas {
			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

func main() {
	path := "../q9.input.txt"

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	points := make([][]int, 0)
	for scanner.Scan() {
		pointStr := strings.Split(scanner.Text(), ",")

		x, err := strconv.Atoi(pointStr[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(pointStr[1])
		if err != nil {
			panic(err)
		}

		point := []int{x, y}

		points = append(points, point)
	}

	fmt.Println(points)

	fmt.Println("Largest Area:", getLargestArea(points))

}
