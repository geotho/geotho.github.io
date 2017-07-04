---
layout: "post"
title: "I'm an integer overflow survivor"
date: "2017-07-03 18:48"
---

I always knew integer overflow existed, but I never thought it could happen to me.

I want to recount my cautionary tale, to avoid the same mistake in future.

The task was to write a `func isPerfectSquare(num int) bool`, with `isPerfectSquare(16) == true` and `isPerfectSquare(5) == false` etc. No standard library sqrt functions are allowed.

So I thought "I can use a binary search!" Go makes binary searching very, very easy:

```go
func isPerfectSquare(num int) bool {
	i := sort.Search(num, func(i int) bool {
		return i*i >= num
	})

	return i*i == num
}
```

And if you were to inline the sort.Search function, you'd get something like this:

```go
func isPerfectSquare(num int) bool {
	i := 0
	j := num

	for i < j {
		h := i + (j-i)/2

		if !(h*h >= num) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i*i == num
}
```

Everything worked fine!

I'm practicing for coding interviews at the moment, so I wrote it in Java too:

```java
public static boolean isPerfectSquare(int num) {
    int i = 0;
    int j = num;

    while (i < j) {
      int h = i + (j-i)/2;

      if (!(h*h >= num)) {
        i = h + 1;
      } else {
        j = h;
      }
    }

    return i*i == num;
  }
}
```

But this one failed with `isPerfectSquare(808201) == isPerfectSquare(899*899) == false`.

It was late. I thought I was crazy. "They're exactly the same!" I yelled.

The problem was this part here: `h*h`. This ends up computing ((899*899)/2)^2 which is 163,296,810,000. Which is bigger than 2^32-1. D'oh!

Why does the Go version work? Go's `int` is 64-bit on 64-bit machines. But Java's int is only ever 32-bit. Switching the `int`s to `long`s in the Java version made it work as expected.

Funnily enough, I used `int h = i + (j-i)/2` (as opposed to `int h = (i + j)/2`) to avoid integer integer overflow when calculating the middle value!

So there we are: integer overflow isn't just academic after all.
