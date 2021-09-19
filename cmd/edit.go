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

	"github.com/spf13/cobra"
)

func editTask(todo *TodoList, taskField string, taskName string, newValue string){
	if task, exists := (*todo)[taskName]; exists {
		// task will be a copy of the TodoEntry struct at todo[taskName]
		if taskField == "status" {
			task.TaskStatus = newValue
		} else if taskField == "updated" {
			task.TaskUpdatedTimestamp = newValue
		}
		
		// TODO: fill in the other options
		
		// update map with a new updated struct
		(*todo)[taskName] = task
	}
}

func DisplayTasks(filename string){
	var todoList = GetFileContents(filename) // returns map[string]TodoEntry
	
	// we need to pass an int ptr to get back the selected task index
	var tasks = []string{}
	for k, _ := range todoList {
		tasks = append(tasks, k)
	}
	
	var selectedTaskName = SurveyAskOneSelect("Select task", tasks)
	
	var todo = SurveyAskOneSelect(
		"What would you like to do", 
		[]string{"edit task description", "edit status", "remove task", "nevermind"},
	)
	
	if todo == "edit task description" {
		fmt.Println("edit task description")
		
	}else if todo == "edit status" {
		fmt.Println("edit status")
		
		var newStatus = SurveyAskOneSelect(
			"Change Status", 
			[]string{"TODO", "IN PROGRESS", "DONE"},
		)
		
		// add back to the map the new task and update the updated field with current timestamp
		editTask(&todoList, "status", selectedTaskName, newStatus)
		
		currTime := time.Now()
		formattedTime := currTime.Format("Mon Jan 2 15:04:05 MST 2006")
		editTask(&todoList, "updated", selectedTaskName, formattedTime)
		
		dataToWrite, err := json.MarshalIndent(todoList, "", " ")
		HandleError(err, "error marshalling the updated data!")
		
		err2 := ioutil.WriteFile(filename, dataToWrite, 0644)
		HandleError(err2, "error writing the updated data to file!")
	}
}


var SelectedFileNameEdit string

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a TODO item",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := "todo-lists"
	
		if SelectedFileNameEdit != "" {
			// maybe check if the file extension exists in the name already?
			DisplayTasks(fmt.Sprintf("%s/%s.txt", dirPath, SelectedFileNameEdit))
		}else{
			// list the files to choose from and then show tasks to edit
			filepath := SelectFile(dirPath)
			DisplayTasks(filepath)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	editCmd.Flags().StringVarP(&SelectedFileNameEdit, "file", "f", "", "specify the TODO list to list tasks from")
}
