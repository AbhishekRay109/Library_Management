package main

type UserService interface {
	CreateUser(*User) error
	//GetUser(*string) (*User, error)
	GetAllByName(*string) ([]*User, error)
	GetAllByRange(*string) ([]*User, error)
	GetAllByCat(*string) ([]*User, error)
	UpdateIssue(*User) error
	UpdateReturn(*User) error
}

// List of all the Actions(Functions)
