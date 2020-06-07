package controllers

type StorageInterface interface {
	IsUsernameTaken(username string) bool
	CreateUser(username string, password string, profilePic string) error
}
