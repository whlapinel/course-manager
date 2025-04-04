---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Phones in bags, bags against the wall
- Create the following and open it in VS Code

```text
/python
  /unit_5
    /lesson_5
```
# Agenda

- Building CLI apps with OOP

# Housekeeping

- Permission forms

## 🏆 Goal:
Create a **To-Do List App** where the user can:
- Add a task  
- View all tasks  
- Mark a task as complete  
- Exit the program  

# Step 1: Create the `Task` Class
Create a `Task` class to store the task details and status:

```python
class Task:
    def __init__(self, description):
        self.description = description
        self.completed = False
    
    def mark_complete(self):
        self.completed = True
    
    def display(self):
        status = "✓" if self.completed else "✗"
        return f"[{status}] {self.description}"
```

# ✅ **Explanation:**
- `__init__` sets up the task description and status  
- `mark_complete()` sets `completed` to `True`  
- `display()` formats the task for display  

# Step 2: Create a `ToDoList` Class

Create a `ToDoList` class to manage a list of tasks:

```python
class ToDoList:
    def __init__(self):
        self.tasks = []
    
    def add_task(self, description):
        task = Task(description)
        self.tasks.append(task)
        print(f"Task '{description}' added.")
```

# Step 2: Create a `ToDoList` Class (continued)

```python
    # continued from previous, this is within the class ToDoList
    def display_tasks(self):
        if not self.tasks:
            print("\nNo tasks available.\n")
            return
        print("\nTasks:")
        for i, task in enumerate(self.tasks):
            print(f"{i + 1}. {task.display()}")
        print()
```

# Step 2: Create a `ToDoList` Class (continued)

```python
    # continued from previous, this is within the class ToDoList
    def mark_task_complete(self, index):
        if 0 <= index < len(self.tasks):
            self.tasks[index].mark_complete()
            print(f"Task '{self.tasks[index].description}' marked as complete.")
        else:
            print("Invalid task number.")
```

# ✅ **Explanation:**
- `__init__` creates an empty list of tasks  
- `add_task()` creates and adds a new task  
- `display_tasks()` shows all tasks  
- `mark_task_complete()` marks a task as complete  

# Step 3: Build the Command-Line Interface
Create a loop to allow the user to interact with the app:

```python
def main():
    todo = ToDoList()

    while True:
        print("\nOptions:")
        print("1. Add task")
        print("2. View tasks")
        print("3. Mark task complete")
        print("4. Exit")

        choice = input("Enter your choice: ")

```
# Step 3: Build the Command-Line Interface (continued)
```python
        if choice == "1":
            description = input("Enter task description: ")
            todo.add_task(description)
        elif choice == "2":
            todo.display_tasks()
        elif choice == "3":
            todo.display_tasks()
            try:
                index = int(input("Enter task number to mark as complete: ")) - 1
                todo.mark_task_complete(index)
            except ValueError:
                print("Please enter a valid number.")
        elif choice == "4":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Try again.")
```

# Step 3: Build the Command-Line Interface (continued)
```python
if __name__ == "__main__":
    main()
```

## ✅ **Explanation:**
- The loop keeps running until the user selects "Exit"  
- Input is collected and passed to the appropriate `ToDoList` method  
- `try/except` handles invalid inputs  

# ✅ Sample Run
```text
Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 1
Enter task description: Finish homework
Task 'Finish homework' added.

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 2

Tasks:
1. [✗] Finish homework

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 3
Enter task number to mark as complete: 1
Task 'Finish homework' marked as complete.

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 4
Goodbye!
```

# 🏆 Done!
✅ Fully working app using OOP  
✅ Handles edge cases like invalid inputs  
✅ Clean and readable code  
