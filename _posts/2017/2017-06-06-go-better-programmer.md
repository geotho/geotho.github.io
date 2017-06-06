---
layout: "post"
title: "Ways Go made me a better programmer"
date: "2017-06-06 12:39"
---

Language influences thought, and programming languages are no different. Writing code in a new language makes you frame the problem in fresh way. Go has narrowed my focus in a few areas, and I feel like a better programmer as a result. Some of the features of Go that have helped me improve are below.

### Interfaces
Interfaces in Go are easy to create and to implement. You can make any type implement any interface by creating a custom type and attaching methods to it. You don’t explicitly declare that the type implements the interface, as you do in e.g. Java, but you can still assert this at compile time.

The Go mantra "functions should expect interfaces and return structs" drums home this idea that you don’t actually care about the thing you are getting, you only care about what it can do. This requires you to focus on the narrowest set of required functionality. The standard library’s `sort` package provides a beautiful illustration.

### Style guide
The standard library code formatter, `gofmt`, is a revelation. Thank goodness I no longer need to spend time formatting code, or deciding whether to use tabs or spaces, or anything else that doesn’t matter as long as it’s consistent.

But it goes further than that: [Effective Go](https://golang.org/doc/effective_go.html) provides details on how to write Go that cannot be captured by an automatic formatter. Tips like ["keep the main logic unindented"](https://github.com/golang/go/wiki/CodeReviewComments#indent-error-flow) help ensure code written by different people is consistent and readable.

Go optimises for consistency and readability. You and your team will be able to understand each other’s code faster and the developing-reviewing-maintaining lifecycle benefits.

### Opinionated – there’s only one way to code
Go’s language design is deliberately simple. It lacks generics. It lacks while loops. It lacks classes. If you are a Lisp programmer, or have read [Paul Graham’s blog post](http://www.paulgraham.com/avg.html) on it, you might be tempted to ask "How can you get anything done in Go? It doesn't even have {generics, while loops, classes, metaprogramming...}."

I’ve used languages which are more powerful (i.e. more abstract) than Go. But I feel more productive in Go than in these other languages. Why? I find the simplicity (i.e. lack of power) refreshing. Power comes with choice. And choice leads to decision making. And decision making leads to procrastination and disagreement.

How many times have I started writing a for loop in Python, until I remember than I can also write a list comprehension to accomplish the same thing? Or the map function with a lambda? Any time I have to ask myself "should I use language feature X or Y to accomplish this same goal?" is a source of programming impedance, and a chance to ruin flow.

I remember another time where one engineer used a Java 8 lambda when a more senior engineer replaced it with a for loop. How much time could have been saved here if only the for loop existed in the first place? How much time and decision fatigue might be saved in aggregate?

Go avoids these problems by only have one way to accomplish something. Want to create a new list where each element is the square of some other list? Write a loop. Want to filter some of these elements? Write a loop.

I found that by narrowing the amount of ways I could accomplish the same goal, and thus removing minor decisions I had to make during day-to-day coding, I found a sense of flow much easier.

### Concurrency
Go’s concurrency primitives are built into the language and are first class citizens of it. I’ve never had concurrent programming make so much sense to me as when I use Go’s `sync` package. I found it far easier to inject performance-enhancing concurrency into my programs in Go than in e.g. Java.

### Table-driven testing
Table-drive tests contain a list of (input, expected output) pairs. This enables unit tests which:
- Force modular, unit-testable code.
- Focus on the test data, not the extraneous mocking etc. that is usually required for testing.
- Clearly define and document the function under test.

It is Go’s structs that make table-driven tests easier to write. Table driven tests in Java would require another testcase class, which always seems more tedious to write and to use.
