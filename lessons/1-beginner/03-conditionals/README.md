# Lesson 03: Conditionals

## Goal

Fix the `Grade` function so it returns the correct grade based on a student's score.

The function should follow these rules:

| Score | Grade |
|-------|-------|
| 90-100 | `A` |
| 75-89 | `B` |
| 60-74 | `C` |
| 40-59 | `D` |
| 0-39 | `F` |
| Below 0 or above 100 | `Invalid` |

## What you will learn

- `if` statements
- `else if`
- `else`
- Comparison operators
- Logical operators
- Returning values from conditional branches

## Steps

1. Open `main.go`.
2. Find the `TODO` inside the `Grade` function.
3. Use `if`, `else if`, and `else` to implement the grading rules.
4. Run the test:

```bash
go test ./lessons/1-beginner/03-conditionals/
```

5. Keep fixing your code until the test passes.
6. Run the program:

```bash
go run ./lessons/1-beginner/03-conditionals/
```

7. When the test passes, commit and push:

```bash
git add .
git commit -m "finish lesson 03"
git push
```

## Examples

```go
if score >= 90 {
    return "A"
} else if score >= 75 {
    return "B"
} else {
    return "C"
}
```

You can also combine conditions with logical operators:

```go
if score < 0 || score > 100 {
    return "Invalid"
}
```

### Important

Go does **not** require parentheses around the condition:

```go
if score >= 90 {
    // correct
}
```

Not:

```go
if (score >= 90) {
    // unnecessary in Go
}
```

## Hint

A good way to start is to check whether the score is invalid first:

```go
if score < 0 || score > 100 {
    return "Invalid"
}
```

Then check the grades from the highest score to the lowest.

## Challenge

After the test passes, try creating your own conditional function.

For example:

```go
func NumberType(n int) string
```

Make it return:

- `Positive` for numbers greater than 0
- `Negative` for numbers less than 0
- `Zero` for 0

Don't change the existing `Grade` function. Create the new function yourself and write tests for it.