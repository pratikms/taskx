# taskx

Taskx is a CLI task manager to manage all your TODOs from the comfort of your CLI

## Installation

`$ go get -u github.com/pratikms/taskx`

## Usage

```
$ taskx help
Taskx is a CLI task manager to manage all your TODOs from the comfort of your CLI

Usage:
  taskx [command]

Available Commands:
  add         Adds a task to your task list
  do          Marks a task as complete
  help        Help about any command
  list        Lists all of your tasks

Flags:
  -h, --help   help for taskx

Use "taskx [command] --help" for more information about a command.
```

To add a task:
```
$ taskx add Check if Hello, world! is added as a task
$ taskx add Talk the talk
$ taskx add Walk the walk
```

To list all tasks:
```
$ taskx list
You have the following tasks: 
1. Check if Hello, world! is added as a task
2. Talk the talk
3. Walk the walk
```

To mark a task as done:
```
$ taskx do 1
Marked "Check if Hello, world! is added as a task" as completed
```