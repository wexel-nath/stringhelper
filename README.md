# stringhelper [![Build Status](https://travis-ci.com/wexel-nath/stringhelper.svg?branch=master)](https://travis-ci.com/wexel-nath/stringhelper)
A go package for converting string case types


# Installation
```
go get -v github.com/wexel-nath/stringhelper
```

# Usage
Converting a kebab-case string:
```go
package main

import (
	"fmt"
	"github.com/wexel-nath/stringhelper"
)

func main() {
    myString := "kev-kev"

    camelized := stringhelper.Camelize(myString)
    pascalized := stringhelper.Pascalize(myString)
    titled := stringhelper.Titleize(myString)

    fmt.Println(myString)   // kev-kev
    fmt.Println(camelized)  // kevKev
    fmt.Println(pascalized) // KevKev
    fmt.Println(titled)     // Kev Kev
}
```
