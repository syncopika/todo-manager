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
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
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
			filepath = fmt.Sprintf("%s/%s.txt", dirPath, SelectedFileNameList)
		}
		
		validateAdd := func(input string) error {
			if len(strings.TrimSpace(input)) <= 0 {
				return errors.New("you didn't enter anything!")
			}
			return nil
		}
		
		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}
		
		// add new task
		addPrompt := promptui.Prompt{
			Label: "Add TODO Item",
			Templates: templates,
			Validate: validateAdd,
		}
		
		taskName, err := addPrompt.Run()
		HandleError(err, "prompt failed!")
		
		// check status of task with user
		statusCheckPrompt := promptui.Select{
			Label: "Check TODO Item Status",
			Items: []string{"TODO", "IN PROGRESS", "DONE"},
		}
		
		_, statusRes, err2 := statusCheckPrompt.Run()
		HandleError(err2, "status check prompt failed")
		
		// check if file exists? then add (create file if needed?)
		currTime := time.Now()
		formattedTime := currTime.Format("Mon Jan 2 15:04:05 MST 2006")
		
		file, err3 := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0644)
		defer file.Close()
		HandleError(err3, "failed to open todo.txt!")
		
		newTask := taskName + "|" + statusRes + "|" + formattedTime + "\n"
		_, err4 := file.WriteString(newTask)
		HandleError(err4, "There was a problem writing to todo.txt!")
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
