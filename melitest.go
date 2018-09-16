package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int32
	arr = append(arr, 6)
	arr = append(arr, 3)
	arr = append(arr, 3)
	arr = append(arr, 1)
	arr = append(arr, 7)

	result := cutTheSticks(arr)
	fmt.Println(result)
}

func cutTheSticks(arr []int32) []int32 {
	var sizes []int32
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
	pos := 0
	remaining := len(arr)
	for 0 < remaining {
		sizes = append(sizes, int32(remaining))
		count := countEqualFrom(arr, pos)
		pos += count
		remaining -= count

	}

	return sizes
}

func countEqualFrom(sticks []int32, from int) int {
	value := sticks[from]
	for i := 1; from+i < len(sticks); i++ {
		if value != sticks[from+i] {
			return i
		}
	}
	return len(sticks) - from
}
