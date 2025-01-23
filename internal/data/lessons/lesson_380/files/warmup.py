

def halve(s: str)->str:
    length_of_s = len(s)
    midpoint = length_of_s//2
    return s[0:midpoint]

print(halve("hello"))




#     # find the middle index of the string
#     middle = len(s) // 2
#     # return the first half of the string
#     return s[:middle]

# assert halve('oooo') == 'oo'
# assert halve("I don't want to be a king") == "I don't want"

# extra warmup

def name(s: str)->list[str]:
    # if not '_' in s:
    #     raise ValueError
    return s.split('_')

print(name('hello-world'))
assert name('hello_world') == ['hello', 'world']