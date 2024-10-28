package main

import (
	"log"
	"os"

	"github.com/Franceskynov/go-tasks-tracker/actions"

	"github.com/urfave/cli/v2"
)

func main() {
	const fileName string = "tasks.json"
	app := &cli.App{
		Name:  "Go Tasks Tracker",
		Usage: "Is a simple CLI for tracking all of your daily tasks",
		Action: func(cCtx *cli.Context) error {
			return actions.Welcome()
		},
		Version: "v0.1",
		Authors: []*cli.Author{
			{
				Name:  "Franceskynov",
				Email: "franceskynov@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "add",
				Aliases:     []string{"Add", "ADD"},
				Category:    "Add new task",
				Usage:       `add "task to do"`,
				Description: "This command is used to add a new task to the todo list",
				Action: func(ctx *cli.Context) error {
					return actions.AddNewTask(ctx, fileName)
				},
			},
			{
				Name:        "update",
				Aliases:     []string{"Update", "UPDATE"},
				Category:    "Update an existing task",
				Usage:       `update 1 "Read a book"`,
				Description: "This command is used to update an existing task",
				Action: func(ctx *cli.Context) error {
					return actions.UpdateExistingTask(ctx, fileName, "updateName")
				},
			},
			{
				Name:        "list",
				Aliases:     []string{"List", "LIST"},
				Category:    "List all the existing tasks",
				Usage:       `list`,
				Description: "This command is used to list all existing tasks",
				Action: func(ctx *cli.Context) error {
					return actions.ListAllTasks(ctx, fileName, "all")
				},
				Subcommands: []*cli.Command{
					{
						Name: "done",
						Usage: "",
						Action: func(ctx *cli.Context) error {
							return actions.ListAllTasks(ctx, fileName, "done")
						},
					},
					{
						Name: "todo",
						Usage: "",
						Action: func(ctx *cli.Context) error {
							return actions.ListAllTasks(ctx, fileName, "todo")
						},
					},
					{
						Name: "in-progress",
						Usage: "",
						Action: func(ctx *cli.Context) error {
							return actions.ListAllTasks(ctx, fileName, "inProgress")
						},
					},
				},
			},
			{
				Name:        "delete",
				Aliases:     []string{"Delete", "DELETE"},
				Category:    "Delete a task",
				Usage:       `delete ID`,
				Description: "This command is used to delete an existing task",
				Action: func(ctx *cli.Context) error {
					return actions.DeleteATask(ctx, fileName)
				},
			},
			{
				Name:        "mark-in-progress",
				Aliases:     []string{"Mark-In-Progress", "MARK-IN-PROGRESS"},
				Category:    "Mark in progress a task",
				Usage:       `mark-in-progress ID`,
				Description: "This command is used to mark in progress a task",
				Action: func(ctx *cli.Context) error {
					return actions.UpdateExistingTask(ctx, fileName, "updateStatusInProgress")
				},
			},
			{
				Name:        "mark-done",
				Aliases:     []string{"Mark-Done", "MARK-DONE"},
				Category:    "Mark as done a task",
				Usage:       `done ID`,
				Description: "This command is used to mark as done a task",
				Action: func(ctx *cli.Context) error {
					return actions.UpdateExistingTask(ctx, fileName, "updateStatusDone")
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
