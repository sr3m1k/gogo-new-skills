package main

import "fmt"

func add(x int, y int) int {
	return (x + y)
}
func main() {
	var x int
	var y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(add(x, y))

}
