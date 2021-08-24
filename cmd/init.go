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

	"github.com/spf13/cobra"
)

// initCmd represents the serve command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init command",
	Long: "create the todo note",
	Run: func(cmd *cobra.Command, args []string) {
		// create a new text file for holding TODO tasks
		// but check if one exists already
		_, err := os.Stat("todo.txt")
		if err == nil {
			// file already exists
			fmt.Println("todo.txt already exists!")
		}else{
			file, err := os.Create("todo.txt")
			defer file.Close()
			if err != nil {
				fmt.Println("oh no, something went wrong! todo.txt was unable to be created.")
			}else{
				fmt.Println("new TODO file created!")
			}
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
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
