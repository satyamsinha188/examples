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
	"log"
	"strings"
	"strconv"

	"todolist/helper"

	"github.com/spf13/cobra"
)

const (
	fileName = "tasklist.txt"
	cleanedUpTasksFile = "cleaneduptasklist.txt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a task to todo list",
	Long: `To add a task to todo list
	sample usage todolist add "task1"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Provide a task to add. For help run -> todolist add --help")
			return
		}
		task := strings.Join(args,"")
		Add(task,fileName,true)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

//Add a task to the provided file. It checks the count of existing task and append the new task to existing task+1 pos
func Add(task, fileName string, confirmBeforeAdding bool)  {
	
	if confirmBeforeAdding {
		userConfirmation := userConfirmation(task)
		if userConfirmation != "y" {
			return
		}
	}
	currentTaskCount := helper.GetExistingTaskCount(fileName)
	if currentTaskCount==0 {
		err := ioutil.WriteFile(fileName, []byte("1:"+task), 0644)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	nextTaskCount := currentTaskCount+1
	newTask:= "\n"+strconv.Itoa(nextTaskCount)+":"+task
	helper.AppendToFile(newTask, fileName)	
}


// Check for user confirmation before adding a tasks to todolist
func userConfirmation(task string) string {
	fmt.Printf(`Do you want to add "%s" in task list? Enter y or n : `, task)
	var confirmation string
	fmt.Scanln(&confirmation)
	return confirmation
}
