package repositories

import (
	"backend/task/models"
	"database/sql"
)

type TaskStatus struct {
	db *sql.DB
}

func NewTaskStatusRepository(db *sql.DB) *TaskStatus {
	return &TaskStatus{
		db: db,
	}
}

func (ts *TaskStatus) GetByName(name string) (*models.TaskStatus, error) {
	sqlQuery := `
		SELECT id, name
		FROM task_status
		where name = $1
	`

	row := ts.db.QueryRow(sqlQuery, name)
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return nil, row.Err()
	}
	if row.Err() != nil && row.Err() == sql.ErrNoRows {
		return nil, nil
	}

	taskStatus := &models.TaskStatus{}
	err := row.Scan(&taskStatus.Id, &taskStatus.Name)
	if err != nil {
		return nil, err
	}

	return taskStatus, nil
}

func (ts *TaskStatus) Get(id string) (*models.TaskStatus, error) {
	sqlQuery := `
		SELECT name
		FROM task_status
		where id = $1
	`

	row := ts.db.QueryRow(sqlQuery, id)
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return nil, row.Err()
	}
	if row.Err() != nil && row.Err() == sql.ErrNoRows {
		return nil, nil
	}

	taskStatus := &models.TaskStatus{}
	taskStatus.Id = id
	err := row.Scan(&taskStatus.Name)
	if err != nil {
		return nil, err
	}

	return taskStatus, nil
}
