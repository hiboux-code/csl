# hippwn/csl

A basic package to identify the type of console attached to the program.

## Installation

```bash
go get github.com/hippwn/csl
```

## Usage

CLS currently provides only one API method. `GetConsole()` returns an integer you'll have to compare with the package constants.

```go
package main

import (
	"fmt"

	"github.com/hippwn/csl"
)

func main() {
	switch csl.GetConsole() {
	case csl.CSL_NULL:
		// no console available :(

	case csl.CSL_TERM:
		fmt.Println("Hi terminal!")

	case csl.CSL_REDIRECT:
		fmt.Println("That's redirected.")

	case csl.CSL_ERR:
		fmt.Println("Oops, idk what this is :o")
	}

}
```
