# Practice 1: Find Maximum in a List
def find_max(lst):
    pass


# Practice 2: Palindrome Check
def is_palindrome(s):
    pass


# Practice 3: Count Vowels in a String
def count_vowels(s):
    pass


# Practice 4: Flatten a Nested Dictionary
def flatten_dict(d, parent_key=""):
    pass


# Practice 5: Generate All Subsets of a List
def subsets(lst):
    pass


# Run sample tests
if __name__ == "__main__":
    print("Find Max:", find_max([3, 1, 4, 1, 5, 9, 2, 6]))
    print("Is 'racecar' a palindrome?:", is_palindrome("racecar"))
    print("Count vowels in 'recursion':", count_vowels("recursion"))

    nested_dict = {"a": 1, "b": {"c": 2, "d": {"e": 3}}}
    print("Flattened Dictionary:", flatten_dict(nested_dict))

    print("Subsets of [1, 2, 3]:", subsets([1, 2, 3]))
