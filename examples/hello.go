//go:build ignore
// +build ignore

package main

import (
	"github.com/gdamore/tcell/v2"
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
