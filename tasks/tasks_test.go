package tasks

import "testing"

func TestNewTask(t *testing.T) {
	task := NewTask("Make this test pass")

	if len(task.id) < 5 {
		t.Error("Expected len(task.id) to be >= 5, but got", len(task.id))
	}

	if task.description != "Make this test pass" {
		t.Error("Task description not as expected:", task.description)
	}

	if task.completed != false {
		t.Error("New task `completed` status should be false, got", task.completed)
	}
}

func TestMarkComplete(t *testing.T) {
	task := NewTask("Make this test pass")

	task.MarkComplete()

	if task.completed != true {
		t.Error("Expected task.completed to be true, got", task.completed)
	}
}

func TestMarkIncomplete(t *testing.T) {
	task := NewTask("Make this test pass")

	task.MarkComplete()
	task.MarkIncomplete()

	if task.completed != false {
		t.Error("Expected task.completed to be false, got", task.completed)
	}
}
