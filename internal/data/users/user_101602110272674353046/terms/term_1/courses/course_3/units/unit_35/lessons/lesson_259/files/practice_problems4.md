# Practice Problems 9.4 (Advanced Python Features)

## **1. List Comprehensions**

**Question**: What does the following code produce?

```python
nums = [1, 2, 3, 4]
squares = [n ** 2 for n in nums]
print(squares)
```

- A) `[1, 4, 9, 16]`
- B) `[1, 2, 3, 4]`
- C) `[2, 4, 6, 8]`
- D) `None`

## **2. Lambda Functions**

**Question**: What is the output of this code?

```python
double = lambda x: x * 2
print(double(5))
```

- A) `5`
- B) `25`
- C) `10`
- D) `Error`

## **3. Closures**

**Question**: What will this code output?

```python
def outer(x):
    def inner(y):
        return x + y
    return inner

add_five = outer(5)
print(add_five(10))
```

- A) `15`
- B) `5`
- C) `10`
- D) `Error`

## **4. File Handling**

**Question**: What does the following code do?

```python
with open("data.txt", "w") as file:
    file.write("Hello, world!")
```

- A) Appends "Hello, world!" to `data.txt`
- B) Writes "Hello, world!" to `data.txt`, overwriting its contents
- C) Reads the contents of `data.txt`
- D) Creates `data.txt` without writing anything

## **5. Error Handling**

**Question**: What is the output of this code?

```python
try:
    result = 10 / 0
except ZeroDivisionError:
    print("Cannot divide by zero!")
finally:
    print("Execution complete.")
```

- A) `"Execution complete."`
- B) `"Cannot divide by zero!"`
- C) `"Cannot divide by zero!\nExecution complete."`
- D) `Error`

## **6. List Comprehensions with Conditionals**

**Question**: What does the following code produce?

```python
nums = [1, 2, 3, 4]
evens = [n for n in nums if n % 2 == 0]
print(evens)
```

- A) `[1, 3]`
- B) `[2, 4]`
- C) `[1, 2, 3, 4]`
- D) `[]`

## **7. Nested Functions**

**Question**: What will this code output?

```python
def outer():
    def inner():
        return "Hello from inner!"
    return inner()

print(outer())
```

- A) `"Hello from inner!"`
- B) `"Hello from outer!"`
- C) `None`
- D) `Error`

## **8. Reading from Files**

**Question**: What will this code output, assuming `data.txt` contains the line `"Python"`?

```python
with open("data.txt", "r") as file:
    print(file.read())
```

- A) `"Python"`
- B) `None`
- C) `Error`
- D) `""` (empty string)

## **9. Lambda and Map**

**Question**: What does the following code produce?

```python
nums = [1, 2, 3]
result = list(map(lambda x: x * 3, nums))
print(result)
```

- A) `[1, 2, 3]`
- B) `[3, 6, 9]`
- C) `[2, 4, 6]`
- D) `None`

## **10. Combining Comprehensions and Files**

**Question**: What does the following code do?

```python
lines = ["Line 1", "Line 2", "Line 3"]
with open("output.txt", "w") as file:
    file.writelines([line + "\n" for line in lines])
```

- A) Writes each line to `output.txt` with a newline
- B) Appends each line to `output.txt` without newlines
- C) Creates `output.txt` but writes nothing
- D) Raises an error
