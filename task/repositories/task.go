package repositories

import (
	"backend/task/models"
	"database/sql"
)

type TaskRepository struct {
	db                   *sql.DB
	taskStatusRepository *TaskStatus
}

func NewTaskRepository(db *sql.DB, taskStatusRepository *TaskStatus) *TaskRepository {
	return &TaskRepository{
		db:                   db,
		taskStatusRepository: taskStatusRepository,
	}
}

func (ts *TaskRepository) Create(task *models.Task) error {
	sql := `
		INSERT INTO task(id, description, id_task_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := ts.db.Exec(sql, task.Id, task.Description, task.Status.Id, task.CreatedAt, task.UpdateAt)
	return err
}

func (ts *TaskRepository) Update(taskId string, task *models.Task) error {
	return nil
}

func (ts *TaskRepository) UpdateStatus(newStatus *models.TaskStatus, task *models.Task) error {
	return nil
}

func (ts *TaskRepository) Delete(task *models.Task) (*models.Task, error) {
	return nil, nil
}

func (ts *TaskRepository) GetAll(task *models.Task) ([]*models.Task, error) {
	return []*models.Task{}, nil
}

func (ts *TaskRepository) GetByDescription(description string) (*models.Task, error) {
	sqlQuery := `
		SELECT id, description, created_at, updated_at, id_task_status
		FROM task
		where description = $1 
	`

	row := ts.db.QueryRow(sqlQuery, description)
	if row.Err() != nil {
		return nil, row.Err()
	}

	task := &models.Task{}
	taskStatusId := ""
	err := row.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.UpdateAt, &taskStatusId)
	if row.Err() != nil && row.Err() != sql.ErrNoRows {
		return nil, row.Err()
	}
	if row.Err() != nil && row.Err() == sql.ErrNoRows {
		return nil, nil
	}

	task.Status, err = ts.taskStatusRepository.Get(taskStatusId)
	if err != nil {
		return nil, err
	}

	return task, nil
}
