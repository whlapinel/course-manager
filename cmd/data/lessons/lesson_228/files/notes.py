import numpy


arr = numpy.random.randint(1, 10, (3, 3))

print(arr)

print(numpy.mean(arr, 0))
print(numpy.mean(arr, 1))

print(numpy.sum(arr, 0))
print(numpy.sum(arr, 1))

print(numpy.median(arr, 0))
print(numpy.median(arr, 1))
