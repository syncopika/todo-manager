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

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// let the user find the TODO file they want to list the tasks of
// TODO: make this function just return the selected file (or an error)
// and put in a utils.go file so it can used elsewhere too
func SelectFile() {
	// get all the files in todo-lists/
	files, err := ioutil.ReadDir("todo-lists") // []fs.FileInfo
	
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
	}else{
		// list all tasks in the selected file
		path := fmt.Sprintf("todo-lists/%s", fileToSelect)
		data, err2 := ioutil.ReadFile(path)
		if err2 != nil {
			fmt.Printf("had trouble reading %s!\n", path)
		}else{
			fmt.Println(string(data))
		}
	}
}

var SelectedFileNameList string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks in the TODO list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: have list be able to accept an argument for how many tasks to list (subcommand?)
		// offer option to show only todo/in-progress tasks or finished tasks?
		// also be able to edit a task? maybe that should be a separate command (or subcommand)
		// also format the info nicely
		
		if SelectedFileNameList == "" {
			SelectFile()
		}else{
			data, err := ioutil.ReadFile(fmt.Sprintf("todo-lists/%s.txt", SelectedFileNameList))
			if err != nil {
				fmt.Println("had trouble reading todo.txt!")
			}else{
				fmt.Println(string(data))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().StringVarP(&SelectedFileNameList, "file", "f", "", "specify the TODO list to list tasks from")
}
