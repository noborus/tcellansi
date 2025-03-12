package tcellansi

import (
	"bytes"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// StyleToANSI converts the tcell style to ANSI escape sequence.
// StyleToANSI converts a tcell.Style to an ANSI escape sequence string.
// It handles foreground color, background color, and various text attributes such as bold, italic, underline, etc.
// The function supports both palette colors and RGB colors for foreground and background.
//
// Parameters:
//   - style: tcell.Style to be converted.
//
// Returns:
//   - A string containing the ANSI escape sequence representing the given style.
func StyleToANSI(style tcell.Style) string {
	var ansi bytes.Buffer
	fg, bg, attr := style.Decompose()

	// Foreground color
	if fg != tcell.ColorDefault {
		if fg&^tcell.ColorValid == 0 {
			// Palette color
			ansi.WriteString(fmt.Sprintf("\x1b[38;5;%dm", int(fg)))
		} else if fg&tcell.ColorIsRGB != 0 {
			// RGB color
			r, g, b := fg.RGB()
			ansi.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b))
		} else {
			// Other cases (treated as palette color)
			ansi.WriteString(fmt.Sprintf("\x1b[38;5;%dm", fg&^tcell.ColorValid))
		}
	}

	// Background color
	if bg != tcell.ColorDefault {
		if bg&^tcell.ColorValid == 0 {
			// Palette color
			ansi.WriteString(fmt.Sprintf("\x1b[48;5;%dm", int(bg)))
		} else if bg&tcell.ColorIsRGB != 0 {
			// RGB color
			r, g, b := bg.RGB()
			ansi.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b))
		} else {
			// Other cases (treated as palette color)
			ansi.WriteString(fmt.Sprintf("\x1b[48;5;%dm", bg&^tcell.ColorValid))
		}
	}
	// Style attributes (Bold, Italic, Underline, etc.)
	if attr&tcell.AttrBold != 0 {
		ansi.WriteString("\x1b[1m")
	}
	if attr&tcell.AttrDim != 0 {
		ansi.WriteString("\x1b[2m")
	}
	if attr&tcell.AttrItalic != 0 {
		ansi.WriteString("\x1b[3m")
	}
	if attr&tcell.AttrUnderline != 0 {
		ansi.WriteString("\x1b[4m")
	}

	if attr&tcell.AttrBlink != 0 {
		ansi.WriteString("\x1b[5m")
	}
	if attr&tcell.AttrReverse != 0 {
		ansi.WriteString("\x1b[7m")
	}
	if attr&tcell.AttrStrikeThrough != 0 {
		ansi.WriteString("\x1b[9m")
	}
	return ansi.String()
}
