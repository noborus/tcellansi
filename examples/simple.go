//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/noborus/tcellansi"
)

func main() {
	// Initialize tcell screen
	screen, _ := tcell.NewScreen()
	screen.Init()

	// Create a style
	style := tcell.StyleDefault.Background(tcell.ColorBlue).Underline(true).Underline(tcell.ColorGreen).Underline(tcell.UnderlineStyleDouble)

	setContents(screen, "Hello world!█", style)

	screen.Show()
	// Convert the style to ANSI escape sequence
	ansiSeq := tcellansi.ToAnsi(style)
	time.Sleep(2 * time.Second)
	screen.Fini()

	fmt.Println(ansiSeq + "█Hello world!█")
}

func setContents(screen tcell.Screen, str string, style tcell.Style) {
	for i, r := range str {
		screen.SetContent(i, 0, r, nil, style)
	}
}
