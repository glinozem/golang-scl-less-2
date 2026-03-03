package main

import "fmt"

func process(temp *int32) {
	var value2 int32 = 200

	*temp = value2
}

func main() {
	var value2 int32 = 100
	pointer := &value2

	fmt.Println("Value: ", *pointer) // 100
	fmt.Println("Pointer: ", pointer)

	process(pointer)

	fmt.Println("Value: ", *pointer) // 200
	fmt.Println("Pointer: ", pointer)
}
