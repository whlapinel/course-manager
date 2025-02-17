import numpy as np

import pandas as pd

data = {
    "Name": ["John", "Anna", "James", "Linda"],
    "Age": [28, 22, 35, 32],
    "City": ["New York", "Paris", "Berlin", "London"],
}
df = pd.DataFrame(data)
print(df)

df.to_csv("hello.csv")

df["Salary"] = np.random.randint(50000, 100000, size=(4))

df["Salary"] = [80000, 20000, 30000, 10000]

print(df)

df.to_csv("test.csv")
