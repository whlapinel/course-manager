# 12.1 Tuples
# Q1: Write a pure function `tuple_to_list` that converts a tuple into a list.
def tuple_to_list(tup):
    return list(tup)


# Sample Test
print(tuple_to_list((1, 2, 3)))  # Expected: [1, 2, 3]
print(tuple_to_list(("a", "b", "c")))  # Expected: ["a", "b", "c"]


# Q2: Write a pure function `multiply_tuple_elements` that takes a tuple of numbers and returns the product of all its elements.
def multiply_tuple_elements(tup):
    result = 1
    for num in tup:
        result *= num
    return result


# Sample Test
print(multiply_tuple_elements((1, 2, 3, 4)))  # Expected: 24
print(multiply_tuple_elements((5, 6)))  # Expected: 30


# 12.2 Lists
# Q3: Write a pure function `list_length` that returns the length of a list.
def list_length(lst):
    return len(lst)


# Sample Test
print(list_length([1, 2, 3]))  # Expected: 3
print(list_length([]))  # Expected: 0


# Q4: Write a pure function `list_average` that calculates the average of numbers in a list.
def list_average(lst):
    return sum(lst) / len(lst) if lst else 0


# Sample Test
print(list_average([1, 2, 3, 4]))  # Expected: 2.5
print(list_average([]))  # Expected: 0


# 12.3 For Loops and Lists
# Q5: Write a pure function `filter_positive` that takes a list of integers and returns a new list containing only the positive numbers.
def filter_positive(lst):
    return [num for num in lst if num > 0]


# Sample Test
print(filter_positive([-3, -2, 0, 1, 2]))  # Expected: [1, 2]
print(filter_positive([-1, -5]))  # Expected: []


# Q6: Write a pure function `find_max_index` that returns the index of the largest number in a list.
def find_max_index(lst):
    return lst.index(max(lst))


# Sample Test
print(find_max_index([10, 20, 30]))  # Expected: 2
print(find_max_index([5, 1, 5, 7]))  # Expected: 3


# 12.4 List Methods
# Q7: Write a pure function `reverse_list` that takes a list and returns a reversed version of it.
def reverse_list(lst):
    return lst[::-1]


# Sample Test
print(reverse_list([1, 2, 3]))  # Expected: [3, 2, 1]
print(reverse_list(["a", "b", "c"]))  # Expected: ["c", "b", "a"]


# Q8: Write a pure function `remove_all_occurrences` that removes all occurrences of a given value from a list.
def remove_all_occurrences(lst, value):
    return [x for x in lst if x != value]


# Sample Test
print(remove_all_occurrences([1, 2, 2, 3], 2))  # Expected: [1, 3]
print(remove_all_occurrences([4, 5, 4], 4))  # Expected: [5]


# 12.5 Creating and Altering Data Structures
# Q9: Write a pure function `create_repeated_list` that creates a list containing a value repeated `n` times.
def create_repeated_list(value, n):
    return [value] * n


# Sample Test
print(create_repeated_list(7, 3))  # Expected: [7, 7, 7]
print(create_repeated_list("a", 5))  # Expected: ["a", "a", "a", "a", "a"]


# Q10: Write a pure function `merge_lists` that takes two lists and returns a new list with elements from both.
def merge_lists(lst1, lst2):
    return lst1 + lst2


# Sample Test
print(merge_lists([1, 2], [3, 4]))  # Expected: [1, 2, 3, 4]
print(merge_lists(["a"], ["b", "c"]))  # Expected: ["a", "b", "c"]


# Q11: Write a pure function `double_numbers` that takes a list of numbers and returns a list with each number doubled.
def double_numbers(lst):
    return [x * 2 for x in lst]


# Sample Test
print(double_numbers([1, 2, 3]))  # Expected: [2, 4, 6]
print(double_numbers([-1, -2]))  # Expected: [-2, -4]


# 13.1 2D Lists
# Q12: Write a pure function `column_sum` that takes a 2D list and a column index and returns the sum of the numbers in that column.
def column_sum(matrix, col):
    return sum(row[col] for row in matrix)


# Sample Test
print(column_sum([[1, 2], [3, 4], [5, 6]], 1))  # Expected: 12
print(column_sum([[7, 8], [9, 10]], 0))  # Expected: 16


# 13.2 List Comprehensions
# Q13: Write a pure function `flatten_2d_list` that takes a 2D list and flattens it into a 1D list.
def flatten_2d_list(matrix):
    return [item for row in matrix for item in row]


# Sample Test
print(flatten_2d_list([[1, 2], [3, 4], [5, 6]]))  # Expected: [1, 2, 3, 4, 5, 6]
print(flatten_2d_list([[10, 20], [30]]))  # Expected: [10, 20, 30]


# 13.3 Packing and Unpacking
# Q14: Write a pure function `pack_list` that packs multiple arguments into a list.
def pack_list(*args):
    return list(args)


# Sample Test
print(pack_list(1, 2, 3))  # Expected: [1, 2, 3]
print(pack_list("a", "b"))  # Expected: ["a", "b"]


# 13.4 Dictionaries
# Q15: Write a pure function `add_to_dict` that takes a dictionary, a key, and a value, and returns a new dictionary with the key-value pair added.
def add_to_dict(d, key, value):
    new_dict = d.copy()
    new_dict[key] = value
    return new_dict


# Sample Test
print(add_to_dict({"a": 1}, "b", 2))  # Expected: {"a": 1, "b": 2}
print(add_to_dict({}, "x", 42))  # Expected: {"x": 42}
