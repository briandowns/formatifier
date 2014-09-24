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

### SSN

```Go
str := formatifier.New("786540986")
fmt.Println(str.ToSSN("-"))
```