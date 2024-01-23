package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"subproblem/management-service/dto"
	"subproblem/management-service/service"

	"github.com/gorilla/mux"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{
		service: taskService,
	}
}

func (task *TaskController) GetAllTasksForUserById(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Incorrect user Id", http.StatusBadRequest)
		return
	}

	tasks, err := task.service.GetAllTasksForUserById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasksJson, err := json.Marshal(tasks)

	if err != nil {
		http.Error(w, "Error converting tasks to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(tasksJson)
}


func (task *TaskController) AddTask(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")
	fmt.Printf("userId: %v\n", userId)
	
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Incorrect user Id", http.StatusBadRequest)
		return
	}

	var taskRequest dto.TaskRequestDto

	err2 := json.NewDecoder(r.Body).Decode(&taskRequest)

	if err2 != nil {
		http.Error(w, "Error decoding the body", http.StatusInternalServerError)
		return
	}

	err3 := task.service.AddTask(&taskRequest, id)

	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}


func (task *TaskController) CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	taskId, err1 := strconv.Atoi(vars["taskId"])
	userId, err2 := strconv.Atoi(vars["userId"])

	if err1 != nil || err2 != nil {
		http.Error(w, "Incorrect path variable", http.StatusBadRequest)
		return
	}

	err := task.service.CompleteTask(taskId, userId)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	return
}

func (task *TaskController) DeleteTaskById(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("X-User-Id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Incorrect user Id", http.StatusBadRequest)
		return
	}

	ok, err := task.service.DeleteTaskById(id)

	if err != nil {
		http.Error(w, "Something went wrong, Task could not be deleted", http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Something went wrong, Task could not be deleted", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
