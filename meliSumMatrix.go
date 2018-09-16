package main

import (
	"fmt"
)

func diagonalDifference(arr [][]int32) int32 {
	var sumA int32
	var sumB int32
	sizeArr := len(arr) - 1
	for index := range arr {
		sumA += arr[index][index]
		sumB += arr[index][sizeArr-index]
	}

	return sumA + sumB
}

func main() {

	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12}}

	result := diagonalDifference(arr)

	fmt.Println(result)
}
