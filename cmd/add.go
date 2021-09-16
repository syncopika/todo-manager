/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var SelectedFileNameAdd string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new item to the TODO list",
	Long: "add a new item to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
	
		var filepath string
		dirPath := "todo-lists"
		
		// select TODO file to add to
		if SelectedFileNameAdd == "" {
			filepath = SelectFile(dirPath)
		}else{
			filepath = fmt.Sprintf("%s/%s.json", dirPath, SelectedFileNameList)
		}
		
		// ask for name of new todo item
		var taskName = ""
		getTodoItemName := &survey.Input{ Message: "new TODO item name: " }
		survey.AskOne(getTodoItemName, &taskName, survey.WithValidator(survey.MinLength(1)))

		// ask for the description of the todo item
		var taskDescription = ""
		getTodoItemDescription := &survey.Input{ Message: "new TODO item description: " }
		survey.AskOne(getTodoItemDescription, &taskDescription, survey.WithValidator(survey.MinLength(1)))
		
		// ask for todo item status
		var statusRes = SurveyAskOneSelect(
			"Check TODO Item Status",
			[]string{"TODO", "IN PROGRESS", "DONE"},
		)
		
		currTime := time.Now()
		formattedTime := currTime.Format("Mon Jan 2 15:04:05 MST 2006")

		// build the todo entry
		newEntry := TodoEntry{
			TaskName:             taskName,
			TaskDescription:      taskDescription,
			TaskStatus:           statusRes,
			TaskCreatedTimestamp: formattedTime,
			TaskUpdatedTimestamp: "",
		}
		
		// get curr json data and add to it
		var currTodoList = GetFileContents(filepath)
		
		if val, exists := currTodoList[taskName]; exists {
			// value already exists
			// TODO: allow user to overwrite?
			fmt.Println("%s already exists!", taskName);
		}else{
			currTodoList[taskName] = newEntry
		}
		
		dataToWrite, err := json.MarshalIndent(currTodoList, "", " ")
		HandleError(err, "error marshalling the new data!")
		
		err2 := ioutil.WriteFile(filepath, dataToWrite, 0644)
		HandleError(err2, "error writing the new data to file!")
		
		fmt.Println("New task added!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringVarP(&SelectedFileNameList, "file", "f", "", "specify the TODO list to add tasks to")
}
