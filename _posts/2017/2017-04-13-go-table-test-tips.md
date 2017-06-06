---
layout: "post"
title: "Levelling up Go table-driven tests"
date: "2017-04-13 20:18"
---

Writing table-driven tests allows me to focus on the inputs and expected outputs of my code, rather than the boilerplate needed to make the test work.



Here are some opinionated tips on how to make yours even better, both for you and others on your team:



### Use a named testcase struct

```go
// bad
func TestFooer(t *testing.T) {
  tests := []struct{
    inputOne, inputTwo string
    expected int
  }{}
}

// good
func TestFooer(t *testing.T) {
  type testcase struct {
    inputOne, inputTwo string
    expected int
  }

  tests := []testcase{}
}
```

Naming the testcase struct allows you to write the struct name inside the slice, which aids readability and copy-pastability. Name the struct `testcase` for consistency.

### Name each of your testcases

```go
// bad - how do you know what failed?
tests := []struct{
  inputOne, inputTwo string
  expected int
}{}

// good
type testcase struct {
  name string // <-- name summarises testcase
  inputOne, inputTwo string
  expected int
}

tests := []testcase{
  testcase{name: "empty_strings"}
}

```

Good test names describe the behaviour you are attempting to capture. They also make the failing test case easier to find.

Do not use spaces in your test name: `t.Run` will replace them with underscores.

### Run each testcase inside t.Run


```go
// bad
for _, test := range tests {
  if test.expected != Foo(inputOne, inputTwo) {
    t.Errorf("%s failed", test.name)
  }
}

// good
for _, test := range tests {
  // run as a subtest
  t.Run(test.name, func(t *testing.T) {
    if test.expected != Foo(inputOne, inputTwo) {
      t.Errorf("%s failed", test.name)
    }
  })
}
```

`t.Run` runs subtests inside your test. It makes it very clear which subtest is failing, and in the case of panics you'll see which subtest caused the panic.

Don't use spaces in your `test.name` as `t.Run` replaces them with underscores. This means you won't be able to copy the failing test name directly from your terminal to your editor to find it.

### Leave struct names in slices

```go
// bad
func TestFooer(t *testing.T) {
  type testcase struct {
    inputOne, inputTwo string
    expected int
  }

  tests := []testcase{
    {
      inputOne: "foo",
      inputTwo: "bar",
      expected: 3,
    },
  }
}

// good
func TestFooer(t *testing.T) {
  type testcase struct {
    inputOne, inputTwo string
    expected int
  }

  tests := []testcase{
    // Can clearly see where each testcase begins.
    testcase{
      inputOne: "foo",
      inputTwo: "bar",
      expected: 3,
    },
  }
}
```

Go allows you to omit the struct name in slices of structs. Avoid this temptation in table-driven tests. Always write out the struct name. It is easy to get lost in nested slices of slices of structs.

VSCode also has a bug where, when writing a slice of structs, IntelliSense tries to suggest struct fields even though you are not inside the struct. Writing the struct name works around this bug.

![VSCode doesn't know to suggest "HelloWorld"]({{site.baseurl}}/images/vscode-bug-2.png "VSCode doesn't know to suggest 'HelloWorld'")
<p style="text-align: center; font-style: italic;">VSCode fails to suggest "HelloWorld".</p>

![Specifying the struct works around the bug]({{site.baseurl}}/images/vscode-bug-3.png "Specifying the struct works around the bug")
<p style="text-align: center; font-style: italic;">Specifying the struct works around the bug.</p>



### Use named struct fields

```go
// bad
tests := []testcase{
  testcase{"foo", "bar", 3},
}

// good
tests := []testcase{
  testcase{
    inputOne: "foo",
    inputTwo: "bar",
    expected: 3,
  },
}
```

Named struct fields mean it is easier to copy and paste existing testcases and know what to fill in. It also means that you can reorder fields without breaking all existing code, and that you don't need to specify all fields in the struct.

### Bind CMD+T to "Run this test"

I use VSCode with the Go plugin. Binding CMD+T to "Test function at cursor" has improved my productivity significantly. I like quick feedback loops, so I can write another testcase and check it without switching windows etc.

I've also bound CMD+R to "Rerun previous test", which is also useful.

### Use custom assert libraries, but not test suites.

Custom assert libraries can be very useful, as you'll otherwise be writing the same deep-equals cruft everywhere.

Leave out custom test suite builders (with functions like like `Before`, `BeforeEach` etc.) They are not particularly useful and are incompatible with the key bindings above.

### Avoid test panics

If a test panics, all tests stop execution. This is immensely frustrating if you know multiple tests are broken, but you can only fix one at a time because only one panics at a time.

So test for nil pointers and array out of bounds errors as you usually world. Fail and return from tests early if things are broken, pointers are nil, slices are empty, or you cannot make any more meaningful assertions.
