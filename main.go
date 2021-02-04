package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"promo/infrastructure/database"
	"promo/infrastructure/promoProductApi"
	"promo/userInterface/controllers"
)

func main() {
	err := godotenv.Load()
	if nil != err {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if nil != err {
		log.Fatal(err.Error())
	}

	controller := controllers.NewProductController(db, promoProductApi.NewPromoProductApi())

	router := mux.NewRouter()
	router.HandleFunc("/create", controller.Create).Methods("POST")
	router.HandleFunc("/index", controller.Index).Methods("GET")

	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token := request.Header.Get("X-token")
			if token != os.Getenv("AUTH_TOKEN") {
				http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			} else {
				handler.ServeHTTP(writer, request)
			}
		})
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
