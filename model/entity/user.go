package entity

type User struct {
	Operation      string
	Id             string
	PersonalNumber string
	Password       string
	Email          string
	Name           string
	RoleID         string
	Role           Role
	IsActive       int32
}
