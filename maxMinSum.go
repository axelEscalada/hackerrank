package main

import (
	"fmt"
	"sort"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	min := arr[0]
	max := arr[len(arr)-1]
	var sum int64
	for _, elem := range arr {
		sum += int64(elem)
	}
	fmt.Print(sum-int64(max), sum-int64(min))

}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
