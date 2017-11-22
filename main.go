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
	taskList, err := tasks.LoadFromFile(jsonFilename)

	if err != nil {
		fmt.Printf("Error loading %s: %s", jsonFilename, err)
		os.Exit(1)
	}

	for {
		input := ""
		second := ""

		displayHelp("Welcome!")
		fmt.Printf("Command> ")
		fmt.Scanf("%s %s", &input, &second)

		switch strings.ToLower(input) {
		case "help":
			displayHelp("Help:")
		case "list":
			// todo
		case "add":
			var newTask tasks.Task
			if newTask, err = getNewTask(); err != nil {
				fmt.Println("Error getting task:", err)
				os.Exit(1)
			}

			taskList.Add(newTask)

			// Write to file
			tasks.SaveToFile(&taskList, jsonFilename)
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
del <id>	delete a task (specify last 5+ characters)
com <id>	mark a task complete (specify last 5+ characters)
inc <id>	mark a task incomplete (specify last 5+ characters)
quit|q		exit the application

`)
}
