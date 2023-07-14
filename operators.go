package main

import "fmt"

func main() {
	var task int
	fmt.Println("Write how many tasks you have")
	fmt.Scan(&task)
	if task > 2 {
		fmt.Println("Ohhh, you have to stay at work")
	} else if task <= 2 {
		fmt.Println("You may go home")
	}
}
