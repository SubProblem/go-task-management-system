package main

import (
	"log"
	"net/http"
	"subproblem/management-service/controller"
	db "subproblem/management-service/database"
	"subproblem/management-service/service"
	"subproblem/management-service/util"

	"github.com/gorilla/mux"
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

	taskService := service.NewTaskService(db)
	
	taskService.PrintTasks()
	
	taskController := controller.NewTaskController(taskService)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/task", taskController.GetAllTasksForUserById).Methods("GET")
	router.HandleFunc("/api/v1/task", taskController.AddTask).Methods("POST")
	router.HandleFunc("/api/v1/task", taskController.DeleteTaskById).Methods("DELETE")

	http.ListenAndServe(":8082", router)
}
