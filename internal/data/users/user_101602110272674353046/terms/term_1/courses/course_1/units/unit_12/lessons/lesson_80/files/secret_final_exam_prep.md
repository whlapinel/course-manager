# Final Project Prep Activities

## **Scaffolding Activities**

### **Activity 1: File Reading and Writing**

**Goal**: Practice reading from and writing to files.

**Task**:
1. Create a file `sample.txt` with the following content:
   ```text
   Task 1: Buy groceries
   Task 2: Finish homework
   Task 3: Call Mom
   ```
2. Write a program to:
   - Read the file and print its contents line by line.
   - Append a new task to the file.

**Code Template**:
```python
def read_file(file_name):
    with open(file_name, "r") as file:
        for line in file:
            print(line.strip())

def append_to_file(file_name, task):
    with open(file_name, "a") as file:
        file.write(f"{task}\n")

read_file("sample.txt")
append_to_file("sample.txt", "Task 4: Walk the dog")
```

### **Activity 2: Displaying Tasks with Formatting**

**Goal**: Learn to format and display a list of tasks.

**Task**:
1. Create a list of tasks:
   ```python
   tasks = ["[ ] Buy groceries", "[ ] Finish homework", "[x] Call Mom"]
   ```
2. Write a program to display the tasks in a numbered list.

**Code Template**:
```python
def display_tasks(tasks):
    print("To-Do List:")
    for i, task in enumerate(tasks, start=1):
        print(f"{i}. {task}")

tasks = ["[ ] Buy groceries", "[ ] Finish homework", "[x] Call Mom"]
display_tasks(tasks)
```

---

### **Activity 3: User Input for Adding Tasks**

**Goal**: Teach students to add items to a list based on user input.

**Task**:
1. Start with an empty list `tasks = []`.
2. Write a program that asks the user to enter new tasks.
3. Add tasks to the list and print the updated list after each addition.

**Code Template**:
```python
def add_task(tasks):
    new_task = input("Enter a new task: ")
    tasks.append(f"[ ] {new_task}")
    print("Task added!")

tasks = []
add_task(tasks)
print(tasks)
```

---

### **Activity 4: Marking Tasks as Complete**

**Goal**: Teach students to update list items based on user input.

**Task**:
1. Use the list of tasks:
   ```python
   tasks = ["[ ] Buy groceries", "[ ] Finish homework", "[x] Call Mom"]
   ```
2. Ask the user for the task number to mark as complete.
3. Update the corresponding task with `[x]`.

**Code Template**:
```python
def mark_task_complete(tasks):
    task_num = int(input("Enter the task number to mark as complete: "))
    tasks[task_num - 1] = tasks[task_num - 1].replace("[ ]", "[x]")
    print("Task marked as complete!")

tasks = ["[ ] Buy groceries", "[ ] Finish homework", "[x] Call Mom"]
mark_task_complete(tasks)
print(tasks)
```

---

### **Activity 5: Combining File and Task Management**

**Goal**: Practice saving tasks to a file and loading them back.

**Task**:
1. Use the list of tasks from `Activity 4`.
2. Write tasks to a file (`todo.txt`).
3. Reload tasks from the file and display them.

**Code Template**:
```python
def save_tasks(file_name, tasks):
    with open(file_name, "w") as file:
        file.write("\n".join(tasks))

def load_tasks(file_name):
    with open(file_name, "r") as file:
        return [line.strip() for line in file.readlines()]

tasks = ["[ ] Buy groceries", "[ ] Finish homework", "[x] Call Mom"]
save_tasks("todo.txt", tasks)
loaded_tasks = load_tasks("todo.txt")
print("Loaded tasks:", loaded_tasks)
```
