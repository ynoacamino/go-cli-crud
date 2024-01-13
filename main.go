package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ynoacamino/go-cli-crud/task"
)

func main() {
	file, error := os.OpenFile("task.json", os.O_RDWR|os.O_CREATE, 0666)
	reader := bufio.NewReader(os.Stdin)

	if error != nil {
		fmt.Println("Error:", error)
		panic(error)
	}

	defer file.Close()

	var tasks []task.Task

	info, error := file.Stat()

	if error != nil {
		fmt.Println("Error:", error)
		panic(error)
	}

	if info.Size() > 0 {
		bytes, err := io.ReadAll(file)

		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)

		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)

	case "add":
		fmt.Println("What is your task?")

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(name, tasks)
		task.SaveTask(file, tasks)

	case "complete":
		task.SaveTask(file, task.CompleteTask(tasks))

	case "delete":
		fmt.Println("Enter the id of the task")
		str, _ := reader.ReadString('\n')

		index := task.ParseInt(str)

		task.SaveTask(file, task.DeleteById(index, tasks))

	default:
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Uso de go-cli-crud: [list | add | complete | delete]")
	os.Exit(0)
}
