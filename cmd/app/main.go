package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server running on port 8080")

	e := echo.New()
	e.GET("/date", Handler)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	curDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	durToNewYear := time.Until(time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC))
	msg := fmt.Sprintf("Current date: %s. \nDays until New Year: %d", curDate, int64(durToNewYear.Hours())/24)
	err := c.String(http.StatusOK, msg)
	if err != nil {
		return err
	}

	return nil
}
