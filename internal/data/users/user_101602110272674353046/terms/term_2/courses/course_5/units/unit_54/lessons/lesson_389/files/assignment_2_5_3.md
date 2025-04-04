# Assignment 5.3: Generators

1. **Create a simple generator**  
   Write a generator function `countdown(n)` that yields numbers from `n` down to `1`.

2. **Create a generator for file processing**  
   Write a generator `read_lines(file)` that reads lines from a file one at a time and yields each line.  

3. **Create a generator for even numbers**  
   Write a generator `even_numbers(n)` that yields even numbers from `0` to `n`.

4. **Filter using a generator**  
   Write a generator `filter_long_words(words, n)` that yields only the words longer than `n` characters from a list.

5. **Combine generator values**  
   Create two generators:  
   - `gen1` yields numbers from 1 to 5  
   - `gen2` yields letters from `'a'` to `'e'`  
   Write a generator `combine(gen1, gen2)` that yields values from both generators alternately.  

6. **Repeat values from a generator**  
   Write a generator `repeat_values(gen, times)` that repeats each value from another generator `times` times.

7. **Use a generator expression**  
   Create a generator expression that yields the uppercase version of each string in a list.

8. **Create an infinite generator**  
   Write an infinite generator `cycle_words(words)` that cycles through a list of words endlessly.

9. **Create a generator with `send()`**  
   Write a generator `counter(start)` that starts at `start` and increments by 1 each time you call `send()`.

10. **Create a generator with cleanup**  
    Write a generator `open_file(file)` that yields each line from a file and closes the file automatically when done (use `finally`).
