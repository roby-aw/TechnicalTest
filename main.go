package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}

func MultiplicationS() string {
	var str string
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			str1 := fmt.Sprintf("%d*%d=%d\t", i, j, i*j)
			str = str + str1
		}
		str = str + "\n"
	}
	return str
}

func main() {
	hasil := MultiplicationS()
	port := "8080"

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, hasil)
	})

	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit

}
