package tasks

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Make this test pass")

	if len(task.Id) < 5 {
		t.Error("Expected len(task.Id) to be >= 5, but got", len(task.Id))
	}

	if task.Description != "Make this test pass" {
		t.Error("Task Description not as expected:", task.Description)
	}

	if task.Completed != false {
		t.Error("New task `Completed` status should be false, got", task.Completed)
	}
}

func TestMarkComplete(t *testing.T) {
	task := NewTask("Make this test pass")

	task.MarkComplete()

	if task.Completed != true {
		t.Error("Expected task.Completed to be true, got", task.Completed)
	}
}

func TestMarkIncomplete(t *testing.T) {
	task := NewTask("Make this test pass")

	task.MarkComplete()
	task.MarkIncomplete()

	if task.Completed != false {
		t.Error("Expected task.Completed to be false, got", task.Completed)
	}
}
