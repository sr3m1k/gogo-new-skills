package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's out of work week?")
	today := time.Now().Weekday()
	switch time.Friday {

	case today + 0:
		fmt.Println("Today you may go to bar")

	case today + 1:
		fmt.Println("Tomorrow you can go for a walk")

	case today + 2:
		fmt.Println("just a little more")

	case today + 3:
		fmt.Println("it's only tuesday, what a vacation")

	case today + 4:
		fmt.Println("work the fuck up and don't think about rest :)")
	}
}
