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
	"os"
	fp "path/filepath"

	"github.com/spf13/cobra"
)

func createFile(filepath string) {
	err := os.MkdirAll(fp.Dir(filepath), 0770)
	if err != nil {
		fmt.Printf("oh no, %s was not able to be created!", fp.Dir(filepath)) 
	}else{
		file, err2 := os.Create(filepath)
		defer file.Close()
		if err2 != nil {
			fmt.Printf("oh no, something went wrong! %s was unable to be created.\n", filepath)
		}else{
			fmt.Println("new TODO file created!")
		}
	}
}

// name of a TODO list to create if passed in as an arg to init
var TodoFileName string

// initCmd represents the serve command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init command",
	Long: "create the todo note",
	Run: func(cmd *cobra.Command, args []string) {
		// create a new text file for holding TODO tasks
		// but check if one exists already
		// all TODO lists should go in the todo-lists directory
		
		var filepath = "todo-lists/"
		if TodoFileName != "" {
			filepath += TodoFileName + ".json"
		}else{
			filepath += "todo.json"
		}
		
		_, err := os.Stat(filepath)
		
		if err == nil {
			// file already exists
			fmt.Printf("%s already exists!", filepath)
		}else{
			createFile(filepath)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringVarP(&TodoFileName, "create", "c", "", "create a new TODO list")
}
