package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findClosestBoxes(lines [][]int, box int) [][]float64 {
	x1, y1, z1 := lines[box][0], lines[box][1], lines[box][2]
	distances := make([][]float64, 0)

	for i, line := range lines {
		if i <= box {
			continue
		}

		x2, y2, z2 := line[0], line[1], line[2]

		dist := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2) + math.Pow(float64(z2-z1), 2))

		distances = append(distances, []float64{float64(box), float64(i), dist})
	}
	return distances
}

func getResult(groups [][]int) bool {
	isCombined := false
	if (len(groups) == 1) {
		return false
	}
	for _, group := range groups {
		if len(group) > 0 {
			if isCombined {
				return false
			}
			isCombined = true
		}
	}
	return true
}

func makeGroups(pairs [][]float64, xs []int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][2] <= pairs[j][2]
	})
	groups := make([][]int, 0)
	assignments := make(map[int]int)
	groupNum := 0

	count := 0

	previous1 := -1
	previous2 := -1

	for _, pair := range pairs {
		box1, box2, _ := int(pair[0]), int(pair[1]), pair[2]


		box1Group, okBox1 := assignments[box1]
		box2Group, okBox2 := assignments[box2]

		previous2 = xs[box2]
		previous1 = xs[box1]

		if !okBox1 && !okBox2 {
			groups = append(groups, []int{box1, box2})
			assignments[box1] = groupNum
			assignments[box2] = groupNum
			groupNum += 1
		} else if okBox1 && !okBox2 {
			groups[box1Group] = append(groups[box1Group], box2)
			assignments[box2] = box1Group
		} else if !okBox1 && okBox2 {
			groups[box2Group] = append(groups[box2Group], box1)
			assignments[box1] = box2Group
		} else {
			if box1Group != box2Group {
				groups[box1Group] = append(groups[box1Group], groups[box2Group]...)
				for _, b := range groups[box2Group] {
					assignments[b] = box1Group
				}
				groups[box2Group] = []int{}
			}
		}

		if (len(assignments) == len(xs)) {
			if (getResult(groups)) {
		 		break
			}
		 }

		count += 1
	}
	fmt.Println(groups)
	fmt.Println("prev1:", previous1, "prev2:", previous2)
	fmt.Println(previous1 * previous2)
	return 0
}

func makeJunctionBoxes(lines [][]int) int {
	pairs := make([][]float64, 0)
	xs := make([]int, 0)

	for i, coords := range lines {
		xs = append(xs, coords[0])
		boxDistances := findClosestBoxes(lines, i)

		pairs = append(pairs, boxDistances...)
	}
	makeGroups(pairs, xs)
	return 0
}

func main() {
	path := "../q8.input.txt"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		intCoords := make([]int, 0)
		for _, coord := range coords {
			num, err := strconv.Atoi(coord)
			if err != nil {
				panic(err)
			}
			intCoords = append(intCoords, num)
		}

		lines = append(lines, intCoords)
	}

	makeJunctionBoxes(lines)

}