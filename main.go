package main

import (
	"errors"
	"fmt"
	// "net/http"
	// "github.com/hanxuema/gogetstart/controllers"
	 "github.com/hanxuema/gogetstart/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "XAVIER",
		LastName:  "XIE",
	}
	fmt.Println(u)
	// port := 3000
	// _, err := startWebServer(port, 2)
	// fmt.Println(port, err)

	// controllers.RegisterControllers()
	// http.ListenAndServe(":3000", nil)

	loops()
	panicMain()
}

func startWebServer(port, numberOfRetry int) (int, error) {
	fmt.Println("starting server...")
	fmt.Println("server started...", port)
	fmt.Println("number of retry...", numberOfRetry)
	return port, errors.New("something went wrong")
}
