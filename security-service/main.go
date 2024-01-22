package main

import (
	"log"
	"net/http"
	"subproblem/security-service/controller"
	db "subproblem/security-service/database"
	"subproblem/security-service/service"
	"github.com/gorilla/mux"
	"subproblem/security-service/util"
)

func main() {
	util.LoadEnv()

	db, err := db.NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	

	authService := service.NewAuthService(db)
	authController := controller.NewAuthController(authService)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/auth/register", authController.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", authController.Login).Methods("POST")
	router.HandleFunc("/api/v1/auth/validate", authController.ValidateToken).Methods("POST")

	http.ListenAndServe(":8081", router)
}