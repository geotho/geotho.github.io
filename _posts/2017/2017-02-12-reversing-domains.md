---
layout: "post"
title: "Whiteboarding coding in the wild: reversing domains for BigTable"
date: "2017-02-12 12:24"
---

A common complaint about whiteboard coding questions is that they do not resemble a software engineer’s day-to-day work.
So I like to note the times when these fun toy algorithms have actually come in handy. 

A classic whiteboard coding question is as follows: 

> "Given a string containing a sentence, reverse the order of the words in that sentence."

So "the quick brown fox" should become "fox brown quick the".

This is the same problem as reversing domain components for Google BigTable keys. Given some row keys that should contain domains like "foo.example.co.uk" and "bar.example.co.uk", it is wise to store them in component reverse order e.g. as "uk.co.example.foo" and "uk.co.example.bar". That way, you can make use of BigTable’s prefix scanning to find all entries related to "*.example.co.uk". (Our actual use case was to reverse emails so that "hello.world@example.co.uk" became "uk.co.example@hello.world", for the same reason).

### So how might we write an algorithm to do this?

Well, one method might be to split the string on dots, reverse that resulting array and then join it back on dots:

```go
func ReverseDomain(domain string) string {
	parts := strings.Split(domain, ".")
	rev(parts)
	return strings.Join(parts, ".")
}

func rev(ss []string) {
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-i-1] = ss[len(ss)-i-1], ss[i]
	}
}
```

How fast does this go? We could write a benchmark to find out:

```go
func BenchmarkReverseDomain(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReverseDomain("foo.bar.baz.co.uk")
	}
}
```

We'd get something like the following result:

```
BenchmarkReverseDomain-8       	 5000000       	       394 ns/op       	     144 B/op  	       3 allocs/op
```

which means it takes ~400 nanoseconds to reverse `foo.bar.baz.co.uk` on my machine, and does three allocations in the process.

### Can it go any faster?

We can use a common trick to avoid allocating slices of strings (aka slices of slices of bytes): reverse the entire string, then unreverse each component separately.

So you do the following steps:
```
foo.bar.baz.co.uk
ku.oc.zab.rab.oof (by reversing the whole string)
uk.oc.zab.rab.oof (reverse the ku)
uk.co.zab.rab.oof (reverse the oc)
uk.co.baz.rab.oof ( etc. )
uk.co.baz.bar.oof
uk.co.baz.bar.foo
```

Because this reversing happens in place, you don't need to allocate any extra memory for doing it:

```go
func ReverseDomainFaster(domain string) string {
	dom := []byte(domain)
	revB(dom)

	lastDotIdx := 0
	for i, b := range dom {
		if b == '.' {
			revB(dom[lastDotIdx:i])
			lastDotIdx = i + 1
		}
	}

	revB(dom[lastDotIdx:])

	return string(dom)
}

func revB(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
```

And if you benchmark that:

```go
func BenchmarkReverseDomainFaster(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReverseDomainFaster("foo.bar.baz.co.uk")
	}
}
```

You might get this:

```
BenchmarkReverseDomainFaster-8         	10000000       	       139 ns/op       	      32 B/op  	       1 allocs/op
```

Hurray! It's about three times faster.

Where does that one allocation come from? Because we're returning a copy of the string rather than the original. Eagle-eyed readers might suspect there should be two allocations: `dom := []byte(domain)` and `return string(dom)`, but it seems like Go's compiler optimises one of these away (after all, a string is just a byte slice).

We can get some evidence for this theory by running our tests with escape analysis: `go test -v ./... -bench . -gcflags '-m'` and seeing the following lines:

```
./domainreverse.go:27: string(dom) escapes to heap
./domainreverse.go:13: ReverseDomainFaster domain does not escape
./domainreverse.go:14: ReverseDomainFaster ([]byte)(domain) does not escape
```

As a final note, if we could avoid the allocation by mutating the original domain instead, our code should go much faster:

```go
func ReverseDomainInPlace(dom []byte) {
	revB(dom)

	lastDotIdx := 0
	for i, b := range dom {
		if b == '.' {
			revB(dom[lastDotIdx:i])
			lastDotIdx = i + 1
		}
	}

	revB(dom[lastDotIdx:])
}
```

Note that to make it clear that the function mutates its arguments, I choose not to return anything.

If you write one final benchmark for this in-place version:
```go
func BenchmarkReverseDomainFasterInPlace(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReverseDomainInPlace([]byte("foo.bar.baz.co.uk"))
	}
}
```

We can see that it is a lot faster than our original function and does no allocations:
```
BenchmarkReverseDomain-8                3000000   404 ns/op    144 B/op    3 allocs/op
BenchmarkReverseDomainFaster-8         10000000   133 ns/op     32 B/op    1 allocs/op
BenchmarkReverseDomainFasterInplace-8  20000000    73.4 ns/op    0 B/op    0 allocs/op
```

What will you do with your extra 330ns and 144 bytes of memory?
