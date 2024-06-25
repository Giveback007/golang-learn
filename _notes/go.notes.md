## WHICH TYPE SHOULD I USE?
Unless you have a good reason to, stick to the following types:
- bool
- string
- int
- uint (??)
- byte (??)
- rune (??)
- float64
- complex128 (??)

---

## CONSTANTS
```go
const pi = 3.14159
```

Can be the fallowing:
- char
- string
- bool
- numeric (any of)

Can't be complex types:
- slices
- maps
- structs

#### Can't do this:
```go
// the current time can only be known when the program is running
const currentTime = time.Now()
```

Constants can only be computed at runtime

---

## CONDITIONALS
```go
if height > 4 {
    fmt.Println("You are tall enough!")
}
```

`if` statements don't use `()`

### This:
```go
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
```
### Can be turned to this:
```go
if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
```

---

## STRUCTS

