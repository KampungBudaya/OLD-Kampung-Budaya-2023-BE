package app

import (
	"log"
	"net/http"
	"os"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/database"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := database.StartMySQLConn()
	if err != nil {
		return err
	}

	port := os.Getenv("APP_PORT")

	app := mux.NewRouter()

	app.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	api := app.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(&w, http.StatusOK, "I'm fine and healthy! nice to see you :)")
	}).Methods(http.MethodGet)

	http.ListenAndServe(":"+port, app)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	return nil
}
