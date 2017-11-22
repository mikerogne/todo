package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"todo/tasks"
)

func main() {
	jsonFilename := "tasks.json"
	var taskList tasks.TaskList

	if _, err := os.Stat(jsonFilename); !os.IsNotExist(err) {
		taskList, err = tasks.LoadFromFile(jsonFilename)

		if err != nil {
			fmt.Printf("Error loading %s: %s", jsonFilename, err)
			os.Exit(1)
		}
	} else {
		taskList = tasks.TaskList{}
	}

	displayHelp("Welcome!")

	for {
		input := ""
		second := ""

		fmt.Printf("Command> ")
		fmt.Scanf("%s %s", &input, &second)

		switch strings.ToLower(input) {
		case "help":
			displayHelp("Help:")

		case "list":
			displayTasks(&taskList)

		case "add":
			newTask, err := getNewTask()
			if err != nil {
				fmt.Println("Error getting task:", err)
				continue
			}

			taskList.Add(newTask)
			err = tasks.SaveToFile(&taskList, jsonFilename)

			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(1)
			}

			fmt.Println("Task saved!")

		case "del":
			// todo

		case "com":
			// todo

		case "inc":
			// todo

		case "quit":
			os.Exit(0)

		case "q":
			os.Exit(0)
		}
	}
}

func getNewTask() (tasks.Task, error) {
	scanner := bufio.NewScanner(os.Stdin)
	taskName := ""

	fmt.Printf("Task name: ")
	scanner.Scan()
	taskName = scanner.Text()
	newTask := tasks.NewTask(taskName)

	if len(strings.TrimSpace(taskName)) == 0 {
		return newTask, errors.New("You cannot enter a blank task")
	}

	return newTask, nil
}

func displayHelp(title string) {
	fmt.Print(title, `
-----
help		display this menu
list		list all tasks (all, incomplete, complete)
add		add new task
del <#>	delete a task (specify last 5+ characters)
com <#>	mark a task complete (specify last 5+ characters)
inc <#>	mark a task incomplete (specify last 5+ characters)
quit|q		exit the application

`)
}

func displayTasks(taskList *tasks.TaskList) {
	fmt.Printf("%5s  %10s %s %s\n", "Index", "Status", "UUID", "Description")

	for i, task := range *taskList {
		status := "Complete"

		if task.Completed == false {
			status = "Incomplete"
		}

		fmt.Printf("%5d  %10s %s %s\n", i, status, task.Id, task.Description)
	}
}
