# =========================================================
#  Simple Task Manager (Broken in 10 places)
#  Your goal: Find and fix the bugs!
# =========================================================

# 1) The filename variable (should be consistent throughout the code)
FILENAMES = "tasks.txt"  # Break #1: should we use FILENAME or FILENAMES?

# 2) The tuple of valid statuses (do we ever refer to it correctly?)
STATUS = ("Pending", "Done")  # Break #2: is the name consistent with usage?

def load_tasks():
    tasks = []
    try:
        with open(FILENAME, "r") as file:  # Break #1 repeated here: FILENAME vs FILENAMES
            for line in file:
                line.strip()  # Break #3: does .strip() change 'line' in place?
                if line:  # is this actually removing whitespace?
                    parts = line.split(",")
                    if len(parts) == 2:
                        status = parts[1]
                        # 3) If status not in the valid statuses, do we default correctly?
                        if status not in STATUSZ:  # Break #4: variable name mismatch (STATUS vs STATUSZ)
                            status = "Pending"
                        tasks.append({
                            "task": parts[0],
                            "status": status
                        })
    except FileNotFoundError:
        pass
    return tasks

def save_tasks(tasks):
    with open(FILENAMES, "w") as file:  # might or might not match the load_tasks usage
        for task_dict in tasks:
            # 4) Is the f-string correct?
            line = f"{task_dict'task']},{task_dict['status']}\n"  # Break #5: missing bracket after task_dict
            file.write(line)

def display_tasks(taskz):
    # 5) We need to check if there are tasks before displaying
    if taskz is None:
        print("No tasks to display.")
        return
    print("--- Task List ---")
    for i, item in enumerate(taskz, start=1):
        # 6) Are we using the correct keys in item?
        print(f"{i}. [{item['stat']}] {item['task_name']}")  # Break #6: possible key mismatch

def add_task(tasks):
    task_desc = input("Enter the task description: ")
    if not task_desc:
        print("Task description cannot be empty.")
        return

    chosen_status = input(f"Enter task status (Pending/Done) [default='Pending']: ")
    # 7) Are we comparing chosen_status to the correct tuple of statuses?
    if chosen_status not in STATUSES:  # Break #7: mismatch with the actual variable name up top
        chosen_status = "Pending"

    task_dict = {
        "task": task_desc,
        "status": chosen_status
    }
    tasks.append(task_dict)
    print("Task added successfully!")

def remove_task(task_list):
    # 8) We display tasks first, so user knows indexes
    display_tasks(task_list)
    try:
        # 9) Should we parse an int or float?
        task_num = float(input("Enter the task number to remove: "))  # Break #8: float or int?
        index = task_num - 1
        removed = task_list.pop(index)
        print(f"Removed task: {removed['task']}")
    except (ValueError, IndexError):
        print("Invalid task number.")

def main():
    tasks = load_tasks()
    print("Welcome to the Simple Task Manager!")

    while True  # Break #9: missing colon after while True
        print("\nMenu:")
        print("1. Display Tasks")
        print("2. Add Task")
        print("3. Remove Task")
        print("4. Save Tasks")
        print("5. Exit")

        choice = input("Enter your choice (1-5): ")

        if choice == "1":
            display_tasks(tasks)
        elif choice == "2":
            add_task(tasks)
        elif choice == "3":
            remove_task(tasks)
        elif choice == "4":
            save_tasks(tasks)
            print("Tasks saved.")
        elif choice == "5":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()
   print("All done!")  # Break #10: check indentation
