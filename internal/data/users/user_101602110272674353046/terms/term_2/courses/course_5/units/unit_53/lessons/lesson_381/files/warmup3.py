# Warmup
# Write a function num_to_word that accepts a number and returns a worded version.
# Example
# example = num_to_word(9)
# print(example) # nine


def num_to_word(num: int) -> str:
    num_to_word = {
        0: "zero",
        1: "one",
        2: "two",
        3: "three",
        9: "nine",
    }
    return num_to_word[num]


assert num_to_word(2) == "two"
