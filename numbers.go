package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var your_num int
	fmt.Println("Let`s start the game ;) ")
	fmt.Println("Enter up to what number you have made a wish! And I will be guess your number :)")
	fmt.Scan(&your_num)
	fmt.Println("Your num is: " + fmt.Sprint(rand.Intn(your_num)))

}
