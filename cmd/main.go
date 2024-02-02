package main

import (
	"context"
	"log"
	"developer/api"
	"developer/config"
	"developer/storage/postgres"
)

func main() {
	cfg := config.Load()

	pgStore, err := postgres.New(context.Background(), cfg)
	if err != nil {
		log.Fatalln("Error while connecting to db err:", err.Error())
		return
	}
	defer pgStore.Close()

	server := api.New(pgStore)

	err = server.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}