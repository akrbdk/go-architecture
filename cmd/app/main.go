package main

import (
	"go-architecture/internal/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Start()
	if err != nil {
		log.Fatal(err)
	}
}
