# Import the numpy library
import numpy as np


def main():
    # Exercise 1: Create a 3D numpy array and print it
    # Create a 3D array with shape (2, 3, 4) filled with random numbers
    arr_3d = np.random.rand(2, 3, 4)
    print("3D Array:")
    print(arr_3d)

    # Exercise 2: Perform boolean indexing on a 2D array
    # Create a 2D array and filter out elements less than 0.5
    arr_2d = np.random.rand(5, 5)
    filtered_arr = arr_2d[arr_2d < 0.5]
    print("Filtered Array (Elements < 0.5):")
    print(filtered_arr)

    # Exercise 3: Reshape an array from 1D to 3D and print the new array
    arr_1d = np.arange(24)
    reshaped_arr = arr_1d.reshape(2, 3, 4)
    print("Reshaped Array from 1D to 3D:")
    print(reshaped_arr)

    # Exercise 4: Calculate and print the variance and standard deviation of a 1D array
    data = np.random.rand(10)
    variance = np.var(data)
    std_dev = np.std(data)
    print("Variance of the array:", variance)
    print("Standard deviation of the array:", std_dev)

    # Exercise 5: Calculate the sum across columns of a 2D array
    print("Original 2D array:")
    print(arr_2d)
    sum_columns = np.sum(arr_2d, axis=0)
    print("Sum across columns:")
    print(sum_columns)

    # Exercise 6: Calculate and print the correlation coefficient between two 1D arrays
    x = np.random.rand(10)
    y = 3 * x + np.random.normal(0, 0.1, 10)
    correlation = np.corrcoef(x, y)
    print("Correlation coefficient between x and y:")
    print(correlation)


if __name__ == "__main__":
    main()
