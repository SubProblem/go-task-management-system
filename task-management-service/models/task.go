package models

import "time"

type Task struct {
	ID        int
	User_ID   int
	Task      string
	CreatedAt time.Time
	Deadline  time.Time
	Completed bool
}
