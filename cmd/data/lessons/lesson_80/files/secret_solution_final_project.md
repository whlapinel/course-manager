# Solution for Final Project

```python

def load_tasks(file_name):
    try:
        with open(file_name, "r") as file:
            return [line.strip() for line in file.readlines()]
    except FileNotFoundError:
        return []

def save_tasks(file_name, tasks):
    with open(file_name, "w") as file:
        file.write("\n".join(tasks))

def display_tasks(tasks):
    print("\nTo-Do List:")
    for i, task in enumerate(tasks, start=1):
        print(f"{i}. {task}")
    print()

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
            display_tasks(tasks)
        elif choice == "2":
            new_task = input("Enter the new task: ")
            tasks.append(f"[ ] {new_task}")
            save_tasks(file_name, tasks)
            print("\nTask added!")
        elif choice == "3":
            display_tasks(tasks)
            try:
                task_num = int(input("Enter the number of the task to mark as complete: "))
                tasks[task_num - 1] = tasks[task_num - 1].replace("[ ]", "[x]")
                save_tasks(file_name, tasks)
                print("\nTask marked as complete!")
            except (IndexError, ValueError):
                print("\nInvalid task number.")
        elif choice == "4":
            print("Goodbye!")
            break
        else:
            print("\nInvalid choice. Please try again.")

if __name__ == "__main__":
    main()
```
