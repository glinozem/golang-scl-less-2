package main

import (
	"github.com/glinozem/golang-scl-less-2/modules/internal/usecase"
	"github.com/glinozem/golang-scl-less-2/modules/pkg/logger"
	"github.com/google/uuid"
	_ "go.uber.org/automaxprocs"
)

func main() {
	log := logger.New("app")
	log.SetPrefix("app")

	hello := usecase.NewHello()
	log.Info("%s", hello.Say())

	log.Info("UUID: %s", uuid.New().String())
}
