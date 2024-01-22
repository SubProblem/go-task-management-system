package dto

import (
	"subproblem/management-service/models"
	"time"
)

type TaskResponseDto struct {
	ID        int       `json:"id"`
	User_ID   int       `json:"userID"`
	Task      string    `json:"task"`
	CreatedAt time.Time `json:"createdAt"`
	Deadline  time.Time `json:"deadline"`
}

type TaskRequestDto struct {
	Task     string    `json:"task"`
	Deadline time.Time `json:"deadline"`
}

func ToTaskResponseDto(task *models.Task) *TaskResponseDto {
	return &TaskResponseDto{
		ID:        task.ID,
		User_ID:   task.User_ID,
		Task:      task.Task,
		CreatedAt: task.CreatedAt,
		Deadline:  task.Deadline,
	}
}

func ToTask(task *TaskRequestDto, userId int) *models.Task {
	return &models.Task{
		User_ID:   userId,
		Task:      task.Task,
		CreatedAt: time.Now(),
		Deadline:  task.Deadline,
		Completed: false,
	}
}

func ReponsesToList(tasks []*models.Task) []*TaskResponseDto {

	var responses []*TaskResponseDto

	for _, t := range tasks {
		response := ToTaskResponseDto(t)
		responses = append(responses, response)
	}

	return responses
}
