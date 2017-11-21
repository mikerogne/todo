package tasks

import uuid "github.com/satori/go.uuid"

// Task is a simple task type that stores relevant info & state.
type Task struct {
	id          string
	description string
	completed   bool
}

// NewTask creates and returns a new Task with the provided description. Completed is set to false.
func NewTask(description string) Task {
	task := Task{id: uuid.NewV4().String(), description: description, completed: false}

	return task
}

// MarkComplete marks the task complete.
func (t *Task) MarkComplete() {
	t.completed = true
}

// MarkIncomplete marks the task incomplete.
func (t *Task) MarkIncomplete() {
	t.completed = false
}
