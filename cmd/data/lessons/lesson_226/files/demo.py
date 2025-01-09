import numpy as np
import time

# Create a large list and a numpy array
size = 10000000
list1 = list(range(size))
array1 = np.array(list1)

# Time list operations
start_time = time.time()
list2 = [x * 2 for x in list1]
end_time = time.time()
list_duration = end_time - start_time
print("List operation time:", end_time - start_time)

# Time numpy array operations
start_time = time.time()
array2 = array1 * 2
end_time = time.time()
array_duration = end_time - start_time
print("Numpy array operation time:", array_duration)

times_faster = round(list_duration / array_duration, 1)
print(f"Array operation was {times_faster} times faster than list operation")
