
# 🔄 While Loop Practice - Hints

# 1️⃣ Print Odd Numbers
💡 **Hint:** Use the modulo operator `%` to check if a number is odd.

```python
if num % 2 != 0:
    # It's odd
```

# 2️⃣ Repeat Until Uppercase
💡 **Hint:** The `.isupper()` string method checks if a string is uppercase.

```python
word = input("Enter a word: ")
while not word.isupper():
    word = input("Try again: ")
```

# 3️⃣ Multiplication Table
💡 **Hint:** Use a counter variable and update it inside the loop.

```python
num = int(input("Enter a number: "))
i = 1
while i <= 10:
    print(num * i)
    i += 1
```

# 4️⃣ Factorial Calculation
💡 **Hint:** Use a loop with `*=`, starting from `1`.

```python
factorial = 1
while num > 1:
    factorial *= num
    num -= 1
```

# 5️⃣ Sum of Digits
💡 **Hint:** Use `num % 10` to get the last digit and `num // 10` to remove it.

```python
sum_digits = 0
while num > 0:
    sum_digits += num % 10
    num //= 10
```

# 6️⃣ Reverse a Number
💡 **Hint:** Multiply the reversed number by 10 and add the last digit.

```python
reversed_num = 0
while num > 0:
    reversed_num = reversed_num * 10 + num % 10
    num //= 10
```

# 7️⃣ Repeat Until Palindrome
💡 **Hint:** Use slicing `[::-1]` to check if a word is a palindrome.

```python
if word == word[::-1]:
    print("It's a palindrome!")
```

# 8️⃣ Divide Until 1
💡 **Hint:** Use integer division `//` to repeatedly divide by 2.

```python
while num > 1:
    num //= 2
```

# 9️⃣ Count Letter Frequency
💡 **Hint:** The `.count()` method can count occurrences of a letter.

```python
count = text.count('e')
```

# 🔟 Number Guesser with Hints
💡 **Hint:** Use conditionals to check if the guess is too high or too low.

```python
if guess > target:
    print("Too high!")
elif guess < target:
    print("Too low!")
```
