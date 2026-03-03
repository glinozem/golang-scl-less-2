package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	printer := func(s []int) {
		fmt.Println("Slice s:", s)
	}

	printer(s)

	s[0] = 0

	printer(s)
}
