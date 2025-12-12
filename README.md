# tcellansi

[![Go Reference](https://pkg.go.dev/badge/github.com/noborus/tcellansi.svg)](https://pkg.go.dev/github.com/noborus/tcellansi)

`tcellansi` is a package for converting [tcell](https://github.com/gdamore/tcell) styles to ANSI escape sequences.

## Version Compatibility

- For **tcell v3**, use the latest `main` branch of this repository.
- For **tcell v2**, use the `v0.3.1` tag:

```sh
  go get github.com/noborus/tcellansi@v0.3.1
```

## Installation

To install the package, run:

```sh
go get github.com/noborus/tcellansi
```

## Usage

Here is a simple example of how to use `tcellansi`:

```go
package main

import (
	"github.com/gdamore/tcell/v3"
	"github.com/noborus/tcellansi"
)

func main() {
	// Initialize tcell screen
	screen, _ := tcell.NewScreen()
	screen.Init()

	// Create a style
	style := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)

	// Convert the style to ANSI escape sequence
	ansiSeq := tcellansi.ToAnsi(style)
	screen.Fini()

	println(ansiSeq + "Hello world!")
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
