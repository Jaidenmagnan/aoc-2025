package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func search(l []int, target int) int {
	if len(l) == 0 {
		return -1
	}

	left, right := 0, len(l)-1
	closest := l[left]

	for left <= right {
		mid := left + (right-left)/2

		if l[mid] == target {
			return l[mid]

		} else if l[mid] < target {
			closest = l[mid]
			left = mid + 1

		} else {
			right = mid - 1
		}
	}

	return closest
}

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

func countAvailableIngredients(idRanges []string, ingredientIds []string) int {
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
	fmt.Println(m)

	starts = make([]int, 0)
	for start := range m {
		starts = append(starts, start)
	}
	slices.Sort(starts)

	// now we find available ingredients
	for _, ingredientIdStr := range ingredientIds {
		ingredientIdInt, err := strconv.Atoi(ingredientIdStr)
		if err != nil {
			panic(err)
		}

		closestStart := search(starts, ingredientIdInt)
		fmt.Println(ingredientIdInt, ":", closestStart)

		if ingredientIdInt >= closestStart && ingredientIdInt <= m[closestStart] {
			fmt.Println("true")
			res++
		}
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
	
	var ingredientIds []string
	for scanner.Scan() {
		ingredientIds = append(ingredientIds, scanner.Text())
	}

	fmt.Println(countAvailableIngredients(idRanges, ingredientIds))
}