package service

import (
	"backend/task/models"
	repository "backend/task/repositories"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository       *repository.TaskRepository
	taskStatusRepository *repository.TaskStatus
}

func NewTaskService(tr *repository.TaskRepository, tsr *repository.TaskStatus) *TaskService {
	return &TaskService{
		taskRepository:       tr,
		taskStatusRepository: tsr,
	}
}

func (ts *TaskService) Create(task *models.Task) (*models.Task, error) {
	task.Id = fmt.Sprintf("%s", uuid.New())

	status, err := ts.taskStatusRepository.GetByName(models.TASK_STATUS_TODO)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return nil, errors.New("houve um erro, tente novamente mais tarde")
	}
	task.Status = status

	taskNameDuplicated, err := ts.taskRepository.GetByDescription(task.Description)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return nil, errors.New("houve um erro, tente novamente mais tarde")
	}
	if taskNameDuplicated != nil {
		return nil, errors.New("já existe uma tarefa com essa descrição")
	}

	err = ts.taskRepository.Create(task)
	return task, err
}

func (ts *TaskService) Update(taskId string, task *models.Task) error {
	return nil
}

func (ts *TaskService) UpdateStatus(newStatus *models.TaskStatus, task *models.Task) error {
	return nil
}

func (ts *TaskService) Delete(task *models.Task) (*models.Task, error) {
	return nil, nil
}

func (ts *TaskService) GetAll() ([]*models.Task, error) {
	return ts.taskRepository.GetAll()
}

func (ts *TaskService) Get(task *models.Task) (*models.Task, error) {
	return nil, nil
}
