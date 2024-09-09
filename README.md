# CLI-ToDo Application

A simple and lightweight command-line ToDo application built with Go and SQLite. It helps you manage your daily tasks efficiently through commands such as adding, listing, completing, deleting, and clearing tasks.

## Features

- **Add tasks**: Add new tasks to your to-do list.
- **List tasks**: View all pending and completed tasks.
- **Complete tasks**: Mark tasks as completed.
- **Delete tasks**: Remove tasks from the list.
- **Clear all tasks**: Clear the entire list of tasks.

## Prerequisites

- **Go** version 1.16 or higher installed on your system.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/vinij-03/CLI-ToDo.git
   ```
2. Navigate to the project directory:
   ```bash
   cd CLI-ToDo
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build -o todo cmd/todo/main.go
   ```

## Usage

Once you've built the application, you can start managing your tasks from the command line.

### Add a New Task

To add a task, use the `-add` flag followed by the task description:

```bash
./todo -add "Buy groceries"
```

### List All Tasks

To display all the tasks, use the `-list` flag:

```bash
./todo -list
```

### Mark a Task as Completed

To mark a specific task as completed, use the `-complete` flag followed by the task number:

```bash
./todo -complete 1
```

### Delete a Task

To delete a task, use the `-del` flag followed by the task number:

```bash
./todo -del 2
```

### Clear All Tasks

To remove all tasks from the list, use the `-clear` flag:

```bash
./todo -clear
```

### Example Workflow

1. Add tasks:
   ```bash
   ./todo -add "Write CLI app"
   ./todo -add "Test the app"
   ```
2. List tasks:
   ```bash
   ./todo -list
   ```
3. Mark task 1 as completed:
   ```bash
   ./todo -complete 1
   ```
4. Delete task 2:
   ```bash
   ./todo -del 2
   ```
5. Clear all tasks:
   ```bash
   ./todo -clear
   ```

## Task Storage

All tasks are saved in a `.todos.json` file in the root directory of the project. You can manually edit this file if necessary.

## Contributing

If you would like to contribute to this project, feel free to submit issues and pull requests. Any contributions are welcome!

