# Python Exam: Basic Data Structures
# For each problem, write a function that meets the requirements specified. Use the assert statements to test your code by uncommenting them.

# 1. 'sum_of_list': Return the sum of all numbers in a list
# Accepts a list of integers
# Returns an integer

# assert sum_of_list([1, 2, 3, 4, 5]) == 15
# assert sum_of_list([]) == 0

# 2. 'min_of_list': Return the smallest number in a list
# Accepts a list of integers
# Returns an integer

# assert min_of_list([3, 1, 4, 1, 5]) == 1
# assert min_of_list([10, 20, 30]) == 10

# 3. 'reverse_string': Return a reversed string
# Accepts a string
# Returns a string

# assert reverse_string("hello") == "olleh"
# assert reverse_string("") == ""

# 4. 'is_palindrome': Check if a string is a palindrome
# Accepts a string
# Returns a boolean

# assert is_palindrome("racecar") == True
# assert is_palindrome("hello") == False

# 5. 'count_occurrences': Count the occurrences of an item in a list
# Accepts a list and an item
# Returns an integer

# assert count_occurrences([1, 2, 3, 2, 1], 2) == 2
# assert count_occurrences([1, 2, 3], 4) == 0

# 6. 'find_index': Return the index of an item in a list (or -1 if not found)
# Accepts a list and an item
# Returns an integer

# assert find_index([1, 2, 3], 2) == 1
# assert find_index([1, 2, 3], 4) == -1

# 7. 'remove_duplicates': Remove all duplicates from a list
# Accepts a list
# Returns a list

# assert remove_duplicates([1, 2, 2, 3, 4, 4]) == [1, 2, 3, 4]
# assert remove_duplicates([]) == []

# 8. 'union_lists': Return the union of two lists (unique elements)
# Accepts two lists
# Returns a list

# assert union_lists([1, 2, 3], [3, 4, 5]) == [1, 2, 3, 4, 5]
# assert union_lists([], []) == []

# 9. 'is_sorted': Check if a list is sorted in ascending order
# Accepts a list
# Returns a boolean

# assert is_sorted([1, 2, 3]) == True
# assert is_sorted([3, 2, 1]) == False

# 10. 'item_counts': Return a dictionary with counts of each item in a list
# Accepts a list
# Returns a dictionary

# assert item_counts(["a", "b", "a"]) == {"a": 2, "b": 1}
# assert item_counts([]) == {}

# 11. 'sorted_keys': Return the keys of a dictionary sorted alphabetically
# Accepts a dictionary
# Returns a list

# assert sorted_keys({"b": 1, "a": 2}) == ["a", "b"]
# assert sorted_keys({}) == []

# 12. 'max_dict_value': Return the maximum value in a dictionary
# Accepts a dictionary with numeric values
# Returns a number or None

# assert max_dict_value({"a": 1, "b": 2, "c": 3}) == 3
# assert max_dict_value({}) == None

# 13. 'filter_evens': Return a list of even numbers from a list
# Accepts a list of integers
# Returns a list

# assert filter_evens([1, 2, 3, 4]) == [2, 4]
# assert filter_evens([1, 3, 5]) == []

# 14. 'merge_dicts': Merge two dictionaries
# Accepts two dictionaries
# Returns a dictionary

# assert merge_dicts({"a": 1}, {"b": 2}) == {"a": 1, "b": 2}
# assert merge_dicts({}, {}) == {}

# 15. 'flatten': Flatten a list of lists into a single list
# Accepts a list of lists
# Returns a list

# assert flatten([[1, 2], [3, 4]]) == [1, 2, 3, 4]
# assert flatten([]) == []
