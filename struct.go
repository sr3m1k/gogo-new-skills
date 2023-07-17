package main

import "fmt"

type user struct {
	name     string
	age      int64
	password string
}

func change(u *user) {
	var name1 string
	var age1 int64
	var password1 string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name1)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age1)

	fmt.Print("Enter your pass: ")
	fmt.Scan(&password1)

	u.name = name1
	u.age = age1
	u.password = password1

}

func main() {
	var name string
	var age int64
	var password string

	user := user{name, age, password}

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your pass: ")
	fmt.Scan(&password)

	user.name = name
	user.age = age
	user.password = password

	fmt.Println(user)

	change(&user)

	fmt.Println(user)
}
