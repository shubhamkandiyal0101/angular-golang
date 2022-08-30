package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	router "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(" -- -- Basic Shopping App Golang -- -- ")

	r := router.Router()
	http.ListenAndServe(":4000", r)

}
