cat > README.md << 'EOF'
# Task Tracker CLI

A simple command-line task tracker written in Go. Tasks are stored in a JSON file in your home directory, so they persist between runs.

## Features

- Add tasks (default status: `todo`)
- List all tasks, or filter by status (`done`, `in-progress`)
- Change task status (`done`, `in-progress`)
- Edit task description
- Delete tasks
- Automatic saving to `~/.task-tracker/tasks.json`
- Creation and last update timestamps
- Built-in help

## Installation

1. Clone the repository:
```bash
   git clone https://github.com/yourusername/task-tracker-cli.git
   cd task-tracker-cli
```
2. Build the binary:

```bash
go build -o task-tracker
```

## Usage
```bash
./task-tracker <command> [arguments]
```
## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `add <description>` | Add a new task | `./task-tracker add "Buy milk"` |
| `list` | Show all tasks | `./task-tracker list` |
| `list done` | Show completed tasks | `./task-tracker list done` |
| `list in-progress` | Show tasks in progress | `./task-tracker list in-progress` |
| `done <id>` | Mark a task as done | `./task-tracker done 1` |
| `in-progress <id>` | Mark a task as in-progress | `./task-tracker in-progress 2` |
| `update <id> <new description>` | Update task description | `./task-tracker update 1 "Buy bread"` |
| `delete <id>` | Delete a task | `./task-tracker delete 1` |
| `help` | Show this help message | `./task-tracker help` |

## Example session
```text
$ ./task-tracker add "Learn Go"
Task added with ID: 1

$ ./task-tracker list
ID: 1; Task: Learn Go; Status: todo; Created: 24.03.2026 15:30; Updated: 24.03.2026 15:30

$ ./task-tracker in-progress 1
Id: 1. Task status changed to "In-Progress".

$ ./task-tracker done 1
Id: 1. Task status changed to "Done".

$ ./task-tracker list done
ID: 1; Task: Learn Go; Status: done; Created: 24.03.2026 15:30; Updated: 24.03.2026 15:32
```

## Data storage
All tasks are stored in ~/.task-tracker/tasks.json in JSON format. Example:
```json
[
    {
        "ID": 1,
        "Description": "Learn Go",
        "Status": "done",
        "CreatedAt": "2026-03-24T15:30:00+03:00",
        "UpdatedAt": "2026-03-24T15:32:00+03:00"
    }
]
```
