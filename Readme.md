Here's a concise README for your Go-based Todo CLI app:

---

# Todo CLI App

A simple command-line tool written in Go to manage your to-do tasks efficiently.

## Features

- Add new tasks
- Mark tasks as complete
- List all tasks
- Edit tasks
- Delete tasks

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory:

   ```bash
   cd todo-cli-app
   ```

3. Build the app:

   ```bash
   go build -o Todo.exe TodoCli
   ```

4. Run the app:

   ```bash
   ./Todo.exe --list
   ```

## Usage

- **Add a task:**

   ```bash
   ./Todo.exe --add "Task Title"
   ```

- **List all tasks:**

   ```bash
   ./Todo.exe --list
   ```

- **Toggle a task:**

   ```bash
   ./Todo.exe --toggle <task-id>
   ```

- **Edit a task:**

   ```bash
   ./Todo.exe --edit <task-id>:<title>
   ```

## License

This project is licensed under the MIT License.

---

Feel free to adjust the content or details as needed!