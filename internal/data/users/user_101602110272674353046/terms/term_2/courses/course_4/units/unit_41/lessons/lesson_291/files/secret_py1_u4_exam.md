# Python 1 Unit 4 Exam
Functions and modules

Choose any 5 problems

1. **Square a Number**
   Write a function `square(n: int) -> int` that returns the square of `n`.

   ```python
   assert square(3) == 9
   assert square(-4) == 16
   assert square(0) == 0
   ```

2. **Get Absolute Value**  
   Write a function `absolute(n: float) -> float` that returns the absolute value of `n`. Use the `abs()` function.  

   ```python
   assert absolute(-5) == 5
   assert absolute(3.2) == 3.2
   assert absolute(0) == 0
   ```

3. **Find Maximum of Two**  
   Write a function `max_of_two(a: int, b: int) -> int` that returns the larger of the two numbers. Use the `max()` function.  

   ```python
   assert max_of_two(3, 7) == 7
   assert max_of_two(-2, -5) == -2
   assert max_of_two(4, 4) == 4
   ```

4. **Find Minimum of Two**  
   Write a function `min_of_two(a: int, b: int) -> int` that returns the smaller of the two numbers. Use the `min()` function.  

   ```python
   assert min_of_two(3, 7) == 3
   assert min_of_two(-2, -5) == -5
   assert min_of_two(4, 4) == 4
   ```

5. **Round a Number**  
   Write a function `round_number(n: float) -> int` that rounds `n` to the nearest whole number. Use the `round()` function.  

   ```python
   assert round_number(3.7) == 4
   assert round_number(2.3) == 2
   assert round_number(-1.5) == -2
   ```

6. **Sum a List**  
   Write a function `sum_list(numbers: list) -> int` that returns the sum of all numbers in the list. Use the `sum()` function.  

   ```python
   assert sum_list([1, 2, 3]) == 6
   assert sum_list([-1, -2, -3]) == -6
   assert sum_list([]) == 0
   ```

7. **Get Length of a String**  
   Write a function `string_length(s: str) -> int` that returns the length of the given string. Use the `len()` function.  

   ```python
   assert string_length("hello") == 5
   assert string_length("") == 0
   assert string_length("Python") == 6
   ```

8. **Check if All Are True**  
   Write a function `all_true(values: list) -> bool` that returns `True` if all elements in `values` are truthy, otherwise returns `False`. Use the `all()` function.  

   ```python
   assert all_true([True, 1, "non-empty"]) == True
   assert all_true([True, 0, "text"]) == False
   assert all_true([]) == True
   ```

9. **Check if Any Are True**  
   Write a function `any_true(values: list) -> bool` that returns `True` if any element in `values` is truthy, otherwise returns `False`. Use the `any()` function.  

   ```python
   assert any_true([False, 0, ""]) == False
   assert any_true([False, 1, ""]) == True
   assert any_true([]) == False
   ```

10. **Convert to Integer**  
    Write a function `to_int(n: float) -> int` that converts `n` to an integer. Use the `int()` function.  

    ```python
    assert to_int(3.9) == 3
    assert to_int(-2.8) == -2
    assert to_int(0.0) == 0
    ```

11. **Convert to Float**  
    Write a function `to_float(n: int) -> float` that converts `n` to a float. Use the `float()` function.  

    ```python
    assert to_float(3) == 3.0
    assert to_float(-2) == -2.0
    assert to_float(0) == 0.0
    ```

12. **Get the Largest Number in a List**  
    Write a function `max_in_list(numbers: list) -> int` that returns the largest number in the list. Use the `max()` function.

     ```python
     assert max_in_list([1, 2, 3]) == 3
     assert max_in_list([-5, -2, -10]) == -2
     assert max_in_list([7]) == 7
     ```

13. **Get the Smallest Number in a List**  
    Write a function `min_in_list(numbers: list) -> int` that returns the smallest number in the list. Use the `min()` function.

     ```python
     assert min_in_list([1, 2, 3]) == 1
     assert min_in_list([-5, -2, -10]) == -10
     assert min_in_list([7]) == 7
     ```

14. **Count Items in a List**  
    Write a function `count_items(lst: list) -> int` that returns the number of items in `lst`. Use the `len()` function.  

    ```python
    assert count_items([1, 2, 3]) == 3
    assert count_items([]) == 0
    assert count_items(["a", "b", "c", "d"]) == 4
    ```

15. **Get the First Letter of a Word**  
    Write a function `first_letter(word: str) -> str` that returns the first letter of the given word.  

    ```python
    assert first_letter("hello") == "h"
    assert first_letter("Python") == "P"
    assert first_letter("world") == "w"
    ```
