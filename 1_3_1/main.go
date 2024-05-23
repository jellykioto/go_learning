package main

import (
	"fmt"
)

func main() {
	var n, num, max int
	_, _ = fmt.Scan(&n, &max)
	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&num)
		if num > max {
			max = num
		}
	}
	fmt.Print(max)
}
