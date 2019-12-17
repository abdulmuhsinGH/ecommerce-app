package users

import (
	"testing"
)

var userServiceTest Service

func TestAddUser(t *testing.T) {
	var user User
	user.Firstname = "a"
	user.Lastname = "b"
	user.Password = "c"
	user.Gender = "d"
	user.Username = "e"

	err := userServiceTest.AddUser(user)
	if err != nil {
		t.Errorf("Test Failed; Users was not added")
	}
}

func TestGetAllusers(t *testing.T) {

	users := userServiceTest.GetAllUsers()
	if len(users) != 1 {
		t.Errorf("Test Failed; No users found")
	}
}
