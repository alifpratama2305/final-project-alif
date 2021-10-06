package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/final_project")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return db, nil
}
