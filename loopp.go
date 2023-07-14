package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	var x float64
	fmt.Println("Write your digit")
	fmt.Scan(&x)
	fmt.Println(sqrt(x))
}
