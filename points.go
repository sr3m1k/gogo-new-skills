package main

import "fmt"

func change(str *string) {
	*str = "LOL"
}

func main() {
	num := 9
	b := &num
	*b = 35
	fmt.Println(*b)
}
