# This will be updated on quiz day


# 12.1 Tuples
# Q1: Write a function `sum_tuple` that takes a tuple of numbers and returns the sum of all its elements.
def sum_tuple(numbers):
    return sum(numbers)


# Sample Test
print(sum_tuple((1, 2, 3)))  # Expected: 6
print(sum_tuple((5, 10, 15)))  # Expected: 30


# 12.2 Lists
# Q2: Write a function `max_list` that takes a list of numbers and returns the largest number.
def max_list(numbers):
    return max(numbers)


# Sample Test
print(max_list([1, 2, 3, 4, 5]))  # Expected: 5
print(max_list([-10, -5, 0, 5]))  # Expected: 5


# 12.3 For Loops and Lists
# Q3: Write a function `count_evens` that takes a list of integers and returns the count of even numbers.
def count_evens(numbers):
    return sum(1 for num in numbers if num % 2 == 0)


# Sample Test
print(count_evens([1, 2, 3, 4, 5, 6]))  # Expected: 3
print(count_evens([1, 3, 5, 7]))  # Expected: 0


# 12.4 List Methods
# Q4: Write a function `remove_first_occurrence` that removes the first occurrence of a given value from a list.
def remove_first_occurrence(lst, value):
    lst_copy = lst.copy()
    if value in lst_copy:
        lst_copy.remove(value)
    return lst_copy


# Sample Test
print(remove_first_occurrence([1, 2, 3, 2, 4], 2))  # Expected: [1, 3, 2, 4]
print(remove_first_occurrence([5, 6, 7], 8))  # Expected: [5, 6, 7]


# 12.5 Creating and Altering Data Structures
# Q5: Write a function `create_number_list` that takes two integers, `start` and `end`, and returns a list of numbers between them (inclusive).
def create_number_list(start, end):
    return list(range(start, end + 1))


# Sample Test
print(create_number_list(1, 5))  # Expected: [1, 2, 3, 4, 5]
print(create_number_list(3, 3))  # Expected: [3]


# 13.1 2D Lists
# Q6: Write a function `row_sum` that takes a 2D list (list of lists) and an integer `row` and returns the sum of the elements in that row.
def row_sum(matrix, row):
    return sum(matrix[row])


# Sample Test
print(row_sum([[1, 2, 3], [4, 5, 6], [7, 8, 9]], 1))  # Expected: 15
print(row_sum([[10, 20], [30, 40], [50, 60]], 2))  # Expected: 110


# 13.2 List Comprehensions
# Q7: Write a function `square_odds` that takes a list of integers and returns a new list with the squares of all the odd numbers.
def square_odds(numbers):
    return [num**2 for num in numbers if num % 2 != 0]


# Sample Test
print(square_odds([1, 2, 3, 4, 5]))  # Expected: [1, 9, 25]
print(square_odds([2, 4, 6]))  # Expected: []


# 13.3 Packing and Unpacking
# Q8: Write a function `swap_first_last` that takes a list and returns a new list with the first and last elements swapped.
def swap_first_last(lst):
    if len(lst) < 2:
        return lst
    return [lst[-1], *lst[1:-1], lst[0]]


# Sample Test
print(swap_first_last([1, 2, 3, 4]))  # Expected: [4, 2, 3, 1]
print(swap_first_last([7]))  # Expected: [7]


# 13.4 Dictionaries
# Q9: Write a pure function `invert_dict` that takes a dictionary and returns a new dictionary where keys and values are swapped.
def invert_dict(d):
    return {v: k for k, v in d.items()}


# Sample Test
print(invert_dict({"a": 1, "b": 2, "c": 3}))  # Expected: {1: "a", 2: "b", 3: "c"}
print(invert_dict({}))  # Expected: {}


# 13.5 Extending Data Structures
# Q10: Write a pure function `merge_dicts` that takes two dictionaries and merges them, with the second dictionary overwriting values of the first for matching keys.
def merge_dicts(d1, d2):
    merged = d1.copy()
    merged.update(d2)
    return merged


# Sample Test
print(
    merge_dicts({"a": 1, "b": 2}, {"b": 3, "c": 4})
)  # Expected: {"a": 1, "b": 3, "c": 4}
print(merge_dicts({}, {"x": 42}))  # Expected: {"x": 42}
