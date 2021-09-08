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
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func editTask(taskString string, field string, newValue string) string {
	selectTaskTokens := strings.Split(taskString, "|")
	
	if field == "status" {
		selectTaskTokens[1] = newValue
	}
	
	return strings.Join(selectTaskTokens, "|")
}

func DisplayTasks(filename string){
	data, err := ioutil.ReadFile(filename)
	HandleError(err, fmt.Sprintf("had trouble reading %s!\n", filename))

	// note: editing a todo file manually causes trouble? something to do with \n
	lines := string(data)
	tasks := strings.Split(lines, "\n")
	
	// we need to pass an int ptr to get back the selected task index
	var taskIndex = SurveyAskOneSelectIndex("Select task", tasks)
	var selectedTask = tasks[taskIndex]
	
	// TODO: after selecting task, allow user to edit
	// add another select prompt to ask if they want to edit the title, status (and later description?)
	// then have another prompt for editing. if status, another select prompt is needed.
	
	var todo = SurveyAskOneSelect(
		"What would you like to do", 
		[]string{"edit task", "edit status", "remove task", "nevermind"},
	)
	
	if todo == "edit task" {
		fmt.Println("edit task")
		
		// prompt user to enter new task
		
	}else if todo == "edit status" {
		fmt.Println("edit status")
		
		var newStatus = SurveyAskOneSelect(
			"Change Status", 
			[]string{"TODO", "IN PROGRESS", "DONE"},
		)
		
		// TODO: we're making an assumption about how
		// we're storing the task info in the TODO list!
		// change later? use a struct?
		editedTask := editTask(selectedTask, "status", newStatus)
		
		// replace the task at the right index of tasks from above
		tasks[taskIndex] = editedTask
		
		// then rejoin tasks as a single string and rewrite to file
		revisedTasks := strings.Join(tasks, "\n")
		
		file, openFileErr := os.OpenFile(filename, os.O_RDWR, 0644)
		defer file.Close()
		HandleError(openFileErr, fmt.Sprintf("failed to open %s!\n", filename))
		
		_, writeFileError := file.WriteString(revisedTasks)
		HandleError(writeFileError, fmt.Sprintf("There was a problem writing to %s!\n", filename))
		fmt.Println("task was revised!")
		
		// TODO - add a another field of the TODO list for recording the timestamp of when the task was modified?
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
