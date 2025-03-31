package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	fmt.Println("welcome to my editor challenge! commands : exit,list,delete,add,save,edit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		commands := strings.Fields(input)
		if len(commands) == 0 {
			fmt.Println("please enter a command!!")
			continue
		}
		switch commands[0] {
		case "add":
			if len(commands) < 2 {
				fmt.Println("please enter a text for command!!")
				continue
			}
			text := strings.Join(commands[1:], " ")
			lines = append(lines, text)
			fmt.Println("text is add: ", text)

		case "delete":
			if len(lines) == 0 {
				fmt.Println("not exist line for remove!")
				continue
			}
			if len(commands) < 2 {
				lines = lines[:len(lines)-1] // remove last line by default
				fmt.Println("last line was deleted!")
			} else {
				lineNum, err := strconv.Atoi(commands[1])
				if err != nil || lineNum < 1 || lineNum > len(lines) {
					fmt.Println("number of line is Invalid!")
					continue
				}
				lines = append(lines[:lineNum-1], lines[lineNum:]...)
				fmt.Printf("line %d removed\n", lineNum)
			}

		case "edit":
			if len(lines) == 0 {
				fmt.Println("There is no text to edit!!")
				continue
			}
			if len(commands) < 3 {
				fmt.Println("wrong command! Format : edit [line Number] [new text]")
				continue
			}
			lineNum, err := strconv.Atoi(commands[1])
			if err != nil || lineNum < 1 || lineNum > len(lines) {
				fmt.Println("number of line is Invalid!")
				continue
			}
			newText := strings.Join(commands[2:], " ")
			lines[lineNum-1] = newText
			fmt.Printf("line %d edited : %s\n", lineNum, newText)

		case "list":
			if len(lines) == 0 {
				fmt.Println("list is empty!")
			} else {
				fmt.Println("Corrent content:")
				for i, line := range lines {
					fmt.Printf("%d: %s\n", i+1, line)
				}
			}

		case "save":
			if len(lines) == 0 {
				fmt.Println("There is text to save!")
				continue
			}
			file, err := os.Create("save.txt")
			if err != nil {
				fmt.Println("Error in create output file: ", err)
				continue
			}
			defer file.Close()
			for _, line := range lines {
				_, err := file.WriteString(line + "\n")
				if err != nil {
					fmt.Println("Error in text file : ", err)
					continue
				}
			}
			fmt.Println("Content saved in file txt")

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("command invalid!")
		}
	}
}
