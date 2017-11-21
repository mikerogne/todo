package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"todo/tasks"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	taskList := new(tasks.TaskList)
	taskName := ""

	fmt.Printf("Task name: ")
	scanner.Scan()
	taskName = scanner.Text()

	taskList.Add(tasks.NewTask(taskName))

	// TESTING
	bs, _ := json.Marshal(taskList)
	fmt.Println(string(bs))
	// ^ does not work as expected. :|

	tasks.SaveToFile(taskList, "tasks.json")

	fmt.Printf("%+v", taskList)
}
