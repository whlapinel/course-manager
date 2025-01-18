# This was created to illustrate how to read a file backwards. This is the challenge in CodeHS 15.6.7 "Reverse a File"

# Read the file
file = open("data.txt")
file.seek(0, 2)
size = file.tell()

# List to hold characters
backwards = []

# For loop to go from the end of the file to the beginning
for i in range(size, -1, -1):
    file.seek(i)
    char = file.read(1)
    backwards.append(char)

file.close()

# Put all the characters in a string
backwards_string = "".join(backwards)


file_new = open("new_data.txt", "w")
file_new.write(backwards_string)

file_new.close()
