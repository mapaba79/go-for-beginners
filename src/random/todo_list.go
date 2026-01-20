package main

import "fmt"

type Task struct {
	ID   int
	Text string
}

type TodoList struct {
	Tasks []Task
}

func (l *TodoList) Add(text string) {
	newID := len(l.Tasks) + 1
	l.Tasks = append(l.Tasks, Task{ID: newID, Text: text})
}

func (l TodoList) Show() {
	fmt.Println("\n--- MY TASKS ---")
	for _, t := range l.Tasks {
		fmt.Printf("[%d] %s\n", t.ID, t.Text)
	}
}

func main() {
	myList := TodoList{}
	myList.Add("Learn Go syntax")
	myList.Add("Build a CLI app")
	myList.Show()
}
