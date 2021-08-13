package mysql

import (
	"fmt"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	record := &RecordRow{"1", 0, 0, time.Now()}

	err = db.Insert(record)
	if err != nil {
		t.Error(err)
	}
}

func TestAddNewPerson(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	person := &EmployeeRow{"1", "abc",0}

	err = db.Add(person)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchExistingEmployeeID(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	name, err := db.SearchEmployeeID("1")

	if err != nil {
		t.Error(err)
	}

	if name != "abc" {
		t.Error("")
	}
}

func TestSearchNonexistentEmployee(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	name, err := db.SearchEmployeeID("2")

	fmt.Println(name, err)

	if name != "" {
		t.Error("")
	}
}

func TestSearchExistingModelID(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	name, err := db.SearchModelID(0)

	if err != nil {
		t.Error(err)
	}

	if name != "abc" {
		t.Error("wrong name")
	}
}

func TestSearchNonexistentModelID(t *testing.T) {
	dbConfig := &DBConfig{
		Username: "larva",
		Password: "AAbb@1234",
		Host:     "192.168.8.178",
		Port:     3306,
		Database: "cicada",
	}

	db, err := CreateDBSession(dbConfig)
	if err != nil {
		panic(err)
	}

	name, err := db.SearchModelID(1)

	if err != nil {
		t.Error(err)
	}

	if name != "" {
		t.Error("wrong name")
	}
}

func TestTimeNow(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
