package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // Added import for strconv.Atoi
	"strings" // Added import for strings.TrimSpace and strings.ToLower
)

// Declared tasks as a global variable to make it accessible to all functions
// Old: var tasks []task (inside main, causing undefined variable errors)
// New: Global declaration
var tasks []task

type task struct {
	id          int
	name        string
	description string
	status      string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Added loop to allow multiple actions until user chooses to exit
	for {
		fmt.Println("Hello World")
		fmt.Print("Enter the process you want to take (add/change_status/update/delete/list/exit): ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		// Trim input and convert to lowercase for case-insensitive comparison
		// Old: if (text == "add") { ... }
		// New: strings.ToLower(strings.TrimSpace(text)) to handle newlines and case variations
		text = strings.ToLower(strings.TrimSpace(text))

		switch text {
		case "add":
			add_new_task(reader) // Pass reader to reuse it
		case "change_status":
			change_status(reader)
		case "update":
			update_existing_task(reader)
		case "delete":
			delete_task(reader)
		case "list":
			list_tasks()
		case "exit":
			fmt.Println("Exiting program")
			return
		default:
			fmt.Println("Please enter a valid choice")
		}
	}
}

func change_status(reader *bufio.Reader) {
	fmt.Print("Enter the task id: ")
	taskIDStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}
	// Trim input and convert to int
	// Old: task_id, _ := reader.ReadString('\n'); if value.id == task_id
	// New: Convert taskIDStr to int and handle error
	taskIDStr = strings.TrimSpace(taskIDStr)
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	found := false
	// Update tasks slice directly using index
	// Old: value.status = task_status (modifying loop copy, no effect)
	// New: tasks[index].status = taskStatus
	fmt.Print("Enter task status: ")
	taskStatus, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading status:", err)
		return
	}
	taskStatus = strings.TrimSpace(taskStatus)

	for index, value := range tasks {
		if value.id == taskID {
			tasks[index].status = taskStatus
			fmt.Println("Updated task is:", tasks[index])
			found = true
			break
		}
	}
	// Print error only if task not found
	// Old: Error printed in every loop iteration
	// New: Use found flag to print error once
	if !found {
		fmt.Println("Couldn't find the task with ID", taskID)
	}
}

func add_new_task(reader *bufio.Reader) {
	length := len(tasks)
	if length >= 10 {
		fmt.Println("First, complete previous tasks!")
		return
	}
	// Old: new_length := length + 1 (unnecessary variable)
	// New: Use length + 1 directly in struct
	fmt.Print("Enter the task name: ")
	taskName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task name:", err)
		return
	}
	taskName = strings.TrimSpace(taskName)

	fmt.Print("Enter task description: ")
	taskDescription, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading description:", err)
		return
	}
	taskDescription = strings.TrimSpace(taskDescription)

	// Corrected struct initialization and used "TODO" as string
	// Old: var new_task task= {id:new_length, name:task_name description:task_description status:TODO}
	// New: Proper struct initialization with commas and quoted "TODO"
	newTask := task{
		id:          length + 1,
		name:        taskName,
		description: taskDescription,
		status:      "TODO",
	}

	tasks = append(tasks, newTask)
	fmt.Println("Task added successfully")
}

func update_existing_task(reader *bufio.Reader) {
	fmt.Print("Enter the task id you want to update: ")
	taskIDStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}
	// Convert task ID to int
	// Old: value.id == task_id (type mismatch)
	// New: Convert taskIDStr to int
	taskIDStr = strings.TrimSpace(taskIDStr)
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	found := false
	fmt.Print("Enter the field you want to update (name/description): ")
	updateField, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading field:", err)
		return
	}
	updateField = strings.TrimSpace(updateField)

	for index, value := range tasks {
		if value.id == taskID {
			if updateField == "name" {
				fmt.Print("Enter the new name: ")
				newName, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading new name:", err)
					return
				}
				// Update slice directly
				// Old: value.name = new_name
				// New: tasks[index].name = newName
				tasks[index].name = strings.TrimSpace(newName)
			} else if updateField == "description" {
				fmt.Print("Enter the new description: ")
				newDescription, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading new description:", err)
					return
				}
				tasks[index].description = strings.TrimSpace(newDescription)
			} else {
				fmt.Println("You need to select a correct field to update (name/description)")
				return
			}
			fmt.Println("Final updated task:", tasks[index])
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Couldn't find the task with ID", taskID)
	}
}

func delete_task(reader *bufio.Reader) {
	fmt.Print("Enter the task id you want to delete: ")
	taskIDStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}
	// Convert task ID to int
	// Old: value.id == task_id (type mismatch)
	// New: Convert taskIDStr to int
	taskIDStr = strings.TrimSpace(taskIDStr)
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	found := false
	// Fix slice deletion
	// Old: tasks = append(tasks[:index], s[index+1:]...) (undefined s)
	// New: Use tasks instead of s
	for index, value := range tasks {
		if value.id == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task deleted. New task list:", tasks)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Couldn't find the task with ID", taskID)
	}
}

func list_tasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	fmt.Println("Task list is:", tasks)
}