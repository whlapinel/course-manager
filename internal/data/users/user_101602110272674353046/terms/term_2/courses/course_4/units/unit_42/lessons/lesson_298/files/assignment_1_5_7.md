# Unit 5 Exam

## Python CLI Application Assignment  
### Build a CLI app using:  
✅ Sets  
✅ Dictionaries  
✅ Tuples  
✅ Lists  

---

## Goal  
Create a command-line app that allows the user to:  
✅ Add data to a list  
✅ Remove data from a list  
✅ Search data using a dictionary  
✅ Store unique values using a set  
✅ Display data as tuples  

---

## Main Loop  
### Provided Code:  
```python
def main():
    data = []
    unique_items = set()
    data_map = {}

    while True:
        print("\nOptions:")
        print("1. Add item")
        print("2. Remove item")
        print("3. Search item")
        print("4. Display unique items")
        print("5. Display all items")
        print("6. Quit")

        choice = input("Enter your choice: ")

        if choice == '1':
            add_item(data, unique_items, data_map)
        elif choice == '2':
            remove_item(data, unique_items, data_map)
        elif choice == '3':
            search_item(data_map)
        elif choice == '4':
            display_unique(unique_items)
        elif choice == '5':
            display_all(data)
        elif choice == '6':
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Try again.")

if __name__ == "__main__":
    main()
```

---

## Task 1 – Add Item  
### Create `add_item()` function:  
✅ Add an item to the list  
✅ Add the item to the set if it's unique  
✅ Update the dictionary value to track **count** of the item  

Example:  
- Adding `"apple"` twice → `data_map = {"apple": 2}`  

---

## Task 2 – Remove Item  
### Create `remove_item()` function:  
✅ Remove an item from the list  
✅ If the count in the dictionary becomes zero, remove the key  
✅ If the item is no longer in the list, remove it from the set  

Example:  
- Removing `"apple"` when `data_map = {"apple": 1}` → `data_map = {}`  

---

## Task 3 – Search Item  
### Create `search_item()` function:  
✅ Allow the user to search for an item in the dictionary  
✅ If found, display the count  

Example:  
- Searching `"apple"` → `"apple found with count = 2"`  

---

## Task 4 – Display Unique Items  
### Create `display_unique()` function:  
✅ Display the set of unique items  

Example:  
- `unique_items = {"apple", "banana"}` → `"Unique items: apple, banana"`  

---

## Task 5 – Display All Items  
### Create `display_all()` function:  
✅ Convert the list into a list of tuples  
✅ Display each tuple `(index, item)`  

Example:  
- `data = ["apple", "banana", "apple"]` →  
```python
(0, 'apple')  
(1, 'banana')  
(2, 'apple')  
```  

---

## Example Output  
```plaintext
Options:
1. Add item
2. Remove item
3. Search item
4. Display unique items
5. Display all items
6. Quit
Enter your choice: 1
Enter item: apple

Options:
1. Add item
2. Remove item
3. Search item
4. Display unique items
5. Display all items
6. Quit
Enter your choice: 4
Unique items: apple
```

---

## Requirements  
✅ The app must use sets, dictionaries, tuples, and lists  
✅ Keep the code clean and readable  
✅ Include error handling for invalid inputs  

---

## Bonus  
💡 Add an option to **update an item**  
💡 Add an option to **sort the list**  
💡 Make the app **case-insensitive**  

---

## Submission  
- ✅ Upload your completed code to Canvas  
- ✅ Due tomorrow Thursday 3/20/25  

---

## ✅ Grading Criteria  

| Task | Points |
|-------|--------|
| Add Item | 20 pts |
| Remove Item | 20 pts |
| Search Item | 20 pts |
| Display Unique Items | 20 pts |
| Display All Items | 20 pts |
| **Total** | **100 pts** |
