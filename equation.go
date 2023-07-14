package main

import (
	"fmt"
	"math"
)

func main() {
	var a_monomial int
	var b_monomial int
	var c_monomail int
	fmt.Println("Write a")
	fmt.Scan(&a_monomial)
	fmt.Println("Write b")
	fmt.Scan(&b_monomial)
	fmt.Println("Write c")
	fmt.Scan(&c_monomail)
	discriminant := b_monomial*b_monomial - 4*a_monomial*c_monomail
	if discriminant > 0 {
		x1 := ((float64(-b_monomial)) + (math.Sqrt(float64(discriminant)))) / (float64(2 * a_monomial))
		x2 := ((float64(-b_monomial)) - (math.Sqrt(float64(discriminant)))) / (float64(2 * a_monomial))
		fmt.Println(x1, x2)
	} else if discriminant == 0 {
		x := -b_monomial / (2 * a_monomial)
		fmt.Println(x)
	} else {
		fmt.Println("There are no roots, I counted in vain")
	}
}
