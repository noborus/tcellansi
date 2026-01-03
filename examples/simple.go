//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/noborus/tcellansi"
)

func main() {
	// Initialize tcell screen
	screen, _ := tcell.NewScreen()
	screen.Init()

	// Create a style
	style := tcell.StyleDefault.Background(color.Blue).Underline(true).Underline(color.Green).Underline(tcell.UnderlineStyleDouble)

	screen.PutStrStyled(0, 0, "Hello world!█", style)

	screen.Show()
	// Convert the style to ANSI escape sequence
	ansiSeq := tcellansi.ToAnsi(style)
	time.Sleep(2 * time.Second)
	screen.Fini()

	fmt.Println(ansiSeq + "█Hello world!█")
}
