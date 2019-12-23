package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User //point to the user array
	nextID = 1
)
