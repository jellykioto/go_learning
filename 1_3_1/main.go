package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Print(sequenceAverage(n) + sequenceAverage(m))
}

func sequenceAverage(end int) float64 {
	var sum float64 = 0.0
	for i := 1.0; i <= float64(end); i++ {
		sum += i
	}
	return sum / float64(end)
}
