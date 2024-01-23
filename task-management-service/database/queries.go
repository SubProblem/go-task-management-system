package database

import (
	"fmt"
	"subproblem/management-service/models"

	_ "github.com/lib/pq"
)

func (pg *PostgresDb) CreateTasksTable() error {

	query := `
		CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		task VARCHAR(250) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		deadline DATE NOT NULL,
		completed BOOLEAN NOT NULL 
		)
	`
	_, err := pg.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (pg *PostgresDb) AddTask(task *models.Task) error {

	query := `
		INSERT INTO tasks
		(task, user_id, created_at, deadline, completed)
		VALUES
		($1, $2, $3, $4, $5)
	`

	res, err := pg.db.Query(
		query,
		task.Task, task.User_ID, task.CreatedAt, task.Deadline, task.Completed,
	)

	if err != nil {
		return err
	}
	defer res.Close()

	fmt.Printf("res: %v\n", res)
	return nil

}

func (pg *PostgresDb) CompleteTask(taskId, userId int) error {

	query := `
		UPDATE tasks
		SET completed = 'true'
		WHERE id = $1 and user_id = $2
	`

	res, err := pg.db.Exec(query, taskId, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}


func (pg *PostgresDb) GetAllTasksByDeadline() ([]*models.Task, error) {

	query := `
		SELECT * FROM tasks
		WHERE deadline = CURRENT_DATE
		
	`
	// WHERE deadline = CURRENT_DATE - INTERVAL '2 days
	res, err := pg.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	var tasks []*models.Task

	for res.Next() {
		task := &models.Task{}

		err := res.Scan(
			&task.ID,
			&task.User_ID,
			&task.Task,
			&task.CreatedAt,
			&task.Deadline,
			&task.Completed,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (pg *PostgresDb) DeleteTaskById(id int) (bool, error) {

	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	res, err := pg.db.Exec(query, id)

	if err != nil {
		return false, nil
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (pg *PostgresDb) GetAllTasksForUserById(userId int) ([]*models.Task, error) {

	query := `
		SELECT * FROM tasks
		WHERE user_id = $1
	`
	var tasks []*models.Task

	res, err := pg.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {

		task := &models.Task{}

		err := res.Scan(
			&task.ID,
			&task.User_ID,
			&task.Task,
			&task.CreatedAt,
			&task.Deadline,
			&task.Completed,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
