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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays the complete list of tasks and list of all tasks still to done",
	Long: `Displays the complete list of tasks and list of all tasks still to done.
	sample usage todolist list`,
	Run: func(cmd *cobra.Command, args []string) {
		list(fileName)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(fileName string)  {
	fileData := helper.ReadFile(fileName)
	fmt.Println("List of all tasks")
	fmt.Println(fileData)
	filedataarray := strings.Split(fileData, "\n")
	var pendingTasksArray = []string{"===============================", "List of Pending Tasks"}
	for _, v := range filedataarray{
		if strings.Contains(v, "-> DONE") {
			continue
		}else{
			pendingTasksArray = append(pendingTasksArray, v)
		}
	}
	
	pendingTasks := strings.Join(pendingTasksArray, "\n")
	fmt.Println(pendingTasks)
}