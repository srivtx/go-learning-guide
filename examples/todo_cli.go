package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoList struct {
	tasks  []Task
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (tl *TodoList) AddTask(title, description string) {
	task := Task{
		ID:          tl.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	tl.tasks = append(tl.tasks, task)
	tl.nextID++
	fmt.Printf("✅ Added task: %s\n", title)
}

func (tl *TodoList) ListTasks() {
	if len(tl.tasks) == 0 {
		fmt.Println("📝 No tasks found. Add some tasks to get started!")
		return
	}

	fmt.Println("\n📋 Your Tasks:")
	fmt.Println(strings.Repeat("-", 50))
	
	for _, task := range tl.tasks {
		status := "⭕"
		if task.Completed {
			status = "✅"
		}
		
		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
		if task.Description != "" {
			fmt.Printf("    %s\n", task.Description)
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}

func (tl *TodoList) CompleteTask(id int) {
	for i := range tl.tasks {
		if tl.tasks[i].ID == id {
			tl.tasks[i].Completed = true
			fmt.Printf("🎉 Completed task: %s\n", tl.tasks[i].Title)
			return
		}
	}
	fmt.Printf("❌ Task with ID %d not found\n", id)
}

func (tl *TodoList) DeleteTask(id int) {
	for i, task := range tl.tasks {
		if task.ID == id {
			// Remove task from slice
			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
			fmt.Printf("🗑️  Deleted task: %s\n", task.Title)
			return
		}
	}
	fmt.Printf("❌ Task with ID %d not found\n", id)
}

func (tl *TodoList) GetStats() {
	total := len(tl.tasks)
	completed := 0
	
	for _, task := range tl.tasks {
		if task.Completed {
			completed++
		}
	}
	
	pending := total - completed
	
	fmt.Println("\n📊 Statistics:")
	fmt.Printf("   Total tasks: %d\n", total)
	fmt.Printf("   Completed: %d\n", completed)
	fmt.Printf("   Pending: %d\n", pending)
	
	if total > 0 {
		percentage := float64(completed) / float64(total) * 100
		fmt.Printf("   Progress: %.1f%%\n", percentage)
	}
}

func printHelp() {
	fmt.Println("\n🔧 Available Commands:")
	fmt.Println("   add <title> [description]  - Add a new task")
	fmt.Println("   list                       - Show all tasks")
	fmt.Println("   complete <id>              - Mark task as completed")
	fmt.Println("   delete <id>                - Delete a task")
	fmt.Println("   stats                      - Show statistics")
	fmt.Println("   help                       - Show this help")
	fmt.Println("   exit                       - Exit the program")
}

func main() {
	todoList := NewTodoList()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🚀 Welcome to Go Todo List Manager!")
	fmt.Println("Type 'help' to see available commands.")

	// Add some sample tasks
	todoList.AddTask("Learn Go basics", "Complete the Go tutorial")
	todoList.AddTask("Build a project", "Create a simple CLI application")
	todoList.AddTask("Practice concurrency", "Learn about goroutines and channels")

	for {
		fmt.Print("\n> ")
		
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("❌ Usage: add <title> [description]")
				continue
			}
			
			title := parts[1]
			description := ""
			
			if len(parts) > 2 {
				description = strings.Join(parts[2:], " ")
			}
			
			todoList.AddTask(title, description)

		case "list":
			todoList.ListTasks()

		case "complete":
			if len(parts) != 2 {
				fmt.Println("❌ Usage: complete <id>")
				continue
			}
			
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("❌ Invalid task ID. Please enter a number.")
				continue
			}
			
			todoList.CompleteTask(id)

		case "delete":
			if len(parts) != 2 {
				fmt.Println("❌ Usage: delete <id>")
				continue
			}
			
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("❌ Invalid task ID. Please enter a number.")
				continue
			}
			
			todoList.DeleteTask(id)

		case "stats":
			todoList.GetStats()

		case "help":
			printHelp()

		case "exit", "quit":
			fmt.Println("👋 Thanks for using Go Todo List Manager!")
			return

		default:
			fmt.Printf("❌ Unknown command: %s\n", command)
			fmt.Println("Type 'help' to see available commands.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

/*
Example usage:

$ go run todo_cli.go

🚀 Welcome to Go Todo List Manager!
Type 'help' to see available commands.
✅ Added task: Learn Go basics
✅ Added task: Build a project
✅ Added task: Practice concurrency

> list

📋 Your Tasks:
--------------------------------------------------
⭕ 1. Learn Go basics
    Complete the Go tutorial
⭕ 2. Build a project
    Create a simple CLI application
⭕ 3. Practice concurrency
    Learn about goroutines and channels
--------------------------------------------------

> complete 1
🎉 Completed task: Learn Go basics

> add "Read Go documentation" "Study the official Go documentation"
✅ Added task: Read Go documentation

> stats

📊 Statistics:
   Total tasks: 4
   Completed: 1
   Pending: 3
   Progress: 25.0%

> exit
👋 Thanks for using Go Todo List Manager!
*/
