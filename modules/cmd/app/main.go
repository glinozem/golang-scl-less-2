package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/glinozem/golang-scl-less-2/modules/internal/usecase"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hello := usecase.NewHello()
	fmt.Println(hello.Say())

	fmt.Println("UUID:", uuid.New().String())
}
