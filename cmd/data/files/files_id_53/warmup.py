# Write a function `write_name` that writes a name to a text file `names.txt`.

# ```python
# def write_name(name: str)->None:
# ```

# Write a second function `read_name` that reads and returns the first name from the text file `names.txt`.

# ```python
# def read_name()->str:
# ```


def write_name(name: str):
    with open("names.txt", "w") as file:
        file.write(name)

def read_name() -> str:
    with open("names.txt", "r") as file:
        contents = file.read()
        return contents


write_name("Bill")
name = read_name()
print(name)
assert name == "Bill"
