package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Print(len(slice), slice)

}
