package cmd

import (
	"io/ioutil"
	"fmt"

	"github.com/manifoldco/promptui"
)

// let the user find the TODO file they want to list the tasks of
func SelectFile(dirPath string) string {
	// get all the files in todo-lists/
	files, err := ioutil.ReadDir(dirPath) // []fs.FileInfo is returned
	
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	
	prompt := promptui.Select{
		Label: "Select the TODO list",
		Items: fileNames,
	}
	
	_, fileToSelect, err := prompt.Run()
	
	if err != nil {
		fmt.Println("ruh roh, TODO file selection failed!")
		return "" // should rather return error?
	}else{
		// return selected filepath
		return fmt.Sprintf("todo-lists/%s", fileToSelect)
	}
}

// general error handler function
func HandleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}