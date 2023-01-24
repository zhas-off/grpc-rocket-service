package main

import (
	"log"

	"github.com/zhas-off/grpc-service/internal/database"
	"github.com/zhas-off/grpc-service/internal/rocket"
)

func Run() error {
	rocketStore, err := database.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
