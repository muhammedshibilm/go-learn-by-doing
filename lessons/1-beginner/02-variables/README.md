# Lesson 02: Variables and Types

## Goal
Fix the `Profile` function so it returns these four values:

| Value | Type | Must be |
|-------|------|---------|
| name | string | `"Gopher"` |
| age | int | `10` |
| height | float64 | `1.75` |
| isStudent | bool | `true` |

## Steps
1. Open `main.go`.
2. Inside `Profile`, create four variables and return them.
3. Run the test:
```bash
   go test ./lessons/1-beginner/02-variables/
```
4. When it passes, commit and push:
```bash
   git add .
   git commit -m "finish lesson 02"
   git push
```

## Hint
Two ways to create a variable:
```go
var city string = "Kerala"   // long form, with the type
country := "India"           // short form, Go works out the type
```
The short form `:=` only works inside functions.