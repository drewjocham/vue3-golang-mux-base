package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Test  TestModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Test:  TestModel{DB: db},
		Users: UserModel{DB: db},
	}
}
