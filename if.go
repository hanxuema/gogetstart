package main

import "github.com/hanxuema/gogetstart/models"

func ifMain()  {
	u1 := models.User{
		ID: 1,
		FirstName: "xavier",
		LastName: "xie",
	}

	u2 := models.User{
		ID: 1,
		FirstName: "winston",
		LastName: "xie",
	}

	if u1.ID == u2.ID {
		println("same user")
	}
}