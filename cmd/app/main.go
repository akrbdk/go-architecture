package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running on port 8080")

	e := echo.New()
	e.GET("/data", Handler)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	err := c.String(http.StatusOK, "Hello World")
	if err != nil {
		return err
	}

	return nil
}
