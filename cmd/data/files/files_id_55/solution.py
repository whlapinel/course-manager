# Import the numpy library
import numpy as np

def main():
    # Exercise 1: Index into a numpy array to get the third element
    # Create an array from 1 to 10 first
    arr = np.arange(1, 11)
    third_element = arr[2]  # Remember, indexing starts at 0
    print("Third element of the array:", third_element)

    # Exercise 2: Slice the array to get the second to fifth element
    sliced_array = arr[1:5]
    print("Sliced array from second to fifth element:", sliced_array)

    # Exercise 3: Reshape the array from 1D to 2D with 2 rows and 5 columns
    reshaped_array = arr.reshape(2, 5)
    print("Reshaped array to 2D (2x5):", reshaped_array)

    # Exercise 4: Find the mean of the entire array
    mean_value = np.mean(arr)
    print("Mean of the array:", mean_value)

    # Exercise 5: Find the median of the array
    median_value = np.median(arr)
    print("Median of the array:", median_value)

    # Exercise 6: Calculate the standard deviation of the array
    std_deviation = np.std(arr)
    print("Standard deviation of the array:", std_deviation)

    # Exercise 7: Sum the elements of the reshaped 2D array across rows
    sum_across_rows = np.sum(reshaped_array, axis=1)
    print("Sum across rows:", sum_across_rows)

    # Exercise 8: Calculate the maximum value of each column in the 2D array
    max_per_column = np.max(reshaped_array, axis=0)
    print("Maximum value per column:", max_per_column)

if __name__ == "__main__":
    main()
