package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var taskFile = "task.json"

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
}


func loadTasks()([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, errors.New("couldn't read tasks")
	}
	
	err = json.Unmarshal(file, &tasks)

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(task Task) error {
	fmt.Print(23)
	log.SetFlags(0)
	tasks, err := loadTasks()
	if err != nil {
		log.Fatal(err)
	}
	totTasks := len(tasks)
	task.ID = fmt.Sprintf("%v", totTasks+1)

	tasks = append(tasks, task)

	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(taskFile, file, 0644)
}

func deleteTaskByID(id string) error {
	tasksList, err := loadTasks()

	if err != nil {
		return errors.New("couldn't load the tasks")
	}
	
	taskNo := 1
	updatedTask := []Task{}
	deleteFlag := false
	for _, task:= range tasksList {
		if task.ID != id {
			task.ID = fmt.Sprintf("%v", taskNo)
			updatedTask = append(updatedTask, task)
			taskNo++;
		} else {
			deleteFlag = true
		}
	}

	if !deleteFlag {
		return errors.New("please provide valid task to be deleted")
	}

	file, err := json.MarshalIndent(updatedTask, "", " ") // encoding : the process of putting a sequence of characters (letters, 
	// numbers, punctuation, and certain symbols) into a specialized format for efficient transmission or storage.

	if err != nil {
		return errors.New("couldn't marshal-indent the tasks")
	}

	return os.WriteFile(taskFile, file, 0644)
}
