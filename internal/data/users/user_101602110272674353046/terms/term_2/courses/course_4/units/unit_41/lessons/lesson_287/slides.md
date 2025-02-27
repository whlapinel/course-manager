---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Phones in bags, bags against the wall
- Create and open `lesson_4` in VS Code

```text
/python
  /unit_4
    /lesson_4
```

# More on Functions in Python 

## Review: What is a Function?
- A function is a block of reusable code.
- Helps to organize and avoid repetition.
- Defined using `def` keyword.

# Function Parameters
### 1. Positional Parameters
- Parameters are assigned values in the order they appear.
```python
def pokemon_battle(trainer, pokemon):
    print(f"{trainer} sends out {pokemon}!")

pokemon_battle("Ash", "Pikachu")
```

# 2. Default Parameters
- Assigns a default value if no argument is provided.
```python
def catch_pokemon(pokemon, pokeball="Poké Ball"):
    print(f"You caught {pokemon} using a {pokeball}!")

catch_pokemon("Charmander")  # Uses default Poké Ball
```

# 3. Keyword Arguments
- Parameters can be specified by name.
```python
def pokemon_stats(name, type, level):
    print(f"{name} is a {type}-type Pokémon at level {level}!")

pokemon_stats(level=25, name="Bulbasaur", type="Grass/Poison")
```

# 4. Arbitrary Arguments (`*args`)
- Allows multiple positional arguments.
```python
def choose_team(*pokemons):
    print("Your Pokémon team:")
    for pokemon in pokemons:
        print(f"- {pokemon}")

choose_team("Squirtle", "Jigglypuff", "Gengar")
```

# 5. Arbitrary Keyword Arguments (`**kwargs`)
- Allows multiple keyword arguments.
```python
def pokemon_profile(**details):
    print("Pokémon Profile:")
    for key, value in details.items():
        print(f"{key}: {value}")

pokemon_profile(name="Eevee", type="Normal", evolves_into="Multiple Forms")
```

# Practice Challenges
## Challenge 1:
Write a function that takes a Pokémon’s name and its region (default: Kanto) and prints its origin.

# Challenge 2:
Modify `pokemon_stats` to accept multiple Pokémon using `*args` and display their stats.

# Challenge 3:
Create a function that accepts a Pokémon’s profile with `**kwargs` and prints their details.

# Summary
- Different types of parameters help create flexible functions.
- `*args` and `**kwargs` allow handling multiple arguments dynamically.
- Default parameters provide fallbacks when values are missing.
- Using keyword arguments makes functions more readable.

# Assignment 4.4

- Complete [problems](./files/assignment_1_4_4.html) and submit as python file
- Complete CodeHS Assignments 5.5 & 5.6

## Reminder
 
- Python Essentials 1 Course Module 2 Quiz is due tomorrow