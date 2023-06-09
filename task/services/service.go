package service

import (
	"backend/task/models"
	repository "backend/task/repositories"
	"errors"
	"fmt"
	"log"
	"time"

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

	task.CreatedAt = time.Now()
	task.UpdateAt = time.Now()

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
	taskFound, err := ts.taskRepository.GetById(taskId)
	if err != nil {
		return err
	}
	if taskFound == nil {
		return errors.New("task não encontrada")
	}

	taskNameDuplicated, err := ts.taskRepository.GetByDescription(task.Description)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	if taskNameDuplicated != nil && taskNameDuplicated.Id != taskId {
		return errors.New("já existe uma tarefa com essa descrição")
	}

	taskStatusFound, err := ts.taskStatusRepository.GetByName(task.Status.Name)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	if taskStatusFound == nil {
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}

	taskFound.Description = task.Description
	taskFound.Status = taskStatusFound
	taskFound.UpdateAt = time.Now()

	err = ts.taskRepository.Update(taskId, taskFound)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	return nil
}

func (ts *TaskService) UpdateStatus(taskId string, taskNewStatus string) error {
	taskFound, err := ts.taskRepository.GetById(taskId)
	if err != nil {
		return err
	}
	if taskFound == nil {
		return errors.New("task não encontrada")
	}

	taskStatusFound, err := ts.taskStatusRepository.GetByName(taskNewStatus)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	if taskStatusFound == nil {
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}

	taskFound.Status = taskStatusFound
	taskFound.UpdateAt = time.Now()

	err = ts.taskRepository.UpdateStatus(taskId, taskFound)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	return nil
}

func (ts *TaskService) GetAllStatus() ([]*models.TaskStatus, error) {
	return ts.taskStatusRepository.GetAllStatus()
}

func (ts *TaskService) Delete(taskId string) error {
	taskFound, err := ts.taskRepository.GetById(taskId)
	if err != nil {
		return err
	}
	if taskFound == nil {
		return errors.New("task não encontrada")
	}

	err = ts.taskRepository.Delete(taskFound.Id)
	if err != nil {
		// TODO: tratar melhor os erros do repositório
		log.Fatalf("%s", err.Error())
		return errors.New("houve um erro, tente novamente mais tarde")
	}
	return nil
}

func (ts *TaskService) GetAll() ([]*models.Task, error) {
	return ts.taskRepository.GetAll()
}

func (ts *TaskService) Get(task *models.Task) (*models.Task, error) {
	return nil, nil
}
