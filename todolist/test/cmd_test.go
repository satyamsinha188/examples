package test

import (
	"testing"
	"os"
	"log"
	"todolist/helper"
	"todolist/cmd"
)

const (
	done = iota
	undone
)

const (
	fileName 			= "inputfiles/test_cleanup.txt"
	cleanedUpTasksFile	 = "inputfiles/test_cleaned_up_tasks.txt"
)

/* This test is used to add tasks in the file, mark few tasks as done, mark few done tasks as undone and then calling 
cleanup command. It will validate if the tasks count are as expected
*/
func TestCleanupCommands(t *testing.T){
	//check if file exists
	helper.CheckIfFileExists(fileName)
	helper.CheckIfFileExists(cleanedUpTasksFile)
	truncateFile(fileName)

	//add tasks to the file
	cmd.Add("task1", fileName, false)
	cmd.Add("task2", fileName, false)
	cmd.Add("task3", fileName, false)
	cmd.Add("task4", fileName, false)

	taskCount := helper.GetExistingTaskCount(fileName)

	if taskCount != 4{
		t.Fail()
	}

	//making task1,task2 and task4 as done
	helper.MarkOrUnmarkTask(fileName, 1, done)
	helper.MarkOrUnmarkTask(fileName, 2, done)
	helper.MarkOrUnmarkTask(fileName, 4, done)

	//making task2 as undone
	helper.MarkOrUnmarkTask(fileName, 2, undone)

	//call CleanupCompletedTasks
	cmd.CleanupCompletedTasks(fileName,cleanedUpTasksFile)

	//get the task count after cleanup
	taskCount = helper.GetExistingTaskCount(fileName)

	//truncate the file
	truncateFile(fileName)
	truncateFile(cleanedUpTasksFile)

	if taskCount !=2 {
		t.Fail()
	}
}

func truncateFile(fileName string){
	err := os.Truncate(fileName, 0)
	if err != nil {
		log.Fatal(err)
	}
}
