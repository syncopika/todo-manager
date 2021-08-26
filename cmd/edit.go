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
	"io/ioutil"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a TODO item",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("todo.txt")
		if err != nil {
			fmt.Println("had trouble reading todo.txt!")
		}else{
			lines := string(data)
			tasks := strings.Split(lines, "\n") // only get title of task?
			
			//fmt.Println(strings.Split(lines, "\n"))
			
			prompt := promptui.Select{
				Label: "Select Task",
				Items: tasks,
			}
			
			idx, result, err2 := prompt.Run()

			if err2 != nil {
				fmt.Println("prompt failed!")
			}else{
				// TODO: after selecting task, allow user to edit
				// add another select prompt to ask if they want to edit the title, status (and later description?)
				// then have another prompt for editing. if status, another select prompt is needed.
				fmt.Printf("you picked: " + result + " at index: %d\n", idx)
				
				prompt2 := promptui.Select{
					Label: "What would you like to do",
					Items: []string{"edit task", "edit status", "nevermind"},
				}
				
				_, result2, err3 := prompt2.Run()
				if err3 != nil {
					fmt.Println("ruh roh, prompt2 failed!")
				}else{
					if result2 == "edit task" {
						fmt.Println("edit task")
						
						// prompt user to enter new task
						
					}else if result2 == "edit status" {
						fmt.Println("edit status")
						
						// prompt user to select new status of task
						promptStatus := promptui.Select{
							Label: "Change Status",
							Items: []string{"TODO", "IN PROGRESS", "DONE"},
						}
						
						_, newStatus, err3 := promptStatus.Run()
						if err3 != nil {
							fmt.Println("promptStatus failed!")
						}else{
							fmt.Println("new status: " + newStatus)
						}
					}
				}
				
			}
			
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
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}