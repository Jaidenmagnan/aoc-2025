package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)


func mergeIntervals(intervals map[int]int, starts []int) map[int]int {
	if len(intervals) == 0 {
		return intervals
	}

	merged := make(map[int]int)

	slices.Sort(starts)
	fmt.Println(starts)

	currentStart := starts[0]
	currentEnd := intervals[currentStart]

	for i := 1; i < len(starts); i++ {
		start := starts[i]
		end := intervals[start]

		if start <= currentEnd {
			currentEnd = max(currentEnd, end)
		} else {
			merged[currentStart] = currentEnd
			currentStart = start
			currentEnd = end
		}
	}

	merged[currentStart] = currentEnd

	return merged
}

func countAvailableIds(idRanges []string) int {
	res := 0

	m := make(map[int]int)
	starts := make([]int, 0)

	// first we create a map of starts
	for _, idRange := range idRanges {
		var start, end int
		fmt.Sscanf(idRange, "%d-%d", &start, &end)

		val, ok := m[start]

		if (ok) {
			end = max(end, val)
		} else {
			starts = append(starts, start)
		}

		m[start] = end
	}

	// now we merge
	m = mergeIntervals(m, starts)

	for start := range m {
		end := m[start]
		res += end - start + 1
	}

	return res
}

func main() {
	path := "../q5.input.txt"
	file, err := os.Open(path)
	
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	
	var idRanges []string
	for scanner.Scan() {
		if (scanner.Text() == "") {
			break
		}
		idRanges = append(idRanges, scanner.Text())
	}

	fmt.Println(countAvailableIds(idRanges))
}