An application for managing to-do task lists.
The app provides below commands

commands:
● add – Add a task to the list. Command argument should be the task name.
● done – Mark a task as done. Command argument should be task ID, starting from 1.
● undone – Mark task as not done. Command argument should be task ID, starting from 1.
● list – List all the tasks and that not been done.
● cleanup – Remove from the store all tasks marked as done.

Example usage of the app could be:

$ todolist help
Usage:
  todolist [command]
Available Commands:
  add Add task to the list
  cleanup Cleanup done tasks
  done Mark task as done
  help Help about any command
  list List all tasks still to do
  undone Mark task as not done
  
Flags:
  -h, --help help for todolist
  
Use "todolist [command] --help" for more information about a command.

Adding tasks:

$ todolist add "Implement command list"
$ todolist add "Implement task2"
