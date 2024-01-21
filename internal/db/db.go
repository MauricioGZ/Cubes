package db

import (
	"database/sql"
	"fmt"

	"github.com/MauricioGZ/Cubes/settings"
	_ "github.com/go-sql-driver/mysql"
)

func New(s settings.Settings) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", s.DB.User, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Name)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5) // only for test can be removed

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
