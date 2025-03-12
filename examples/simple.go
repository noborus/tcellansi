package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/noborus/tcellansi"
)

func main() {
	// Initialize tcell screen
	screen, _ := tcell.NewScreen()
	screen.Init()

	// Create a style
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.Color250).Bold(true)

	// Convert the style to ANSI escape sequence
	ansiSeq := tcellansi.ToAnsi(style)

	screen.Fini()

	fmt.Println(ansiSeq + "█Hello, World!█")
}
