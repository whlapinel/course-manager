import numpy as np

# # Reshape from 1D to 2D
arr = np.arange(10)  # Creates an array with numbers from 0 to 9
# reshaped_arr = arr.reshape(2, 5)
# print(reshaped_arr)
# print(arr.shape)


# # Statistics
# data = np.array([1, 2, 3, 4, 5])

# print("Mean:", np.mean(data))
# print("Median:", np.median(data))
# print("Standard Deviation:", np.std(data))

# matrix = np.array([[1, 2, 3], [4, 5, 6]])

# # Sum across rows
# print("Sum of each column:", np.sum(matrix, axis=0))

# # Sum across columns
# print("Sum of each row:", np.sum(matrix, axis=1))

print(np.shape(arr))
reshaped = arr.reshape(2, 5)
print(reshaped)
print(reshaped.shape)
