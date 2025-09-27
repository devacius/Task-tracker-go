# Task Tracker Go

Task Tracker Go is a simple command-line application written in Go for managing your daily tasks. It allows you to add, update, change status, delete, and list tasks, with all data stored in a local JSON file (`tasks.json`).

## Features

- Add new tasks with a name and description
- Update task name or description
- Change the status of a task (e.g., TODO, IN PROGRESS, DONE)
- Delete tasks
- List all tasks in a readable JSON format
- Persistent storage in `tasks.json`
- Limit of 10 tasks at once to encourage focus and completion

## Installation

1. **Install Go** (if not already installed):
	- Download and install from [golang.org/dl](https://golang.org/dl/)
2. **Clone the repository:**
	```sh
	git clone https://github.com/devacius/tasko.git
	cd tasko
	```
3. **Build the project:**
	```sh
	go build -o task-tracker main.go
	```

## Usage

Run the application from your terminal:

```sh
go install github.com/devacius/tasko@latest
tasko
```

You will be prompted to enter a process:

- `add` — Add a new task
- `change_status` — Change the status of a task
- `update` — Update the name or description of a task
- `delete` — Delete a task
- `list` — List all tasks
- `exit` — Exit the program

Example session:

```
Enter the process you want to take (add/change_status/update/delete/list/exit): add
Enter the task name: Write documentation
Enter task description: Complete the README file
Task added successfully
```

All changes are automatically saved to `tasks.json`.

## Why is there a limit of 10 tasks?

The application restricts the number of active tasks to 10. This design choice is intentional to help users focus on completing existing tasks before adding more. Managing too many tasks at once can lead to overwhelm and reduced productivity. By limiting the list, Task Tracker Go encourages you to prioritize, complete, and clear tasks, fostering better task management habits.

## License

MIT

