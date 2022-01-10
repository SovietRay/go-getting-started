package models

type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
}

var (
	users []*User //Работаем с указателями, чтобы лишний раз не копировать данные
	nextID = 1 
)