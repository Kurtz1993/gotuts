package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 0

	for x < 10 {
		fmt.Println(x)
		x += 5
	}

	// while-true equivalent
	// for {
	// 	fmt.Println("I'm here forever")
	// }

	y := 5

	for {
		fmt.Println("Do stuff with y", y)

		y += 3

		if y > 25 {
			break
		}
	}

	for y := 5; y < 25; y += 3 {
		fmt.Println("Do more stuff with y", y)
	}
}
