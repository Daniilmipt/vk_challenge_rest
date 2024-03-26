package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbInit() (*sql.DB, error) {
	err := godotenv.Load("conv.env")
	if err != nil {
		return nil, errors.New("error loading .env file: " + err.Error())
	}

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=localhost", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		return nil, errors.New("can not open: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("error while check ping db: " + err.Error())
	}
	return db, nil
}
