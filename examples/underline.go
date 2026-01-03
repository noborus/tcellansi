//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/noborus/tcellansi"
)

func main() {
	var lines []string
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer func() {
		screen.Fini()
		for _, line := range lines {
			fmt.Print(line)
		}
	}()
	// title
	titleStyle := tcell.StyleDefault.Foreground(color.White).Bold(true)
	screen.PutStrStyled(0, 0, "UnderlineColor Examples (Press ESC to exit)", titleStyle)

	// Underline examples
	examples := []struct {
		y     int
		text  string
		style tcell.Style
		desc  string
	}{
		{
			y:    2,
			text: "Solid Red Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleSolid, color.Red),
			desc: "Style: Solid, Color: Red",
		},
		{
			y:    4,
			text: "Double Blue Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleDouble, color.Blue),
			desc: "Style: Double, Color: Blue",
		},
		{
			y:    6,
			text: "Curly Green Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleCurly, color.Green),
			desc: "Style: Curly, Color: Green",
		},
		{
			y:    8,
			text: "Dotted Yellow Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleDotted, color.Yellow),
			desc: "Style: Dotted, Color: Yellow",
		},
		{
			y:    10,
			text: "Dashed Magenta Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleDashed, tcell.NewRGBColor(139, 0, 139)),
			desc: "Style: Dashed, Color: Magenta",
		},
		{
			y:    12,
			text: "Custom RGB Orange Curly Underline",
			style: tcell.StyleDefault.Foreground(color.White).
				Underline(tcell.UnderlineStyleCurly, tcell.NewRGBColor(255, 165, 0)),
			desc: "",
		},
	}

	descStyle := tcell.StyleDefault.Foreground(color.Gray)
	for _, ex := range examples {
		screen.PutStrStyled(0, ex.y, ex.text, ex.style)
		if ex.desc != "" {
			screen.PutStrStyled(30, ex.y, padRight(ex.text, 30)+ex.desc, descStyle)
		}
	}

	screen.Show()
	w, _ := screen.Size()
	lines = tcellansi.ScreenContentToStrings(screen, 0, w, 0, 13)
	// Wait for ESC key to exit
	for {
		ev := <-screen.EventQ()
		switch tev := ev.(type) {
		case *tcell.EventKey:
			if tev.Key() == tcell.KeyEscape {
				return
			}
		}
	}
}

// padRight is a helper function to pad a string with spaces to the right up to the specified length.
func padRight(str string, length int) string {
	for len([]rune(str)) < length {
		str += " "
	}
	return str
}
