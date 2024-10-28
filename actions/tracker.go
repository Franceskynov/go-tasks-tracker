package actions

import (
	"errors"
	"fmt"
	"github.com/Franceskynov/go-tasks-tracker/utils"
	"strconv"
)

func Welcome() error {
	
	art := `
██████  ████████ ████████ 
██         ██       ██    
██   ███   ██       ██    
██    ██   ██       ██    
 ██████    ██       ██ 
`
	fmt.Println(art)
	fmt.Println("Welcome to Go Tasks Tracker")
   
	return nil
}

func WriteTasks(fileName string, tasks utils.Tasks, id int, name string) error {
	
	// Create the new taks with todo status
	currentTask := utils.Task{
		ID:     id,
		Name:   name,
		Status: "ToDo",
	}

	// Add the current task to the list of existing tasks
	tasks.Tasks = append(tasks.Tasks, currentTask)

	// Write the new changes
	return utils.JSONWriter(fileName, tasks) 
}

/*
	Add a new task
*/
func AddNewTask(ctx []string, fileName string) error {
	
	id := 1
	taskName := ctx[2]
	tasks, err := utils.JSONReader(fileName)
	if taskName == "" {
		return errors.New("please write a valid task")
	}

	if err != nil {

		// fmt.Printf("There's no existing %s, we are going to create it\n", fileName)
		fmt.Printf("Task added successfully %d", id)
		return WriteTasks(fileName, tasks, id, taskName)
	} else {
		numberOfTasks := len(tasks.Tasks)
		
		// Checks is there are tasks, for increase the id  
		if numberOfTasks > 0 {
			id = tasks.Tasks[numberOfTasks-1].ID + 1
		}

		fmt.Printf("Task added successfully %d", id)
		return WriteTasks(fileName, tasks, id, taskName)
	}
}

/*
	Update an existing task
*/
func UpdateExistingTask(ctx []string, fileName string, mode string) error {

	tasks, err := utils.JSONReader(fileName)
	taskID := ctx[2]
	taskName := ""

	if len(ctx) == 4 {
		taskName = ctx[3]
	}
	
	if taskID == "" || (mode == "updateName" && taskName == "")  {
		return errors.New("please write a valid arguments to update the task")
	}

	if err != nil {
		return fmt.Errorf("there's no %s to update the task: %s, %s", fileName, taskID, taskName)
	} else {

		id, err := strconv.Atoi(taskID)
		seledTasks := tasks.Tasks

		if err != nil {
			return errors.New("invalid ID argument")
		}

		idx := utils.FindIndex(seledTasks, id)

		if idx == -1 {
			return fmt.Errorf("the id %d of the task not exist, please write a existant id after list all the tasks", id)
		}

		selectedTask := seledTasks[idx]
		
		switch mode {
			case "updateName":
				selectedTask.Name = taskName
			case "updateStatusDone":
				selectedTask.Status = "Done"
			case "updateStatusInProgress":
				selectedTask.Status = "In Progress"
			default:
				break
		}

		tasks.Tasks[idx] = selectedTask
		fmt.Println("Task updated")
		return utils.JSONWriter(fileName, tasks)
	}
}

/*
List all existing tasks
*/
func ListAllTasks(fileName string, mode string) error {

	decoration := "-----------------------------------------------" 
	fmt.Println(decoration)
	tasks, err := utils.JSONReader(fileName)
	if err != nil {
		return errors.New("there's no file to open")
	}

	if len(tasks.Tasks) == 0 {
		return errors.New("there are no tasks for list")
	} else {

		// Iterate the list of existing tasks
		for _, v := range tasks.Tasks {

			tpl := fmt.Sprintf("Task details: \nId %d, \nname %s \nstatus: %s\n", v.ID, v.Name, v.Status)
			if mode == "all" {
				fmt.Print(tpl)
				fmt.Println(decoration)
			} else if mode == "todo" && v.Status == "ToDo" {
				fmt.Print(tpl)
				fmt.Println(decoration)
			} else if mode == "inProgress" && v.Status == "In Progress"{
				fmt.Print(tpl)
				fmt.Println(decoration)
			} else if mode == "done" && v.Status == "Done" {
				fmt.Print(tpl)
				fmt.Println(decoration)
			}
		}
	}

	return nil
}

/*
Delete a existing task
*/
func DeleteATask(ctx []string, fileName string) error {
	
	taskID := ctx[2]
	tasks, err := utils.JSONReader(fileName)

	if taskID == ""  {
		return errors.New("please write a valid argument to delete the task")
	}

	if err != nil {
		return fmt.Errorf("there's no %s to delete the task: %s", fileName, taskID)
	} else {
		id, err := strconv.Atoi(taskID)

		seledTasks := tasks.Tasks
		if err != nil {
			return errors.New("invalid ID argument")
		}

		idx := utils.FindIndex(seledTasks, id)

		if idx == -1 {
			return fmt.Errorf("the id %d of the task not exist, please write a existant id after list all the tasks", id)
		}

		tasks.Tasks = append(seledTasks[:idx], seledTasks[idx+1:]...)

		utils.JSONWriter(fileName, tasks)
		
	}

	return nil
}
