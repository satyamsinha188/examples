package test

import (
	"testing"
	"todolist/helper"
)

const testFileName = "inputfiles/test_read_file.txt"
/* 
*/
func TestReadFile(t *testing.T){
	helper.CheckIfFileExists(testFileName)
	helper.WriteToFile(testFileName, "This is the sample test file")
	fileContent := helper.ReadFile(testFileName)

	if fileContent != "This is the sample test file"{
		t.Fail()
	}

	truncateFile(testFileName)
}

func TestCheckIfDone(t *testing.T){
	testCases := []struct {
        task  		string
        expected  	bool
    }{
		{"schedule a meeting-> DONE", true},
		{"document handover", false},
		{"document writing Done", false},
	}

	for _, v := range testCases{
		if v.expected != helper.CheckIfDone(v.task){
			t.Fail()
		}
	}
}