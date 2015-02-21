# formatifier

[![GoDoc](https://godoc.org/github.com/briandowns/formatifier?status.svg)](https://godoc.org/github.com/briandowns/formatifier) [![Build Status](https://travis-ci.org/briandowns/formatifier.svg?branch=master)](https://travis-ci.org/briandowns/formatifier)

formatifier is a Go (golang) library to quickly format strings in common formats.

This is a port of [Jim Sidler's](https://github.com/jvsidler) [formatifier](https://github.com/jvsidler/formatifier)

## Installation

```bash
go get github.com/briandowns/formatifier
```

## Examples

### Import the package

```Go
import "github.com/briandowns/formatifier"
```

### Telephone 1

```Go
result, err := ToPhone("2155551212", "-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
(215) 555-1212
```

### Telephone 2

```Go
result, err := ToPhone("12155551212", "-")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
1 (215) 555-1212
```

### URL 1

```Go
result, err := ToURL("github.com/briandowns", "", true)
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
https://github.com/briandowns
```

### URL 2

```Go
result, err := ToURL("github.com/briandowns", "gist", false)
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
http://gist.github.com/briandowns
```

### SSN

```Go
result, err := ToSSN("786540986", "-")
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
result, err := ToLoclCombo("651096", "-")
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
result, err := ToISBN("6517106483096", "-")
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
result, err := ToMorseCode("hello my friend")
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
result, err := ToPirateSpeak("hello my friend")
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
result, err := ToIRSA("hello my friend")
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
result, err := ToLeet("that is cool dude")
if err != nil {
    fmt.Println(err)
}

fmt.Println(result)
```
```bash
7#47 15 kewl d00d
```

