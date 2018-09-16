package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	var positive float32
	var negative float32
	var zero float32
	for _, elem := range arr {
		if elem > 0 {
			positive++
		} else if elem < 0 {
			negative++
		} else {
			zero++
		}
	}

	positive = positive / float32(len(arr))
	negative = negative / float32(len(arr))
	zero = zero / float32(len(arr))

	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(zero)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(arr)
}
