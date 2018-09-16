package main

import (
	"fmt"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	for i := 0; i <= int(n); i++ {
		spaces := strings.Repeat(" ", int(n)-i)
		sharp := strings.Repeat("#", int(i))
		fmt.Println(spaces + sharp)
	}
}

func main() {
	staircase(6)
}
