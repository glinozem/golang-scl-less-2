package main

import "fmt"

var str string

func main() {
	var i int

	closure := func(i *int, str *string) {
		*(i)++
		*str += "Hello"
	}

	closure(&i, &str)
	closure(&i, &str)
	fmt.Println(str, i)
}
