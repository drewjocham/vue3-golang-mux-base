package data

import (
	"database/sql"
	validator "github.com/interviews/internal/vaildator"
	"time"
)

type Test struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

func ValidateTest(v *validator.Validator, test *Test) {
	v.Check(test.FirstName != "", "firstName", "must be provided")
	v.Check(len(test.FirstName) <= 500, "firstName", "must not be more than 500 bytes long")

	v.Check(test.LastName != "", "lastName", "must be provided")
	v.Check(len(test.FirstName) <= 500, "lastName", "must not be more than 500 bytes long")
}

type TestModel struct {
	DB *sql.DB
}
