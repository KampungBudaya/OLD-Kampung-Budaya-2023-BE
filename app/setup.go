package app

import (
	"fmt"
	"log"
	"net/http"

	"os"

	fordaHandler "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/delivery/http"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/forda/usecase"
	googleHttpDelivery "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/delivery/http"
	googleUsecase "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	oauthRepository "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/repository"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/config"
	_ "github.com/KampungBudaya/Kampung-Budaya-2023-BE/docs"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := config.StartMySQLConn()
	if err != nil {
		return err
	}

	googleConf := config.ConfigureGoogleOAuth()

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
		response.Success(w, http.StatusOK, "I'm fine and healthy! nice to see you :)")
	}).Methods(http.MethodGet)

	oauthRepo := oauthRepository.NewOAuthRepository(db)
	googleUsecase := googleUsecase.NewGoogleUsecase(oauthRepo)
	googleHttpDelivery.NewGoogleHandler(v1, googleUsecase, googleConf)

	fordaRepository := repository.NewFordaRepository(db)
	fordaUsecase := usecase.NewFordaUsecase(fordaRepository)
	fordaHandler.NewFordaHandler(v1, fordaUsecase)

	fmt.Println("Server running on port " + port)
	if err := http.ListenAndServe(":"+port, app); err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	return nil
}
