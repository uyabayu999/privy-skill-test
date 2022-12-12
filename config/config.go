package config

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@tcp(db:3306)/privy_test")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
	}
	s.String, s.Valid = string(data), true
	return nil
}
