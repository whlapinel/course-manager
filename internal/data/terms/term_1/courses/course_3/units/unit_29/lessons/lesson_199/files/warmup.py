# Write a function abbreviated that accepts a string and returns a new string that
# consists of the last 5 characters.

def abbreviated(s: str)->str:
    return s[-5:]

# format for slicing:
# string[start:stop:step]

modified_string = abbreviated('my favorite class is Python')
print(modified_string) # ython