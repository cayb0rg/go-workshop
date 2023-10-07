package models

type Character struct {
	Name  Name
	Class Class
	Race  string
}

type Name struct {
	FirstName string
	LastName  string
}

type Class struct {
	ClassName string
}
