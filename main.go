package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global slice of tasks
var tasks []Task

// ✅ Struct fields must be exported (capitalized) to be JSON serialized
// ✅ Removed invalid `int.` and `string.`
// ✅ Fixed struct tags
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// ✅ Fixed error handling with proper checks
func updateJSON() {
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error during JSON marshal:", err)
		return
	}
	if err := os.WriteFile("tasks.json", jsonData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
		// ✅ Load tasks from tasks.json at startup
	fileData, err := os.ReadFile("tasks.json")
	if err == nil { // file exists
		if len(fileData) > 0 {
			err = json.Unmarshal(fileData, &tasks)
			if err != nil {
				fmt.Println("Error parsing tasks.json:", err)
			} else {
				fmt.Println("Loaded tasks from tasks.json")
			}
		}
	} else {
		// If file doesn't exist, no problem — start fresh
		fmt.Println("No existing tasks.json found, starting with empty task list")
	}

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Enter the process you want to take (add/change_status/update/delete/list/exit): ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		text = strings.ToLower(strings.TrimSpace(text))

		switch text {
		case "add":
			addNewTask(reader)
		case "change_status":
			changeStatus(reader)
		case "update":
			updateExistingTask(reader)
		case "delete":
			deleteTask(reader)
		case "list":
			listTasks()
		case "exit":
			fmt.Println("Exiting program")
			return
		default:
			fmt.Println("Please enter a valid choice")
		}

		updateJSON()
	}
}

func changeStatus(reader *bufio.Reader) {
	fmt.Print("Enter the task id: ")
	taskIDStr, _ := reader.ReadString('\n')
	taskIDStr = strings.TrimSpace(taskIDStr)
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	fmt.Print("Enter task status: ")
	taskStatus, _ := reader.ReadString('\n')
	taskStatus = strings.TrimSpace(taskStatus)

	found := false
	for index, value := range tasks {
		if value.ID == taskID {
			tasks[index].Status = taskStatus
			fmt.Println("Updated task is:", tasks[index])
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Couldn't find the task with ID", taskID)
	}
}

func addNewTask(reader *bufio.Reader) {
	if len(tasks) >= 10 {
		fmt.Println("First, complete previous tasks!")
		return
	}

	fmt.Print("Enter the task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	fmt.Print("Enter task description: ")
	taskDescription, _ := reader.ReadString('\n')
	taskDescription = strings.TrimSpace(taskDescription)

	newTask := Task{
		ID:          len(tasks) + 1,
		Name:        taskName,
		Description: taskDescription,
		Status:      "TODO",
	}

	tasks = append(tasks, newTask)
	fmt.Println("Task added successfully")
}

func updateExistingTask(reader *bufio.Reader) {
	fmt.Print("Enter the task id you want to update: ")
	taskIDStr, _ := reader.ReadString('\n')
	taskID, err := strconv.Atoi(strings.TrimSpace(taskIDStr))
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	fmt.Print("Enter the field you want to update (name/description): ")
	updateField, _ := reader.ReadString('\n')
	updateField = strings.TrimSpace(updateField)

	found := false
	for index, value := range tasks {
		if value.ID == taskID {
			if updateField == "name" {
				fmt.Print("Enter the new name: ")
				newName, _ := reader.ReadString('\n')
				tasks[index].Name = strings.TrimSpace(newName)
			} else if updateField == "description" {
				fmt.Print("Enter the new description: ")
				newDescription, _ := reader.ReadString('\n')
				tasks[index].Description = strings.TrimSpace(newDescription)
			} else {
				fmt.Println("Invalid field (must be name or description)")
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

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter the task id you want to delete: ")
	taskIDStr, _ := reader.ReadString('\n')
	taskID, err := strconv.Atoi(strings.TrimSpace(taskIDStr))
	if err != nil {
		fmt.Println("Invalid task ID, must be a number")
		return
	}

	found := false
	for index, value := range tasks {
		if value.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			listTasks()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Couldn't find the task with ID", taskID)
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error converting tasks to JSON:", err)
		return
	}

	fmt.Println("Tasks list is: ",string(jsonData))
}
