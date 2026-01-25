package main

import (
	"github.com/fatih/color"
	"github.com/glinozem/golang-scl-less-2/modules/internal/usecase"
	"github.com/google/uuid"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hello := usecase.NewHello()
	color.Green(hello.Say())
	color.Cyan("UUID: %s", uuid.New().String())

	//fmt.Println("UUID:", uuid.New().String())
}
