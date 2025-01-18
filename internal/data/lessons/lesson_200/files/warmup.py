grades = {
"Kurt": 86,
"Russell": 67,
"Leon": 78
}

print(grades.values())
# Output: dict_values([86, 67, 78])

print(grades.keys())
# Output: dict_keys(['Kurt', 'Russell', 'Leon'])

print(grades.items())
# Output: dict_items([('Kurt', 86), ('Russell', 67), ('Leon', 78)])

print(grades["Kurt"])
# Output: 86

print(grades.get("Russell"))
# Output: 67

print(len(grades))
# Output: 3

print(sum(grades.values()))
# Output: 231

def grade_avg(grades):
    total = 0
    for grade in grades.values():
        total += grade
    return total / len(grades)

avg = grade_avg(grades)
print(avg) # 77.0