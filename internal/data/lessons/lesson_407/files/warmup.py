import numpy as np

arr = np.random.randint(1, 101, size=(100))

print(arr)

filt = arr < 10

print(filt)

filtered_array = arr[filt]

print(filtered_array)
