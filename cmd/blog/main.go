package main

import (
	"blog/internal/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Println("Start ...")
	log.Fatal(http.ListenAndServe(":8080", router.SetRouter()))
}
