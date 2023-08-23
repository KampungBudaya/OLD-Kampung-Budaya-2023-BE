package app

import (
	"log"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/database"
	"github.com/joho/godotenv"
)

func Run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	db, err := database.StartMySQLConn()
	if err != nil {
		return err
	}

	// ...

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	return nil
}
