package tasks

import (
	"testing"
)

func TestCanAddTasks(t *testing.T) {
	taskList := TaskList{}
	task1 := NewTask("First task")
	task2 := NewTask("Second task")
	task3 := NewTask("Third task")

	taskList.Add(task1)
	taskList.Add(task2)
	taskList.Add(task3)

	if count := len(taskList); count != 3 {
		t.Error("Expected len(taskList) to be 3, got", count)
	}
}

func TestCanRemoveTasks(t *testing.T) {
	taskList := TaskList{}
	task1 := NewTask("First task")
	task2 := NewTask("Second task")
	task3 := NewTask("Third task")

	taskList.Add(task1)
	taskList.Add(task2)
	taskList.Add(task3)

	taskList.Remove(task2)

	if count := len(taskList); count != 2 {
		t.Error("Expected len(taskList) to be 2, got", count)
	}

	if taskList[0].Description != "First task" {
		t.Error("Expected taskList[0] to be 'First task', got", taskList[0].Description)
	}

	if taskList[1].Description != "Third task" {
		t.Error("Expected taskList[1] to be 'Third task', got", taskList[1].Description)
	}
}
