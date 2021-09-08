package cmd

import (
	"io/ioutil"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	//"github.com/manifoldco/promptui"
)


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
	
	var answer = struct {
		Select string
	}{}
	
	
	qs := []*survey.Question{
		{
			Name: "select",
			Prompt: &survey.Select{
				Message: "Select the TODO list",
				Options: fileNames,
				Default: "todo.txt",
			},
		},
	}
	
	err2 := survey.Ask(qs, &answer)
	
	
	if err2 != nil {
		fmt.Println(err2.Error())
		fmt.Println("ruh roh, TODO file selection failed!")
		return "" // should rather return error?
	}else{
		// return selected filepath
		fileToSelect := answer.Select
		return fmt.Sprintf("todo-lists/%s", fileToSelect)
	}
}

// general error handler function
func HandleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}