# Go Learn By Doing 🐹

A hands-on, practice-based course to learn Go, from beginner to building a REST API.
No videos. You write code, run the tests, and push to GitHub. Every lesson has a test that fails until you solve it.

## Who is this for?

- Complete beginners who want to learn Go by writing code
- Developers from other languages who want to learn Go quickly
- Anyone who wants to practice Git and GitHub while learning

## How to start

1. Click **Use this template**, then **Create a new repository**.
2. Clone your new repository:

```bash
   git clone https://github.com/<your-username>/<your-repo-name>.git
   cd <your-repo-name>
```

3. Install Go from https://go.dev/dl and check it:

```bash
   go version
```

4. Open the first lesson and read its `README.md`:
   [Lesson 01: Hello World](lessons/1-beginner/01-hello-world/README.md)
5. Edit `main.go` and run the test:

```bash
   go test ./lessons/1-beginner/01-hello-world/
```

6. When the test passes, commit and push. Open the **Actions** tab to see GitHub check your work (green tick = pass).
7. Tick the lesson in the checklist below (change `[ ]` to `[x]`) in your own copy of this README, then commit that too.

## How a lesson works

```
Read README → Edit main.go → go test → Fix → Pass → Commit & push → Tick checklist → Next lesson
```

## Course levels

| Level | Topics |
|-------|--------|
| 1-beginner | Hello world, variables, conditionals, loops, functions |
| 2-intermediate | Slices, maps, structs, methods, interfaces, errors, pointers |
| 3-advanced | Goroutines, channels, context, generics, testing |
| 4-rest-api | net/http, JSON, routing, middleware, database, Docker |

## Progress checklist

Change `[ ]` to `[x]` when you finish a lesson.

### Beginner
- [ ] [01 - Hello World](lessons/1-beginner/01-hello-world/README.md)
- [ ] [02 - Variables and Types] (lessons/1-beginner/02-variables/README.md)
- [ ] 03 - Conditionals
- [ ] 04 - Loops
- [ ] 05 - Functions

### Intermediate
- [ ] Coming soon

### Advanced
- [ ] Coming soon

### REST API
- [ ] Coming soon

## Stuck?

Look at the `solutions` branch of the original repo, but try on your own first.

## Contributing

Found a mistake or want to add a lesson? Open an issue or a pull request.

## License

MIT