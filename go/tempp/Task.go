package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_completed"`
}

func (t *Task) WriteTaskToFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(t)
}

func (t *Task) LoadTaskFromFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(t)
}

func WriteContent(fileName, content string) error {
	tmp := fmt.Sprintf("%s.tmp.%d", fileName, rand.Intn(100000))
	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err := io.Copy(file, strings.NewReader(content)); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	return os.Rename(tmp, fileName)
}

func FileWriteDemo() {
	fmt.Println("Enter file Name: ")
	rdr := bufio.NewReader(os.Stdin)
	fileName, err := rdr.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter content: ")
	content, err := rdr.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if err := WriteContent(fileName, content); err != nil {
		panic(err)
	}
}

func TaskDemo() {
	// Create a sample task
	task := Task{
		Id:          1,
		Title:       "Sample Task",
		Description: "This is a sample task",
		IsComplete:  false,
	}

	// Write the task to file
	err := task.WriteTaskToFile("tasks.json")
	if err != nil {
		panic(err)
	}

	// Create a new task to load into
	var loadedTask Task
	err = loadedTask.LoadTaskFromFile("tasks.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded task: %+v\n", loadedTask)
}
