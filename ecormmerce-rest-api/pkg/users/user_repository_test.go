package users

import (
	postgres "ecormmerce-app/ecormmerce-rest-api/pkg/storage/postgres"
	"os"
	"testing"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
)

var dbTest *pg.DB
var userRepositoryTest Repository

func setupTestCase(t *testing.T, db *pg.DB) func(t *testing.T) {
	t.Log("setup test case")
	err := db.Insert(&UserRole{
		ID:          1,
		RoleName:    "admin",
		Description: "admin",
	})
	if err != nil {
		t.Errorf("Test Failed; Could not insert user role seed data: \n" + err.Error())
	}

	err = db.Insert(&User{
		Firstname: "test first name",
		Lastname:  "test last name",
		Password:  "test password",
		Gender:    "d",
		Username:  "test.username",
		EmailWork: "test@email.com",
		PhoneWork: "02933482",
		Role:      1,
		Status:    true,
	})

	if err != nil {
		t.Errorf("Test Failed; Could not insert  user seed data: \n" + err.Error())
	}

	return func(t *testing.T) {
		t.Log("teardown test case")
		_, err = db.Model((*User)(nil)).Exec(`TRUNCATE TABLE ?TableName`)
		if err != nil {
			t.Errorf("Test Failed; Users Table truncate failed")
		}

		_, err := db.Model((*UserRole)(nil)).Exec(`TRUNCATE TABLE ?TableName cascade`)
		if err != nil {
			t.Errorf("Test Failed; User Roles Table truncate failed")
		}

	}
}

func pgOptions() pg.Options {
	return pg.Options{
		Addr:            os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASS"),
		Database:        os.Getenv("DB_TEST_NAME"),
		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func pgConnect() (*pg.DB, error) {
	dbInfo := pgOptions()
	return postgres.Connect(dbInfo)

}
func TestAddUser(t *testing.T) {
	t.Log(t.Name())
	err := godotenv.Load()
	if err != nil {
		t.Log(err.Error())
	}

	dbTest, err = pgConnect()
	if err != nil {
		t.Errorf("Test Failed; DB Connection failed")
	}
	defer dbTest.Close()
	userRepositoryTest = NewRepository(dbTest)
	// var user User

	teardownTestCase := setupTestCase(t, dbTest)
	defer teardownTestCase(t)
	user := User{
		Firstname: "a",
		Lastname:  "b",
		Password:  "c",
		Gender:    "d",
		Username:  "w",
		EmailWork: "e",
		PhoneWork: "1",
		Role:      1,
		Status:    true,
	}
	status := userRepositoryTest.AddUser(&user)
	if !status {
		t.Errorf("Test Failed; Users was not added")
	}

}

func TestAddUserWithoutEmail(t *testing.T) {
	t.Log(t.Name())
	err := godotenv.Load()
	if err != nil {
		t.Log(err.Error())
	}

	dbTest, err = pgConnect()
	if err != nil {
		t.Errorf("Test Failed; DB Connection failed")
	}
	defer dbTest.Close()
	userRepositoryTest = NewRepository(dbTest)
	user := User{
		Firstname: "a",
		Lastname:  "b",
		Password:  "c",
		Gender:    "d",
		Username:  "qwerty",
		PhoneWork: "1",
		Role:      1,
		Status:    true,
	}

	status := userRepositoryTest.AddUser(&user)
	if status {
		t.Errorf("Test Failed; Users added. User Added Without Email")
	}

}

func TestAddUserWithoutRole(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Log(err.Error())
	}

	dbTest, err = pgConnect()
	if err != nil {
		t.Errorf("Test Failed; DB Connection failed")
	}
	defer dbTest.Close()
	userRepositoryTest = NewRepository(dbTest)
	user := User{
		Firstname: "a",
		Lastname:  "b",
		Password:  "c",
		Gender:    "d",
		Username:  "asdf",
		EmailWork: "e",
		Status:    true,
	}

	status := userRepositoryTest.AddUser(&user)
	if status {
		t.Errorf("Test Failed; Users added. User Added Without Role")
	}

}

func TestGetAllusers(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		t.Log(err.Error())
	}

	dbTest, err = pgConnect()
	if err != nil {
		t.Errorf("Test Failed; DB Connection failed")
	}
	defer dbTest.Close()
	userRepositoryTest = NewRepository(dbTest)

	teardownTestCase := setupTestCase(t, dbTest)
	defer teardownTestCase(t)

	users, err := userRepositoryTest.GetAllUsers()
	if err != nil {
		t.Errorf("Test Failed; No users found")
	}
	if len(users) < 1 {
		t.Errorf("Test Failed; No users found")
	}
	t.Log(len(users))
}
