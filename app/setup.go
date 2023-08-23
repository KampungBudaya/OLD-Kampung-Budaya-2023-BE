package app

import (
	"log"
	"net/http"
	"os"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/database"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := database.StartMySQLConn()
	if err != nil {
		return err
	}

	app := mux.NewRouter()
	api := app.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(&w, http.StatusOK, "I'm fine and healthy! nice to see you :)")
	})

	http.Handle("/", app)

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), app)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	return nil
}
