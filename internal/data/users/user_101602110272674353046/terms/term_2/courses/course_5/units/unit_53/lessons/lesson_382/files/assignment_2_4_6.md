# SQLite Assignment: Minecraft Villager Trading Database

## **Objective**
You are a Minecraft villager librarian who manages trade records using SQLite. Your goal is to retrieve, update, and secure trade data using advanced SQLite techniques in Python.

## **Setup**
1. **Create a Database**: Connect to `villager_trades.db`.
2. **Create a Table**: Define a table `trades` with the following structure:
   ```sql
   CREATE TABLE trades (
       id INTEGER PRIMARY KEY,
       villager_name TEXT,
       item TEXT,
       emerald_cost INTEGER
   );
   ```
3. **Insert Sample Data**:
   ```python
   trades = [
       ("Librarian", "Enchanted Book", 12),
       ("Fletcher", "Bow", 8),
       ("Farmer", "Golden Carrot", 3)
   ]
   cursor.executemany("INSERT INTO trades (villager_name, item, emerald_cost) VALUES (?, ?, ?)", trades)
   conn.commit()
   ```

---

## **Tasks**

### **1. Retrieve Trades Using `row_factory`**
Modify your connection to return results as dictionaries:
```python
conn.row_factory = sqlite3.Row
```
Fetch and print all trades using column names instead of indexes.

### **2. Prevent SQL Injection**
Write a function `get_trades_by_villager(villager_name)` that safely retrieves all trades for a given villager using **parameterized queries**.
```python
def get_trades_by_villager(name):
    pass # insert your code here
```
Test it with safe user input.

### **3. Use `with` for Connection Management**
Rewrite your database interaction using a `with` statement to ensure proper resource handling.

### **4. Handle Errors Gracefully**
Wrap this query inside a `try-except` block to catch potential SQLite errors.
```python
cursor.execute("SELECT * FROM nonexistent_table")
```

### **5. Efficient Data Fetching**
- Fetch only one trade (`fetchone()`)
- Fetch all trades (`fetchall()`)
- Fetch two trades at a time (`fetchmany(2)`) and print them.

---

## **Bonus Challenge**
Create a function `update_trade_cost(villager, item, new_cost)` that safely updates the emerald cost of an item.
