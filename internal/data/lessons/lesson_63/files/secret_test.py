# 12.1 Tuples
# Q1: Write a pure function `tuple_to_list` that converts a tuple into a list.

# Sample Test
# assert tuple_to_list((1, 2, 3)) == [1, 2, 3]
# assert tuple_to_list(("a", "b", "c")) == ["a", "b", "c"]

# Q2: Write a pure function `multiply_tuple_elements` that takes a tuple of numbers and returns the product of all its elements.

# Sample Test
# assert multiply_tuple_elements((1, 2, 3, 4)) == 24
# assert multiply_tuple_elements((5, 6)) == 30

# 12.2 Lists
# Q3: Write a pure function `list_length` that returns the length of a list.

# Sample Test
# assert list_length([1, 2, 3]) == 3
# assert list_length([]) == 0

# Q4: Write a pure function `list_average` that calculates the average of numbers in a list.

# Sample Test
# assert list_average([1, 2, 3, 4]) == 2.5
# assert list_average([]) == 0

# 12.3 For Loops and Lists
# Q5: Write a pure function `filter_positive` that takes a list of integers and returns a new list containing only the positive numbers.

# Sample Test
# assert filter_positive([-3, -2, 0, 1, 2]) == [1, 2]
# assert filter_positive([-1, -5]) == []

# Q6: Write a pure function `find_max_index` that returns the index of the largest number in a list.

# Sample Test
# assert find_max_index([10, 20, 30]) == 2
# assert find_max_index([5, 1, 5, 7]) == 3

# 12.4 List Methods
# Q7: Write a pure function `reverse_list` that takes a list and returns a reversed version of it.

# Sample Test
# assert reverse_list([1, 2, 3]) == [3, 2, 1]
# assert reverse_list(["a", "b", "c"]) == ["c", "b", "a"]

# Q8: Write a pure function `remove_all_occurrences` that removes all occurrences of a given value from a list.

# Sample Test
# assert remove_all_occurrences([1, 2, 2, 3], 2) == [1, 3]
# assert remove_all_occurrences([4, 5, 4], 4) == [5]

# 12.5 Creating and Altering Data Structures
# Q9: Write a pure function `create_repeated_list` that creates a list containing a value repeated `n` times.

# Sample Test
# assert create_repeated_list(7, 3) == [7, 7, 7]
# assert create_repeated_list("a", 5) == ["a", "a", "a", "a", "a"]

# Q10: Write a pure function `merge_lists` that takes two lists and returns a new list with elements from both.

# Sample Test
# assert merge_lists([1, 2], [3, 4]) == [1, 2, 3, 4]
# assert merge_lists(["a"], ["b", "c"]) == ["a", "b", "c"]

# Q11: Write a pure function `double_numbers` that takes a list of numbers and returns a list with each number doubled.

# Sample Test
# assert double_numbers([1, 2, 3]) == [2, 4, 6]
# assert double_numbers([-1, -2]) == [-2, -4]

# 13.1 2D Lists
# Q12: Write a pure function `column_sum` that takes a 2D list and a column index and returns the sum of the numbers in that column.

# Sample Test
# assert column_sum([[1, 2], [3, 4], [5, 6]], 1) == 12
# assert column_sum([[7, 8], [9, 10]], 0) == 16

# 13.2 List Comprehensions
# Q13: Write a pure function `flatten_2d_list` that takes a 2D list and flattens it into a 1D list.

# Sample Test
# assert flatten_2d_list([[1, 2], [3, 4], [5, 6]]) == [1, 2, 3, 4, 5, 6]
# assert flatten_2d_list([[10, 20], [30]]) == [10, 20, 30]

# 13.3 Packing and Unpacking
# Q14: Write a pure function `pack_list` that packs multiple arguments into a list.

# Sample Test
# assert pack_list(1, 2, 3) == [1, 2, 3]
# assert pack_list("a", "b") == ["a", "b"]

# 13.4 Dictionaries
# Q15: Write a pure function `add_to_dict` that takes a dictionary, a key, and a value, and returns a new dictionary with the key-value pair added.

# Sample Test
# assert add_to_dict({"a": 1}, "b", 2) == {"a": 1, "b": 2}
# assert add_to_dict({}, "x", 42) == {"x": 42}
