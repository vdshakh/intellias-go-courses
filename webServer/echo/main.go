package main

import (
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

const numberForRandom = 1

func main() {
	e := echo.New()

	e.GET("/random", GetRandomNumberHandler)
	e.POST("/random", func(c echo.Context) error {
		return c.String(http.StatusMethodNotAllowed, "message: wrong http method")
	})

	err := e.Start(":8080")
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func GetRandomNumberHandler(r echo.Context) error {
	minParameter := r.QueryParam("min")
	maxParameter := r.QueryParam("max")

	min, err := strconv.Atoi(minParameter)
	if err != nil {
		return r.String(http.StatusBadRequest,
			fmt.Sprintf("Can't convert min parameter to int"))
	}

	max, err := strconv.Atoi(maxParameter)
	if err != nil {
		return r.String(http.StatusBadRequest,
			fmt.Sprintf("Can't convert max parameter to int"))
	}

	result := rand.Intn(max-min+numberForRandom) + min

	return r.String(http.StatusOK, fmt.Sprintf("Your random number between %v and %v: %v", min, max, result))
}
