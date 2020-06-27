package users

import (
	postgres "ecormmerce-app/ecormmerce-rest-api/pkg/storage/postgres"
	"os"
	"testing"

	"github.com/go-pg/pg/v9"
)

var dbTest *pg.DB
var userRepositoryTest Repository

func init() {
	dbTest, _ = postgres.Connect(os.Getenv("DB_NAME"))
	userRepositoryTest = NewRepository(dbTest)
}
func TestAddUser(t *testing.T) {
	var user User
	user.Firstname = "a"
	user.Lastname = "b"
	user.Password = "c"
	user.Gender = "d"
	user.Username = "e"

	status := userRepositoryTest.AddUser(&user)
	if !status {
		t.Errorf("Test Failed; Users was not added")
	}
}

func TestGetAllusers(t *testing.T) {

	users, err := userRepositoryTest.GetAllUsers()
	if err != nil {
		t.Errorf("Test Failed; No users found")
	}
	if len(users) != 1 {
		t.Errorf("Test Failed; No users found")
	}
}
