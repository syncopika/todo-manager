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


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new item to the TODO list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")
		
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
		if err != nil {
			fmt.Println("prompt failed!")
		}
		
		// check status of task with user
		statusCheckPrompt := promptui.Select{
			Label: "Check TODO Item Status",
			Items: []string{"TODO", "IN PROGRESS", "DONE"},
		}
		
		_, statusRes, err2 := statusCheckPrompt.Run()
		if err2 != nil {
			fmt.Println("status check prompt failed")
		}
		
		// check if file exists? then add (create file if needed?)
		currTime := time.Now()
		formattedTime := currTime.Format("Mon Jan 2 15:04:05 MST 2006")
		
		file, err3 := os.OpenFile("todo.txt", os.O_RDWR|os.O_APPEND, 0644)
		defer file.Close()
		
		if err3 != nil {
			fmt.Println("failed to open todo.txt!")
		}else{
			newTask := taskName + "|" + statusRes + "|" + formattedTime + "\n"
			_, err4 := file.WriteString(newTask)
			if err4 != nil {
				fmt.Println("There was a problem writing to todo.txt!")
			}else{
				fmt.Println("New task added!")
			}
		}
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
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
