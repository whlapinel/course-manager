# Final Project Python I Programming Honors

## **Project: Simple To-Do List Manager**

### **Project Overview**
Students will build a basic **to-do list manager** that can:
1. Add tasks to a to-do list.
2. View the current list of tasks.
3. Mark tasks as complete.
4. Save the to-do list to a file (`todo.txt`) and load it back.

### **Specifications**
1. **Input**: Students will interact with the application via a menu-driven interface.
2. **Output**: Tasks are displayed in a formatted list, and the file (`todo.txt`) is updated to reflect changes.
3. **File Handling**: The program will read from and write to a file (`todo.txt`) to persist the to-do list between runs.

### **Example Output**
```plaintext
Welcome to the To-Do List Manager!

Menu:
1. View tasks
2. Add a task
3. Mark a task as complete
4. Quit

Choose an option: 1

To-Do List:
1. [ ] Buy groceries
2. [ ] Finish homework

Choose an option: 2
Enter the new task: Call Mom

Task added!

Choose an option: 1

To-Do List:
1. [ ] Buy groceries
2. [ ] Finish homework
3. [ ] Call Mom

Choose an option: 3
Enter the number of the task to mark as complete: 2

Task "Finish homework" marked as complete!

Choose an option: 1

To-Do List:
1. [ ] Buy groceries
2. [x] Finish homework
3. [ ] Call Mom
```

## Starter Code

```python
def load_tasks(file_name):
    # TODO: Read tasks from the file and return them as a list
    pass

def save_tasks(file_name, tasks):
    # TODO: Write tasks to the file
    pass

def display_tasks(tasks):
    # TODO: Print tasks in a numbered list
    pass

def main():
    file_name = "todo.txt"
    tasks = load_tasks(file_name)

    while True:
        print("Menu:")
        print("1. View tasks")
        print("2. Add a task")
        print("3. Mark a task as complete")
        print("4. Quit")
        choice = input("Choose an option: ")

        if choice == "1":
            # TODO: Display tasks
            pass
        elif choice == "2":
            # TODO: Add a new task
            pass
        elif choice == "3":
            # TODO: Mark a task as complete
            pass
        elif choice == "4":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()
```
