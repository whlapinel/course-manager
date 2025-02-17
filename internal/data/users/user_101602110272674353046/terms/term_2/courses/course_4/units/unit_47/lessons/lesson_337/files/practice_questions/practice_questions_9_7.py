# Practice Questions 9.7: Working with Nested Structures
# For each problem, write a function that meets the requirements specified. Use the assert statements to test your code by uncommenting them.

# 1. 'sum_nested_list': Return the sum of all numbers in a list of lists
# Accepts a list of lists of integers
# Returns an integer

# assert sum_nested_list([[1, 2], [3, 4]]) == 10
# assert sum_nested_list([]) == 0

# 2. 'flatten_dict': Flatten a dictionary with lists as values into a single list
# Accepts a dictionary with lists as values
# Returns a list

# assert flatten_dict({"a": [1, 2], "b": [3, 4]}) == [1, 2, 3, 4]
# assert flatten_dict({}) == []

# 3. 'odd_value_keys': Return a list of keys whose values are odd numbers
# Accepts a dictionary with integer values
# Returns a list of keys

# assert odd_value_keys({"a": 1, "b": 2, "c": 3}) == ["a", "c"]
# assert odd_value_keys({}) == []

# 4. 'count_chars_nested': Count the total number of characters in all strings in a nested list
# Accepts a list of lists of strings
# Returns an integer

# assert count_chars_nested([["hello", "world"], ["python"]]) == 15
# assert count_chars_nested([]) == 0

# 5. 'transpose': Transpose a 2D list (list of lists)
# Accepts a list of lists
# Returns a transposed list of lists

# assert transpose([[1, 2], [3, 4]]) == [[1, 3], [2, 4]]
# assert transpose([]) == []
