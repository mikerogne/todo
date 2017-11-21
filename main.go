package main

import (
	"bufio"
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

	fmt.Printf("%+v", taskList)
}