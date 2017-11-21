package tasks

// TaskList is a slice of Tasks
type TaskList []Task

// Add a Task to the list
func (tl *TaskList) Add(t Task) {
	*tl = append(*tl, t)
}

// Remove a Task from the list
func (tl *TaskList) Remove(t Task) {
	foundIndex := -1

	for i, task := range *tl {
		if task.id == t.id {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		(*tl)[foundIndex], (*tl)[0] = (*tl)[0], (*tl)[foundIndex]
		(*tl) = (*tl)[1:]
	}
}
