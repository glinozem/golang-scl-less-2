package main

import "fmt"

func main() {
	var i int
	var str string

	closure := func() {
		i++
		str += "Hello"
	}

	closure()
	closure()
	fmt.Println(str, i)
}
