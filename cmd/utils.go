package cmd

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// define our types for the TODO list JSON
type TodoList map[string]TodoEntry

type TodoEntry struct {
	TaskName                string  `json:"name"`
	TaskDescription         string  `json:"description"`
	TaskStatus              string  `json:"status"`
	TaskCreatedTimestamp	string  `json:"created"`
	TaskUpdatedTimestamp	string  `json:"updated"`
}

func GetFileContents(filepath string) TodoList {
	file, err := ioutil.ReadFile(filepath)
	HandleError(err, fmt.Sprintf("failed to open %s!", filepath))
	
	var todoList TodoList
	
	// check if file is an empty array
	if len(file) == 0 {
		return TodoList{}
	}else{
		json.Unmarshal(file, &todoList)
		return todoList
	}
}

// let the user find the TODO file they want to list the tasks of
func SelectFile(dirPath string) string {
	// get all the files in todo-lists/
	files, err := ioutil.ReadDir(dirPath) // []fs.FileInfo is returned
	if err != nil {
		fmt.Println("error reading todo list directory!")
		return ""
	}
	
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	
	fileToSelect := SurveyAskOneSelect("Select the TODO list", fileNames)
	return fmt.Sprintf("todo-lists/%s", fileToSelect)
}

// wrapper around survey's select and askone for string answers
// survey's askone function can accept a string pointer for storing the selected string;
// if an integer pointer is passed, you get back the selected string's index in options.
func SurveyAskOneSelect(message string, options []string) string {
	answer := ""
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	survey.AskOne(prompt, &answer)
	return answer
}

// like above but for getting the index of the selected item
func SurveyAskOneSelectIndex(message string, options []string) int {
	answer := 0
	prompt := & survey.Select{
		Message: message,
		Options: options,
	}
	survey.AskOne(prompt, &answer)
	return answer
}

// general error handler function
func HandleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}