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
	sqlStatement := `
		INSERT INTO task(id, description, id_task_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := ts.db.Exec(sqlStatement, task.Id, task.Description, task.Status.Id, task.CreatedAt, task.UpdateAt)
	return err
}

func (ts *TaskRepository) Update(taskId string, task *models.Task) error {
	sqlStatement := `
		UPDATE task
		SET 
			description = $1,
			id_task_status = $2,
			updated_at = $3
		WHERE
			id = $4;
	`
	_, err := ts.db.Exec(sqlStatement, task.Description, task.Status.Id, task.UpdateAt, taskId)
	return err
}

func (ts *TaskRepository) UpdateStatus(newStatus *models.TaskStatus, task *models.Task) error {
	return nil
}

func (ts *TaskRepository) Delete(task *models.Task) (*models.Task, error) {
	return nil, nil
}

func (ts *TaskRepository) GetAll() ([]*models.Task, error) {
	sqlStatement := `
		SELECT id, description, created_at, updated_at, id_task_status
		FROM task;
	`
	rows, err := ts.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks = []*models.Task{}

	for rows.Next() {
		var task = &models.Task{}
		var taskStatusId string = ""
		if err := rows.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.UpdateAt, &taskStatusId); err != nil {
			return tasks, err
		}
		task.Status, err = ts.taskStatusRepository.Get(taskStatusId)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (ts *TaskRepository) GetByDescription(description string) (*models.Task, error) {
	sqlStatement := `
		SELECT id, description, created_at, updated_at, id_task_status
		FROM task
		where description = $1 
	`
	task := &models.Task{}
	taskStatusId := ""
	row := ts.db.QueryRow(sqlStatement, description)
	switch err := row.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.UpdateAt, &taskStatusId); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		task.Status, err = ts.taskStatusRepository.Get(taskStatusId)
		if err != nil {
			return nil, err
		}
		return task, nil
	default:
		return nil, err
	}
}

func (ts *TaskRepository) GetById(id string) (*models.Task, error) {
	sqlStatement := `
		SELECT id, description, created_at, updated_at, id_task_status
		FROM task
		where id = $1 
	`
	task := &models.Task{}
	taskStatusId := ""
	row := ts.db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.UpdateAt, &taskStatusId); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		task.Status, err = ts.taskStatusRepository.Get(taskStatusId)
		if err != nil {
			return nil, err
		}
		return task, nil
	default:
		return nil, err
	}
}
