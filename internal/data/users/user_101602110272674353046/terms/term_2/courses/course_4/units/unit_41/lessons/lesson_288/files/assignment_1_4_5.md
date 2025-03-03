# Assignment 4.5

## Instructions
- Complete CodeHS Assignments 7.1 and 7.2 (loops)
- Submit all the below python files as a single .zip file
  - (select all python files, right-click and click send to compressed)

### **📝 Module Practice Problems**

#### **1️⃣ Creating & Importing a Simple Module**
📌 **Task:**  
- Create a module called `math_helpers.py`.  
- Inside the module, define two functions:  
  ```python
  def square(n):  # Returns the square of n
      return n * n

  def cube(n):  # Returns the cube of n
      return n * n * n
  ```
- Create another file `main.py` and:
  - Import `math_helpers`
  - Call `square(4)` and `cube(3)`
  - Print the results

📌 **Submission:** Submit both `math_helpers.py` and `main.py`.

---

#### **2️⃣ Using `from module import function`**
📌 **Task:**  
- Create a module `greetings.py` with:
  ```python
  def hello(name):
      return f"Hello, {name}!"

  def goodbye(name):
      return f"Goodbye, {name}!"
  ```
- In a separate file `test_greetings.py`:
  - Import `hello` **directly** from `greetings`
  - Call `hello("Alice")` and print the result
  - Try calling `goodbye("Bob")` (what happens?)

📌 **Hint:** Use `from greetings import hello`.

📌 **Submission:** Submit `greetings.py` and `test_greetings.py`.

---

#### **3️⃣ Importing with an Alias**
📌 **Task:**  
- Create a module `stats.py` with:
  ```python
  def average(numbers):
      return sum(numbers) / len(numbers)

  def max_number(numbers):
      return max(numbers)
  ```
- In another file `analyze.py`:
  - Import `stats` as `st`
  - Call `st.average([10, 20, 30])`
  - Call `st.max_number([4, 7, 2, 9])`
  - Print the results

📌 **Submission:** Submit `stats.py` and `analyze.py`.

---

#### **4️⃣ `__name__ == "__main__"`**
📌 **Task:**  
- Create a module `converter.py` with:
  ```python
  def inches_to_cm(inches):
      return inches * 2.54

  if __name__ == "__main__":
      print("This file is being run directly.")
  ```
- In another file `test_converter.py`:
  - Import `inches_to_cm`
  - Call `inches_to_cm(10)` and print the result
  - Run **both** files separately and observe output

📌 **Hint:** Run:
```sh
python converter.py
python test_converter.py
```

📌 **Submission:** Submit `converter.py` and `test_converter.py`.

---

#### **5️⃣ Exploring Built-in Modules**
📌 **Task:**  
- Create a module `random_utils.py` that:
  - Uses `random.randint()` to generate a number between 1-100
  - Uses `time.sleep(2)` to delay before printing it  
- In `use_random.py`:
  - Import `random_utils`
  - Call the function from `random_utils`

📌 **Submission:** Submit `random_utils.py` and `use_random.py`.
