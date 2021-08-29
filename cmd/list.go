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

	"github.com/spf13/cobra"
)


var SelectedFileNameList string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the tasks in the TODO list",
	Long: "list the current tasks in the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: have list be able to accept an argument for how many tasks to list (subcommand?)
		// offer option to show only todo/in-progress tasks or finished tasks? also format the info nicely
		
		var filepath string
		dirPath := "todo-lists"
		
		if SelectedFileNameList == "" {
			filepath = SelectFile(dirPath)
		}else{
			filepath = fmt.Sprintf("%s/%s.txt", dirPath, SelectedFileNameList)
		}
		
		data, err := ioutil.ReadFile(filepath)
		HandleError(err, "had trouble reading todo.txt!")
		
		fmt.Println(string(data))
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
