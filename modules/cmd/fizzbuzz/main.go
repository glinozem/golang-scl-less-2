package main

import "fmt"

func FizzBuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%v", num)
	}
}

func main() {
	for i := range 100 {
		fmt.Println(FizzBuzz(i + 1))
	}
}
