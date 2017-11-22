package tasks

import (
	"encoding/json"
	"io/ioutil"
)

// TaskList is a slice of Tasks
type TaskList []Task

// LoadFromFile loads a TaskList from file
func LoadFromFile(filename string) (TaskList, error) {
	bs, err := ioutil.ReadFile(filename)
	tasks := new(TaskList)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, tasks)

	if err != nil {
		return nil, err
	}

	return *tasks, nil
}

// SaveToFile accepts a TaskList reference and writes the serialized version to file.
func SaveToFile(tasks *TaskList, filename string) error {
	bs, err := json.Marshal(*tasks)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bs, 0644)

	if err != nil {
		return err
	}

	return nil
}

// Add a Task to the list
func (tl *TaskList) Add(t Task) {
	*tl = append(*tl, t)
}

// Remove a Task from the list
func (tl *TaskList) Remove(t Task) {
	foundIndex := -1

	for i, task := range *tl {
		if task.Id == t.Id {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		(*tl)[foundIndex], (*tl)[0] = (*tl)[0], (*tl)[foundIndex]
		(*tl) = (*tl)[1:]
	}
}
