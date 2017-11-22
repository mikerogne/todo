package tasks

import uuid "github.com/satori/go.uuid"

// Task is a simple task type that stores relevant info & state.
type Task struct {
	Id          string
	Description string
	Completed   bool
}

// NewTask creates and returns a new Task with the provided description. Completed is set to false.
func NewTask(description string) Task {
	task := Task{Id: uuid.NewV4().String(), Description: description, Completed: false}

	return task
}

// MarkComplete marks the task complete.
func (t *Task) MarkComplete() {
	t.Completed = true
}

// MarkIncomplete marks the task incomplete.
func (t *Task) MarkIncomplete() {
	t.Completed = false
}
