---
layout: "post"
title: "You don't always have to change that function signature"
date: "2017-03-25 12:52"
---

Sometimes, I see functions that have some hardcoded parameters and some that are passed in:

```go
func NewFoo(size int) {
  return Foo{
    Size: size,
    Name: "Billy",
  }
}
```

And the `NewFoo` function is used in many places.

But now imagine that rather than "Billy", we want to pass in the value for name.

One option is to change `NewFoo` to take a size and a name, like this:

```go
func NewFoo(size int, name string) {
  return Foo{
    Size: size,
    Name: name,
  }
}
```

But then you have to update each usage of `NewFoo`. Not ideal.

Instead, you can do something like this:

```go
func NewFooWithName(size int, name string) {
  return Foo{
    Size: size,
    Name: name,
  }
}

func NewFoo(size int) {
  return NewFooWithName(size, "Billy")
}
```

and leave all the existing `NewFoo`s as they are.

Note that there is a limit to how sensible this technique is. If you start doing function calls three or four deep just to specify different parameters then maybe it is time to refactor. But for simple things like this I would opt for changing code in the least number of places.
