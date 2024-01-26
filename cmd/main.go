package main

import (
	"developer/api"
	"developer/api/handler"
	"developer/config"
	"developer/storage/postgres"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}
	fmt.Println("success")
	
	defer store.Close()

	handler := handler.New(store)

	api.New(handler)

	fmt.Println("Server is running on port 8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("error while running server err:", err.Error())
	}
}
