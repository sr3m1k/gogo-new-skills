package main

import "fmt"

func main() {
	slice := []int{}
	var pow = []int{13, 67, 23, 98, 43, 17, 3, 21, 63}
	for i := 0; i != len(pow); i++ {
		switch pow[i] % 3 {

		case 0:
			slice = append(slice, pow[i])

		}
	}
	fmt.Println(slice)
}
