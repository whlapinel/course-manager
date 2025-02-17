# # 12.1 Tuples
# # Q1: Write a function `sum_tuple` that takes a tuple of numbers and returns the sum of all its elements.

# # Sample Test
# assert sum_tuple((1, 2, 3)) == 6  # Expected: 6
# assert sum_tuple((5, 10, 15)) == 30  # Expected: 30

# # 12.2 Lists
# # Q2: Write a function `max_list` that takes a list of numbers and returns the largest number.

# # Sample Test
# assert max_list([1, 2, 3, 4, 5]) == 5  # Expected: 5
# assert max_list([-10, -5, 0, 5]) == 5  # Expected: 5

# # 12.3 For Loops and Lists
# # Q3: Write a function `count_evens` that takes a list of integers and returns the count of even numbers.

# # Sample Test
# assert count_evens([1, 2, 3, 4, 5, 6]) == 3  # Expected: 3
# assert count_evens([1, 3, 5, 7]) == 0  # Expected: 0

# # 12.4 List Methods
# # Q4: Write a function `remove_first_occurrence` that removes the first occurrence of a given value from a list.

# # Sample Test
# assert remove_first_occurrence([1, 2, 3, 2, 4], 2) == [1, 3, 2, 4]  # Expected: [1, 3, 2, 4]
# assert remove_first_occurrence([5, 6, 7], 8) == [5, 6, 7]  # Expected: [5, 6, 7]

# # 12.5 Creating and Altering Data Structures
# # Q5: Write a function `create_number_list` that takes two integers, `start` and `end`, and returns a list of numbers between them (inclusive).

# # Sample Test
# assert create_number_list(1, 5) == [1, 2, 3, 4, 5]  # Expected: [1, 2, 3, 4, 5]
# assert create_number_list(3, 3) == [3]  # Expected: [3]


# # 13.1 2D Lists
# # Q6: Write a function `row_sum` that takes a 2D list (list of lists) and an integer `row` and returns the sum of the elements in that row.


# # Sample Test
# assert row_sum([[1, 2, 3], [4, 5, 6], [7, 8, 9]], 1) == 15  # Expected: 15
# assert row_sum([[10, 20], [30, 40], [50, 60]], 2) == 110  # Expected: 110

# # 13.2 List Comprehensions
# # Q7: Write a function `square_odds` that takes a list of integers and returns a new list with the squares of all the odd numbers.

# # Sample Test
# assert square_odds([1, 2, 3, 4, 5]) == [1, 9, 25]  # Expected: [1, 9, 25]
# assert square_odds([2, 4, 6]) == []  # Expected: []

# # 13.3 Packing and Unpacking
# # Q8: Write a function `swap_first_last` that takes a list and returns a new list with the first and last elements swapped.

# # Sample Test
# assert swap_first_last([1, 2, 3, 4]) == [4, 2, 3, 1]  # Expected: [4, 2, 3, 1]
# assert swap_first_last([7]) == [7]  # Expected: [7]

# # 13.4 Dictionaries
# # Q9: Write a pure function `invert_dict` that takes a dictionary and returns a new dictionary where keys and values are swapped.

# # Sample Test
# assert invert_dict({"a": 1, "b": 2, "c": 3}) == {1: "a", 2: "b", 3: "c"}  # Expected: {1: "a", 2: "b", 3: "c"}
# assert invert_dict({}) == {}  # Expected: {}

# # 13.5 Extending Data Structures
# # Q10: Write a pure function `merge_dicts` that takes two dictionaries and merges them, with the second dictionary overwriting values of the first for matching keys.

# # Sample Test
# assert merge_dicts({"a": 1, "b": 2}, {"b": 3, "c": 4}) == {"a": 1, "b": 3, "c": 4}  # Expected: {"a": 1, "b": 3, "c": 4}
# assert merge_dicts({}, {"x": 42}) == {"x": 42}  # Expected: {"x": 42}
