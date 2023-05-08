package models

import (
	"errors"
	"time"
)

type Task struct {
	Id          string
	Description string
	Status      *TaskStatus
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func (t *Task) Validate() error {
	var err error = nil

	if len(t.Description) == 0 {
		err = errors.New("descrição não pode ser vazia")
		return err
	}
	if len(t.Description) > 255 {
		err = errors.New("descrição não pode ser maior que 255 caracteres")
		return err
	}
	if t.Status == nil {
		err = errors.New("defina um status")
		return err
	}

	errStatus := t.Status.Validate()
	if errStatus == nil {
		err = errStatus
		return err
	}

	return nil
}
