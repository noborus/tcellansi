//go:build ignore
// +build ignore

package main

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/noborus/tcellansi"
)

func main() {
	// Initialize tcell screen
	screen, _ := tcell.NewScreen()
	screen.Init()

	// Create a style
	style := tcell.StyleDefault.Foreground(color.Red).Background(color.Black)

	// Convert the style to ANSI escape sequence
	ansiSeq := tcellansi.ToAnsi(style)
	screen.Fini()

	println(ansiSeq + "Hello world!")
}
