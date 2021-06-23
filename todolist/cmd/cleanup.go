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
	"strings"

	"todolist/helper"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup done tasks",
	Long: `Cleanup done tasks
	sample usage todolist cleanup`,
	Run: func(cmd *cobra.Command, args []string) {
		askBeforeCleaning(fileName)
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

// Check for user confirmation before cleaning up the done tasks
func askBeforeCleaning(fileName string)  {
	fmt.Println("Do you want to cleanup done tasks? Enter y or n")
		var confirmation string
		fmt.Scanln(&confirmation)
		if confirmation == "y" {
			CleanupCompletedTasks(fileName, cleanedUpTasksFile)
		}
}

//This method is used to remove the tasks from todolist if they are marked as done and it also stores the done tasks 
// which will be removed, in a separate file for future reference
func CleanupCompletedTasks(fileName, cleanedUpTasksFile string){
	fileData := helper.ReadFile(fileName)
		filedataArray := strings.Split(fileData, "\n")
		var pendingTasks []string
		var completedTasks []string
		fmt.Println("Cleanup of completed tasks started.....")
		countOfCompletedTasks := 0
		for _,v := range filedataArray {
			if strings.Contains(v, "-> DONE") {
				countOfCompletedTasks +=1
				completedTasks = append(completedTasks,v)
				continue
			}else{
				pendingTasks = append(pendingTasks, v)
			}
		}

		completedTasksLog := strings.Join(completedTasks, "\n")
		helper.AppendToFile(completedTasksLog,cleanedUpTasksFile)
		if countOfCompletedTasks == 0 {
			fmt.Println("Nothing to clean")
			return
		}
		helper.CreateFile(fileName)
				for _, value := range pendingTasks{
					parts := strings.Split(value, ":")
					Add(parts[1],fileName,false)
				}
}