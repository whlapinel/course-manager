---
layout: none
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Go to [VS Code Keyboard Shortcuts](https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf)
  - Identify 3 *new* keyboard shortcuts and try them out
  - Be prepared to share one

# Introduction to Binary Numbers and Bitwise Operations in Python

# What is the Binary Number System?
- Binary is a **base-2** number system.
- It uses only two digits: **0 and 1**.
- Each position represents a power of **2** (e.g., `2^0`, `2^1`, `2^2`, etc.).

# Example: Converting Binary to Decimal
- Example: `1011` (Binary) → **?** (Decimal)

## Steps:
1. `(1 × 2³) + (0 × 2²) + (1 × 2¹) + (1 × 2⁰)`
2. `8 + 0 + 2 + 1 = 11`
- **Answer:** `1011₂ = 11₁₀`

# Example: Converting Decimal to Binary
- Convert `19` (Decimal) to Binary

## Steps:
1. Divide by 2, record the remainder.
2. Continue until quotient is `0`.
3. Read remainders **bottom to top**.

```text
19 ÷ 2 = 9 remainder ?
 9 ÷ 2 = 4 remainder ?
 4 ÷ 2 = 2 remainder ?
 2 ÷ 2 = 1 remainder ?
 1 ÷ 2 = 0 remainder ?
```
# Answer

- **Answer:** `19₁₀ = 10011₂`

# Bitwise Operations in Python

## What are Bitwise Operations?
- Bitwise operations manipulate bits directly.
- They are useful in low-level programming, cryptography, networking, and performance optimization.

# Bitwise Operators in Python

| Operator | Symbol | Example | Explanation |
|----------|--------|---------|-------------|
| AND | `&` | `5 & 3` | Bitwise AND |
| OR | `\|` | `5 \| 3` | Bitwise OR |
| XOR | `^` | `5 ^ 3` | Bitwise XOR |
| NOT | `~` | `~5` | Bitwise NOT |
| Left Shift | `<<` | `5 << 1` | Shift bits left |
| Right Shift | `>>` | `5 >> 1` | Shift bits right |

# Bitwise AND (`&`)
- Compares corresponding bits, returns **1 if both are 1**, otherwise **0**.
- Example: `5 & 3`

```text
  5 = 0b0101
  3 = 0b0011
 -------------
      0b0001 (1 in decimal)
```
# Solution
- **Answer:** `5 & 3 = 1`

# Bitwise OR (`|`)
- Compares corresponding bits, returns **1 if either bit is 1**.
- Example: `5 | 3`

```text
  5 = 0b0101
  3 = 0b0011
 -------------
      0b0111 (7 in decimal)
```

# Solution
- **Answer:** `5 | 3 = 7`

# Bitwise XOR (`^`)
- Compares corresponding bits, returns **1 if bits are different**, otherwise **0**.
- Example: `5 ^ 3`

```text
  5 = 0b0101
  3 = 0b0011
 -------------
      0b0110 (6 in decimal)
```

# Solution
- **Answer:** `5 ^ 3 = 6`

# Bitwise NOT (`~`)
- Flips all bits: **1 → 0 and 0 → 1**.
- Example: `~5`

```text
  5 =  0b0000 0101
 ~5 =  0b1111 1010 (Two's complement: -6)
```

# Solution
- **Answer:** `~5 = -6`

# Left Shift (`<<`)
- Shifts bits **left**, filling with `0`s on the right.
- Equivalent to multiplying by `2^n`.

Example: `5 << 1`
```text
  5 = 0b0101
 -------------
 10 = 0b1010
```
# Solution
- **Answer:** `5 << 1 = 10`

# Right Shift (`>>`)
- Shifts bits **right**, filling with `0`s on the left.
- Equivalent to integer division by `2^n`.

Example: `5 >> 1`
```text
  5 = 0b0101
 -------------
  2 = 0b0010
```
# Solution
- **Answer:** `5 >> 1 = 2`

# Summary
- **Binary numbers** are base-2.
- **Bitwise operators** allow direct bit manipulation.
- **AND, OR, XOR, NOT, Shift** operations are key tools in Python.
- Useful for **performance optimization, cryptography, networking, and low-level programming**.

# Other resources

[YouTube video: FireShip explains Binary in 01100100 Seconds](https://youtu.be/zDNaUi2cjv4?si=LH4GN2UtdLhtmn9k)

# Assignment 2.5

- Solve [Problems](./files/assignment_2_2_5.html) and submit as python file in Canvas
- Continue working on Python Essentials 2 Course Module 1 (Module 1 Quiz is due tomorrow)
