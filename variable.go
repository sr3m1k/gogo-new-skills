package main

import "fmt"

func main() {
	var name string
	var age uint8
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Hello " + name + "!")
	fmt.Println(name + ", " + "how old are you?")
	fmt.Scan(&age)
	fmt.Println("you are " + fmt.Sprint(age) + " years old")

}
