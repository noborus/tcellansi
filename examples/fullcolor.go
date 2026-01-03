//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/noborus/tcellansi"
)

// Converts the screen content to ANSI strings
func toansi(s tcell.Screen) []string {
	w, h := s.Size()
	if w == 0 || h == 0 {
		return nil
	}
	return tcellansi.TrimRightSpaces(tcellansi.ScreenContentToStrings(s, 0, w, 0, h-1))
}

// Draws a color grid on the screen
func drawColorGrid(s tcell.Screen) {
	// System colors (0..15)
	for y := 0; y < 2; y++ {
		for x := 0; x < 8; x++ {
			style := tcell.StyleDefault.Background(color.IsValid + color.Color(x+y*8))
			s.SetContent(x*2, y, ' ', nil, style)
			s.SetContent(x*2+1, y, ' ', nil, style)
		}
	}

	// Color cube (16..231)
	idx := color.White + 1
	for r := 0; r < 6; r++ {
		for g := 0; g < 6; g++ {
			x := r * 7
			y := 4 + g
			for b := 0; b < 6; b++ {
				style := tcell.StyleDefault.Background(idx)
				s.SetContent(x*2, y, ' ', nil, style)
				s.SetContent(x*2+1, y, ' ', nil, style)
				idx++
				x++
			}
		}
	}

	// Grayscale ramp (232..255)
	for i := 232; i < 256; i++ {
		style := tcell.StyleDefault.Background(color.IsValid + color.Color(i))
		s.SetContent((i-232)*2, 12, ' ', nil, style)
		s.SetContent((i-232)*2+1, 12, ' ', nil, style)
	}

	y := 14
	x := 0
	// Full RGB color grid
	for g := 0; g < 16; g++ {
		for r := 0; r < 16; r++ {
			for b := 0; b < 16; b++ {
				style := tcell.StyleDefault.Background(tcell.NewRGBColor(int32(r*16), int32(g*16), int32(b*16)))
				s.SetContent(x, y, ' ', nil, style)
				x++
			}
			x++
		}
		y++
		x = 0
	}

	// Update the screen
	s.Show()
}

// Draws a string on the screen with a specific style
func drawString(s tcell.Screen, str string, style tcell.Style, y int) {
	for i, r := range str {
		s.SetContent(i, y, r, nil, style)
	}
}

// Renders text with various attributes
func renderAttribute(s tcell.Screen) {
	s.Clear()
	str := "Hello world!"
	style := tcell.StyleDefault.Underline(true)
	drawString(s, str, style, 1)
	style = tcell.StyleDefault.Bold(true)
	drawString(s, str, style, 2)
	style = tcell.StyleDefault.Italic(true)
	drawString(s, str, style, 3)
	style = tcell.StyleDefault.Reverse(true)
	drawString(s, str, style, 4)
	style = tcell.StyleDefault.Blink(true)
	drawString(s, str, style, 5)
	style = tcell.StyleDefault.Dim(true)
	drawString(s, str, style, 6)
	s.Show()
}

// Main event loop
func mainLoop() []string {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	drawColorGrid(s)
	defer s.Fini()

	// Event loop
	for {
		ev := <-s.EventQ()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyUp {
				renderAttribute(s)
			}
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyEnter {
				return toansi(s)
			}
		}
	}
}

// Main function
func main() {
	screenStr := mainLoop()
	for _, v := range screenStr {
		fmt.Printf("%s", v)
	}
}
