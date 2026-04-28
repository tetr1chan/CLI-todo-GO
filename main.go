package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	WhTask   string
	Deadline string
}

func (a *Task) NewTask(task string, deadline string) {
	a.WhTask = task
	a.Deadline = deadline
}
func main() {
	flag := true
	var AllTasks []Task
	for flag {
		var s string
		fmt.Println("Do you have new task answer yes/no:")
		fmt.Scanln(&s)
		if s == "Yes" || s == "yes" {
			fmt.Println("Enter your new task:")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error of reading:", err)
				return
			}
			Newt := input
			fmt.Println("Enter deadline of your new task:")
			input1, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error of reading:", err)
				return
			}
			NewD := input1
			var tusk Task
			tusk.NewTask(Newt, NewD)
			AllTasks = append(AllTasks, tusk)
		} else {
			flag = false
		}
	}
	fmt.Println("Do you want check all tasks?")
	var checker string
	fmt.Scanln(&checker)
	if checker == "Yes" || checker == "yes" || checker == "YES" {
		i := 0
		for _, task := range AllTasks {
			i++
			fmt.Print("Task №", i, "\n")
			fmt.Print(task.WhTask)
			fmt.Print(task.Deadline)
			fmt.Println("---------------------------------")
		}
	} else {
		return
	}
}
