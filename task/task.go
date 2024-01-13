package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask(list []Task) {
	if len(list) < 1 {
		fmt.Println("There aren't taks")
		return
	}

	for _, task := range list {

		status := " "

		if task.Complete {
			status = "[✓] Complete\t"
		} else {
			status = "[ ] Uncomplete\t"
		}

		fmt.Printf("%d %s %s \n", task.ID, status, task.Name)
	}
}

func AddTask(name string, list []Task) []Task {
	newTask := Task{
		Name:     name,
		ID:       len(list) + 1,
		Complete: false,
	}
	return append(list, newTask)
}

func SaveTask(file *os.File, list []Task) {
	bytes, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

func CompleteTask(list []Task) []Task {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What task is complete?")

	ListTask(list)

	fmt.Println("Choose a number")
	num, _ := reader.ReadString('\n')

	index := ParseInt(num)

	for i, tarea := range list {
		if tarea.ID == index {
			list[i].Complete = true
			break
		}
	}
	return list
}

func DeleteById(id int, list []Task) []Task {
	var index int

	for i, tarea := range list {
		if tarea.ID == id {
			index = i
		}
	}

	list = append(list[:index], list[index+1:]...)
	return list
}

func ParseInt(str string) int {
	num, err := strconv.Atoi(ClearString(str))

	if err != nil {
		panic(err)
	}

	return num
}

func ClearString(str string) string {
	return strings.Trim(str, "\r\n")
}
