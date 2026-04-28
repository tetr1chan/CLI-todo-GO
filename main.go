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

func main() {
	flag := true
	for flag != false {
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
			NewTask := input
			fmt.Println("Enter deadline of your new task:")
			input1, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error of reading:", err)
				return
			}
			NewDeadline := input1
		}
	}
}
