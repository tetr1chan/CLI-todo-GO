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
	}
	flag = true
	for flag {
		fmt.Println("Do you want to change Task?")
		var checker1 string
		fmt.Scanln(&checker1)
		if checker1 == "YES" || checker1 == "yes" || checker1 == "Yes" {
			fmt.Println("Enter the number of task you want to change")
			var NumofTask int
			fmt.Scanln(&NumofTask)
			fmt.Println("If you want change Task enter 1 if you want change deadline enter 2 if you want delete task enter 3")
			var changeable int
			fmt.Scanln(&changeable)
			if changeable == 1 {
				fmt.Println("Enter New description of Task")
				var Newdesc string
				fmt.Scanln(&Newdesc)
				AllTasks[NumofTask].WhTask = Newdesc
			} else if changeable == 2 {
				fmt.Println("Enter new deadline")
				var Newdead string
				fmt.Scanln(&Newdead)
				AllTasks[NumofTask].Deadline = Newdead
			} else if changeable == 3 {
				AllTasks = append(AllTasks[:NumofTask-1], AllTasks[NumofTask:]...)
			} else {
				return
			}
		} else {
			flag = false
		}
	}
	fmt.Println("Do you want see new list of Task?")
	var checker2 string
	fmt.Scanln(&checker2)
	if checker2 == "YES" || checker2 == "yes" || checker2 == "Yes" {
		i := 0
		for _, task := range AllTasks {
			i++
			fmt.Print("Task №", i, "\n")
			fmt.Print(task.WhTask)
			fmt.Print(task.Deadline)
			fmt.Println("---------------------------------")
		}
	}
}
