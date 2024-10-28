# Go Tasks Tracker

A CLI tool for track all of your daily tasks.
Project coded follow the requirements and constraints of Roadmap.sh [Task tracker](https://roadmap.sh/projects/task-tracker)

## Building the tool
```
go build .
```


## Listing

Listing all tasks
```
./go-tasks-tracker list  
```

Listing tasks by status
```
./go-tasks-tracker in-progress

./go-tasks-tracker todo

./go-tasks-tracker list done
```

## Running the tool

Adding a new task
```
./go-tasks-tracker add "paint the house"
```

## Updating and deleting

Delete a task
```
./go-tasks-tracker delete 1002   
```

Update a task
```
./go-tasks-tracker update 999 "paint the house"
```


## Marking a task

Mark a task as in progress
```
./go-tasks-tracker mark-in-progress 1000
```

Mark a taks as done
```
 ./go-tasks-tracker mark-done 1000
```