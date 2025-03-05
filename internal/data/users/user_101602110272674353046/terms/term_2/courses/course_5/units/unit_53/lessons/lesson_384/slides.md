---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Lesson 3: Relating Tables in SQL with Python
## Foreign Keys & JOIN Statements

## 🏗 Why Relate Tables?
- A **database** is more than just one table.
- Relating tables prevents **data duplication**.
- Example: A Clash Royale game might store **players** and their **decks** in separate tables.
- To connect related data, we use **foreign keys** and **JOIN statements**.

# 🔗 What is a Foreign Key?
A **foreign key** is a column that links to the **primary key** of another table.

## Example: Players & Decks
```sql
CREATE TABLE players (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE decks (
    id INTEGER PRIMARY KEY,
    player_id INTEGER,
    name TEXT NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id)
);
```
- **`players.id`** is the **primary key** in `players`.
- **`decks.player_id`** is a **foreign key**, linking each deck to a player.

# 🎴 Inserting Related Data
```sql
INSERT INTO players (name) VALUES ('Jeanette');
INSERT INTO players (name) VALUES ('Jan');

INSERT INTO decks (player_id, name) VALUES (1, 'Jeanette’s Mega Deck');
INSERT INTO decks (player_id, name) VALUES (2, 'Jan’s Royal Deck');
```
- Jeanette’s deck links to **player_id = 1**.
- Jan’s deck links to **player_id = 2**.

# 🔍 Retrieving Related Data (Without JOIN)
We can manually look up a player's deck in Python:
```python
cursor.execute("SELECT * FROM decks WHERE player_id = ?", (1,))
decks = cursor.fetchall()
```
🚨 **Problem:** This requires multiple queries and is inefficient!

# 🔄 The Power of JOIN
**JOIN** allows us to combine data from multiple tables in **one query**.

## Example: Players & Their Decks
```sql
SELECT players.name, decks.name
FROM players
JOIN decks ON players.id = decks.player_id;
```
This links `players` and `decks`, showing **each player's deck**.

# 🎮 Running a JOIN in Python
```python
cursor.execute('''
    SELECT players.name, decks.name 
    FROM players
    JOIN decks ON players.id = decks.player_id
''')
results = cursor.fetchall()
for row in results:
    print(f"Player: {row[0]}, Deck: {row[1]}")
```
✅ **One query, efficient retrieval!**

# 🔄 Types of JOINs

| JOIN Type | Description |
|-----------|-------------|
| **INNER JOIN** | Returns only matching rows (most common) |
| **LEFT JOIN**  | Returns all rows from the left table, even if no match |
| **RIGHT JOIN** | Returns all rows from the right table, even if no match |
| **FULL JOIN**  | Returns all rows from both tables (not supported in SQLite!) |

# 🔄 Using LEFT JOIN
If some players **don’t have decks**, we still want to show them:
```sql
SELECT players.name, decks.name
FROM players
LEFT JOIN decks ON players.id = decks.player_id;
```
Now players **without decks** will still appear (with NULL deck values).

# 🛠 Practice Task: Finding Player Decks
## Given these tables:
```sql
CREATE TABLE players (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE decks (
    id INTEGER PRIMARY KEY,
    player_id INTEGER,
    name TEXT NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id)
);
```
❓ **Write a JOIN query to list all players and their decks.**

# 🚀 Wrapping Up
- **Foreign keys** link tables together.
- **JOIN statements** efficiently combine data.
- **Different types of JOINs** serve different needs.
