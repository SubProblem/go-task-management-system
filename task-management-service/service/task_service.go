package service

import (
	"fmt"
	"subproblem/management-service/database"
	"subproblem/management-service/dto"
	"time"
)

type TaskService struct {
	db database.PostgresDb
}


func NewTaskService(database *database.PostgresDb) *TaskService {
	return &TaskService{
		db: *database,
	}
}

func (task *TaskService) AddTask(request *dto.TaskRequestDto, userId int) error {

	newTask := dto.ToTask(request, userId)
	newTask.CreatedAt = time.Now()

	err := task.db.AddTask(newTask)

	if err != nil {
		return err
	}

	return nil

}

func (task *TaskService) GetAllTasksForUserById(id int) ([]*dto.TaskResponseDto, error) {

	tasks, err := task.db.GetAllTasksForUserById(id)

	if err != nil {
		return nil, err
	}

	
	return dto.ReponsesToList(tasks), nil
}

func (task *TaskService) DeleteTaskById(id int) (bool, error) {
	
	ok, err := task.db.DeleteTaskById(id)

	if err != nil {
		return false, err
	}

	if ok != true {
		return false, nil
	}

	return true, nil
}



func (task *TaskService) CheckDeadline() error {

	// Retreive Data from Database based on deadline

	users, err := task.db.GetAllTasksByDeadline()


	if err != nil {
		return err
	}

	for _, t := range users {
		fmt.Println(t)
	}

	return nil
}

func (task *TaskService) PrintTasks() {
	interval := time.Duration(10) * time.Second

	ticker := time.NewTicker(interval)


	go func() {
		for {
			select {
			case <- ticker.C:
				if err := task.CheckDeadline(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}

