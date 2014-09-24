# formatifier

formatifier is a Go (golang) library to quickly format strings in common formats.

This is a port of [Jim Sidler's](https://github.com/jvsidler) [formatifier](https://github.com/jvsidler/formatifier) 

## Documentation

```bash
godoc -http=:6060
```

## Warning

This application is a Work In Progress.  All exported functions __except__ ToPhone() are currently working.

## Installation

```bash
go get github.com/bdowns328/formatifier
```

## Examples

### Import the package, create a new formatifier object.

```Go
package main

import (
    "github.com/bdowns328/formatifier"
)

func main() {
    str := formatifier.New("hello world")
}
```

### Telephone

```Go
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### URL

```Go
str := formatifier.New("github.com/bdowns328")

result, err := str.ToURL(true, "www")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### SSN

```Go
str := formatifier.New("786540986")

result, err := str.ToSSN("-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### Lock Combo

```Go
str := formatifier.New("651096")

result, err := str.ToLoclCombo("-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### ISBN

```Go
str := formatifier.New("651096")

result, err := str.ToISBN("-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### Morse Code

```Go
str := formatifier.New("hey this is going to be in morse code")

result, err := str.ToMorseCode()
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### Pirate Speak

```Go
str := formatifier.New("hello my friend")

result, err := str.ToPirateSpeak()
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### IRSA

```Go
str := formatifier.New("hello my friend")

result, err := str.ToIRSA()
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

### Leet Speak

```Go
str := formatifier.New("that is cool dude")

result, err := str.ToLeet()
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```

