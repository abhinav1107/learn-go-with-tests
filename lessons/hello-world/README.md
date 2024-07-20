# Lessons from Hello World

These were my lessons while doing [Hello, World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world) exercise:
- Still not clear with `package main` and its intended use. May be later.
- Go doesn't have Optional Arguments in function, nor does it allow method overloading.
  - This means to achieve something like calling a function with default / undeclared values, we need to use `struct`.
  - [Example](hello.go#11)
- Convert a string to title case:
  ```go
  import (
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
  )
  
  caser := cases.Title(language.English)
  tittleInput = caser.String(givenInput)
  ```
- The declaration `func Hello() string {}` means it will return a string at the end.
- Rules of creating tests in Go:
  - It needs to be in a file with a name like `xxx_test.go`
  - The test function must start with the word `Test`
  - The test function takes only one argument `t *testing.T`
- 2 ways to declare variables:
  - `varName := value`. `:=` is called `short variable declaration`, which declares and initialises the variable
  - `var varName type = expression`. This is an assignment. Here either the `type` or `= expression` can be omitted but not both.
