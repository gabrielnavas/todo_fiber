package models

import "errors"

const (
	TASK_STATUS_TODO   string = "todo"
	TASK_STATUS_DOING         = "doing"
	TASK_STATUS_FINISH        = "finish"
)

type TaskStatus struct {
	Id   string
	Name string
}

func (s *TaskStatus) Validate() error {
	var err error = nil

	if len(s.Name) < 2 {
		err = errors.New("nome do status deve ser maior que 2")
		return err
	}

	return nil
}
