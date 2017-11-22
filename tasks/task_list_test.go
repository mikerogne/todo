package tasks

import (
	"io/ioutil"
	"os"
	"strings"
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

func TestSaveToFile(t *testing.T) {
	// Create task list, save to file.
	tasks := TaskList{NewTask("This is a test")}
	SaveToFile(&tasks, "test.json")

	// Read the raw json file.
	bs, _ := ioutil.ReadFile("test.json")
	tasksJSON := string(bs)
	os.Remove("test.json")

	// See if it contains our expected data.
	if !strings.Contains(tasksJSON, `"Description":"This is a test"`) {
		t.Error("Expected to see description in tasksJSON.")
	}

	if !strings.Contains(tasksJSON, `"Completed":false`) {
		t.Error("Expected to see Completed=false in tasksJSON.")
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create task list, save to file.
	tasks := TaskList{NewTask("This is a test")}
	SaveToFile(&tasks, "test.json")

	// Now load task list from file.
	tasks, err := LoadFromFile("test.json")

	if err != nil {
		t.Error("Error loading from file:", err)
	}
	defer os.Remove("test.json")

	if len(tasks) != 1 {
		t.Error("Expected len(tasks) to be 1, got", len(tasks))
	}

	if len(tasks[0].Id) == 0 {
		t.Error("Expected tasks[0].Id to be set.")
	}

	if tasks[0].Description != "This is a test" {
		t.Error(`"Expected tasks[0].Description to be "This is a test", got`, tasks[0].Description)
	}

	if tasks[0].Completed != false {
		t.Error("Expected tasks[0].Completed to be false, got", tasks[0].Completed)
	}
}
