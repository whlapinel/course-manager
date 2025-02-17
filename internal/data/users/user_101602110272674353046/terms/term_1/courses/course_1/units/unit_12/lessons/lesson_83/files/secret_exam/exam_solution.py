# A Simple Task Management Application

# Define the file where tasks will be stored
FILENAME = "tasks.txt"

# Define a tuple for valid statuses
STATUSES = ("Pending", "Done")

def load_tasks():
    """
    Reads the tasks from the file FILENAME.
    Returns a list of dictionaries, each dictionary representing a task.
    """
    tasks = []
    try:
        with open(FILENAME, "r") as file:
            for line in file:
                line = line.strip()
                if line:  # make sure the line is not empty
                    # Expected format: task_description,status
                    parts = line.split(",")
                    if len(parts) == 2:
                        task_desc = parts[0]
                        status = parts[1]
                        # Validate status, default to "Pending" if invalid
                        if status not in STATUSES:
                            status = "Pending"
                        tasks.append({"task": task_desc, "status": status})
    except FileNotFoundError:
        # If the file doesn't exist, just return an empty list
        pass
    return tasks


def save_tasks(tasks):
    """
    Saves the list of tasks (list of dictionaries) to the file FILENAME.
    Each line in the file will be task_description,status
    """
    with open(FILENAME, "w") as file:
        for task_dict in tasks:
            line = f"{task_dict['task']},{task_dict['status']}\n"
            file.write(line)


def display_tasks(tasks):
    """
    Prints all tasks with an index, so the user can see them on-screen.
    """
    if not tasks:
        print("No tasks to display.")
        return
    print("\n--- Task List ---")
    for index, task_dict in enumerate(tasks, start=1):
        print(f"{index}. [{task_dict['status']}] {task_dict['task']}")
    print("----------------\n")


def add_task(tasks):
    """
    Prompts the user for a new task description and optionally its status.
    Adds the task as a dictionary to the tasks list.
    """
    task_desc = input("Enter the task description: ").strip()
    if not task_desc:
        print("Task description cannot be empty.\n")
        return

    # Optionally ask for a status (Pending/Done). If invalid, default to Pending.
    chosen_status = input(f"Enter task status ({'/'.join(STATUSES)}) [default='Pending']: ").strip()
    if chosen_status not in STATUSES:
        chosen_status = "Pending"

    task_dict = {"task": task_desc, "status": chosen_status}
    tasks.append(task_dict)
    print(f"Task '{task_desc}' added with status '{chosen_status}'.\n")


def remove_task(tasks):
    """
    Prompts the user for a task number to remove.
    Removes the task from the tasks list if valid.
    """
    if not tasks:
        print("No tasks to remove.\n")
        return

    display_tasks(tasks)
    try:
        task_num = int(input("Enter the task number to remove: "))
        # Convert to zero-based index
        index = task_num - 1
        if index < 0 or index >= len(tasks):
            print("Invalid task number.\n")
            return
        removed_task = tasks.pop(index)
        print(f"Removed task: '{removed_task['task']}'.\n")
    except ValueError:
        print("Please enter a valid number.\n")


def main():
    """
    The main function that runs the whole application in a loop,
    showing a menu and responding to user input.
    """
    # Load tasks from file at the start
    tasks = load_tasks()
    print("Welcome to the Simple Task Manager!\n")

    while True:
        print("Menu:")
        print("1. Display Tasks")
        print("2. Add Task")
        print("3. Remove Task")
        print("4. Save Tasks")
        print("5. Exit\n")

        choice = input("Enter your choice (1-5): ").strip()

        if choice == "1":
            display_tasks(tasks)
        elif choice == "2":
            add_task(tasks)
        elif choice == "3":
            remove_task(tasks)
        elif choice == "4":
            save_tasks(tasks)
            print("Tasks saved successfully.\n")
        elif choice == "5":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please enter a number between 1 and 5.\n")


# Run the main function if this file is executed directly
if __name__ == "__main__":
    main()
