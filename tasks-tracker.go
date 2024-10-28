package main

import (
	"os"
	"github.com/Franceskynov/go-tasks-tracker/actions"
)

func main() {
	const fileName string = "tasks.json"
	
	numberOfArguments := len(os.Args)
	
	if numberOfArguments == 1 {
		actions.Welcome()		
	} else if numberOfArguments >= 2 {
		switch os.Args[1] {

			case "add":
				actions.AddNewTask(os.Args, fileName)
			case "update":
				actions.UpdateExistingTask(os.Args, fileName, "updateName")
			case "mark-done":
				actions.UpdateExistingTask(os.Args, fileName, "updateStatusDone")
			case "mark-in-progress":
				actions.UpdateExistingTask(os.Args, fileName, "updateStatusInProgress")
			case "delete":
				actions.DeleteATask(os.Args, fileName)
			case "list":

				if numberOfArguments == 3 {
				
					switch os.Args[2] {
						case "todo":
							actions.ListAllTasks(fileName, "todo")
						case "in-progress":
							actions.ListAllTasks(fileName, "inProgress")
						case "done":
							actions.ListAllTasks(fileName, "done")
						default:
							break
					}

				} else {
					actions.ListAllTasks(fileName, "all")	
				}


			default:
				break
			
		}
	} 

}
