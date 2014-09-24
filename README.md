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
// Not implemented yet
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash

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
```bash
https://www.github.com/bdowns328
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
```bash
786-54-0986
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
```bash
65-10-96
```

### ISBN

```Go
str := formatifier.New("6517106483096")

result, err := str.ToISBN("-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
651-7-10-648309-6
```
### Morse Code

```Go
str := formatifier.New("hello my friend")

result, err := str.ToMorseCode()
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
. . . ... _ . .. _ . ._ _ _ _ __ . _ _ . . _ .. _ .. .._ ._ . .
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
```bash
ahoy my mate
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
```bash
hotel echo lima lima oscar  |  mike yankee  |  foxtrot romeo india echo november delta 
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
```bash
7#47 15 kewl d00d
```

