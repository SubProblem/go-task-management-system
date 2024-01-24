package service

import (
	"fmt"
	"subproblem/management-service/database"
	"subproblem/management-service/dto"
	"subproblem/management-service/producer"
	"time"
)

type TaskService struct {
	db database.PostgresDb
	mp *producer.MessageProducer
}

func NewTaskService(database *database.PostgresDb, producer *producer.MessageProducer) *TaskService {
	return &TaskService{
		db: *database,
		mp: producer,
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

func (task *TaskService) CompleteTask(taskId, userId int) error {

	err := task.db.CompleteTask(taskId, userId)

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

func (task *TaskService) CheckDeadline() {

	interval := time.Duration(10) * time.Second

	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				tasks, err := task.db.GetAllTasksByDeadline()

				if err != nil {
					fmt.Println(err)
					continue
				}

				for _, t := range tasks {
					fmt.Println(t)
					msg := &producer.Message{
						TaskId: t.ID,
						Task:   t.Task,
						UserId: t.User_ID,
					}
					task.mp.ProduceMessage(msg)
				}
			}
		}
	}()

}
