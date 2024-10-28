# Go Tasks Tracker

A CLI tool for track all of your daily tasks.


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