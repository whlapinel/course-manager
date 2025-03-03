# Assignment 4.6

## Instructions

- Complete the practice problems below and submit as a python file
- Complete CodeHS 7.3 and 7.4

## Python Function Practice Problems 🧙🏻‍♂️

---

### 1. The Potion Mixer (Function Parameters & Return Values)
Write a function `brew_potion(ingredient1, ingredient2)` that takes two ingredients as arguments and returns a potion recipe as a string.

```python
def brew_potion(ingredient1, ingredient2):
    # Your code here

print(brew_potion("Eye of Newt", "Dragon Scale"))
# Output: "A potion made of Eye of Newt and Dragon Scale."
```

---

### 2. The Cursed Scroll (Local Scope)
Inside the function `read_scroll()`, define a local variable `spell` with a secret incantation. Print it inside the function but try printing it outside—what happens?

```python
def read_scroll():
    spell = "Wingardium Leviosa"
    print(spell)

read_scroll()
print(spell)  # What happens?
```

---

### 3. The Wand of Power (Global Scope)
The wizard’s wand stores its power in a global variable `wand_power`. Create a function `cast_spell()` that modifies this global variable.

```python
wand_power = 50

```

---

### 4. The Spellbook (Enclosing Scope)
A spellbook contains spells written inside a nested function. Modify the `change_spell()` function so that it updates the spell inside `spellbook()`.

```python
def spellbook():
    spell = "Fireball"
    
    def change_spell():
        spell = "Ice Blast"
    
    change_spell()
    print(spell)  # Output: "Ice Blast"

spellbook()
```

---

### 5. The Forbidden Rune (Shadowing)
Fix the issue in `use_rune()` where the local variable `rune` is shadowing the global `rune`.

```python
rune = "Ancient Power"

def use_rune():
    rune = "Fading Magic"  # This shadows the global variable
    print(rune)  # Output should be "Ancient Power"

use_rune()
```

---

### 6. The Magic Counter (Common Scope Mistakes)
Fix the `count_spells()` function so it correctly updates the `spell_count` global variable.

```python
spell_count = 0

def count_spells():
    spell_count += 1  # What's wrong here?

count_spells()
print(spell_count)
```

---

### 7. (Bonus) The Ultimate Challenge: Summoning the Python Guardian
Write a function `summon_guardian()` that uses both `global` and `nonlocal` to modify a spell stored in an enclosing function while keeping track of a global energy reserve.

```python
energy = 100

def spell_circle():
    spell = "Basic Shield"
    
    def summon_guardian():
        pass
    
    summon_guardian()
    print(spell)
    
spell_circle()
print(energy)
```
