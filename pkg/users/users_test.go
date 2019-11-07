package users

import "testing"



func TestAddUser(t *testing.T) {
	var user User
	user.Firstname = "a"
	user.Lastname = "b"
	user.Password = "c"
	user.Gender = "d"
	user.Username = "e"

	var userServiceTest Service
    err := userServiceTest.AddUser(user)
    if err != nil {
       t.Errorf("Test Failed; Users was not added")
    }
}