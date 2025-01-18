def num_to_word(num: int) -> str:
    num_word_dict = {
        1: "one",
        2: "two",
        3: "three",
        4: "four",
    }
    converted = num_word_dict[num]
    return converted


# example = num_to_word(9)
# print(example)  # nine

assert num_to_word(4) == "four"
