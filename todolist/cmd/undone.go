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
	"strconv"
	
	"todolist/helper"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark task as undone",
	Long: `Mark task as undone
	sample usage todolist undone taskid
	For example to mark taskid 1 as undone -> todolist undone 1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if	err != nil{
			fmt.Printf("Unable to convert %s to int. Please provide a valid integer value. Error: %s", args[0], err)
			return
		}
		helper.MarkOrUnmarkTask(fileName,id,undone)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}
