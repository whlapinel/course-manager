# Assignment 5.4: Closures

## Problems

### 1. **Create an Adder**
Write a function `make_adder` that takes a number `x` and returns a function that takes another number `y` and returns the sum of `x` and `y`.

**Example:**
```python
add_five = make_adder(5)
print(add_five(10))  # Output: 15
```

---

### 2. **Create a Multiplier**
Write a function `make_multiplier` that takes a number `n` and returns a function that multiplies its input by `n`.

**Example:**
```python
double = make_multiplier(2)
print(double(4))  # Output: 8
```

---

### 3. **Counter**
Write a function `make_counter` that returns a function. Each time the inner function is called, it should increase a count by 1 and return the new value.

**Example:**
```python
counter = make_counter()
print(counter())  # Output: 1
print(counter())  # Output: 2
```

---

### 4. **Power Function**
Write a function `make_power` that takes an exponent `n` and returns a function that raises its input to the power of `n`.

**Example:**
```python
square = make_power(2)
print(square(3))  # Output: 9
```

---

### 5. **String Formatter**
Write a function `make_greeter` that takes a greeting string (`"Hello"`, `"Hi"`, etc.) and returns a function that takes a name and returns the greeting followed by the name.

**Example:**
```python
greet = make_greeter("Hello")
print(greet("Alice"))  # Output: Hello Alice
```

---

### 6. **Discount Calculator**
Write a function `make_discount` that takes a discount percentage and returns a function that applies the discount to a price.

**Example:**
```python
ten_percent_off = make_discount(10)
print(ten_percent_off(100))  # Output: 90.0
```

---

### 7. **Repeater**
Write a function `make_repeater` that takes an integer `n` and returns a function that repeats a string `n` times.

**Example:**
```python
repeat_three_times = make_repeater(3)
print(repeat_three_times("Hi"))  # Output: HiHiHi
```
