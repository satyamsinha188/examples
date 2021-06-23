package helper

import (
    "io/ioutil"
	"log"
	"os"
	"fmt"
	"strings"
	"bufio"
)

//read the content of a given file
func ReadFile(fileName string) string{
	content, err := ioutil.ReadFile(fileName)

		if err != nil {
			log.Fatal(err)
		}

		return string(content)
}

//create a file
func CreateFile(fileName string){
	updatedfile, e := os.Create(fileName)
    if e != nil {
        log.Fatal(e)
    }
    updatedfile.Close()
}

//write to a file
func WriteToFile(fileName, filedata string){
	err := ioutil.WriteFile(fileName, []byte(filedata), 0644)
		if err != nil {
			log.Fatal(err)
		}
		return
}

//Check if a file exists, if not create it
func CheckIfFileExists(name string)  {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
			fmt.Println("the file doesn't exist, creating the file....")
            CreateFile(name)	
        }
    }
}

//MarkOrUnmarkTask is used to mark a task as done by appending "-> DONE" to the task
// or undone a already done tasks by removing "-> DONE"
func MarkOrUnmarkTask(fileName string, id int, taskType TaskType){

	filedata := ReadFile(fileName)
	filedataarray := strings.Split(filedata, "\n")
	if filedata == "" {
		fmt.Println("There are no available tasks in file")
		return
	}else if id > len(filedataarray) {
		fmt.Printf("The task id %d is not available. Run todolist list to get the list of available ids", id)
		return
	}
    postitiontomark := id-1
	condition := CheckIfDone(filedataarray[postitiontomark])
	switch taskType {
	case done:
		if condition {
			fmt.Println("task is already done")
			return
		} else {
			filedataarray[postitiontomark] = filedataarray[postitiontomark] + "-> DONE"
		}	
	case undone:
		if condition {
			filedataarray[postitiontomark] = strings.Replace(filedataarray[postitiontomark], "-> DONE", "", -1)
		} else {
			fmt.Printf("task %d is not done yet.", id)
			return
		}	
	default:
		fmt.Println("not a valid task type")	
	}
	updatedfiledata := strings.Join(filedataarray, "\n")
	CreateFile(fileName)
	WriteToFile(fileName,updatedfiledata)
}

//CheckIfDone checks if a task is done or not. If done it returns true else false
func CheckIfDone(task string) bool {
	if strings.Contains(task, "-> DONE") {
		return true
	}
	return false
}

//GetExistingTaskCount gets the count of tasks in the file
func GetExistingTaskCount(fileName string) int{

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}	
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio. ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return count
}

//Append the given string to a existing file
func AppendToFile(newTask, fileName string){

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer file.Close()
    if _, err := file.WriteString(newTask); err != nil {
        log.Fatal(err)
    }
}