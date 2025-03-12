package tcellansi

import (
	"bytes"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// ToAnsi converts the tcell style to an ANSI escape sequence.
// It handles foreground color, background color, and various text attributes such as bold, italic, underline, etc.
// The function supports both palette colors and RGB colors for foreground and background.
//
// Parameters:
//   - style: tcell.Style to be converted.
//
// Returns:
//   - A string containing the ANSI escape sequence representing the given style.
func ToAnsi(style tcell.Style) string {
	var ansi bytes.Buffer
	fg, bg, attr := style.Decompose()

	// Foreground color
	if fg != tcell.ColorDefault {
		if fg&tcell.ColorIsRGB != 0 {
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
		if bg&tcell.ColorIsRGB != 0 {
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

const resetStyle = "\x1b[0m"

// ScreenContentToStrings converts the screen content to a slice of strings.
// It reads the screen content from the specified range (x1, x2, y1, y2) and converts it to a slice of strings.
// Each string in the slice represents a row of the screen content.
//
// Parameters:
//   - screen: tcell.Screen to be converted.
//   - x1: int, the starting column of the range.
//   - x2: int, the ending column of the range.
//   - y1: int, the starting row of the range.
//   - y2: int, the ending row of the range.
//
// Returns:
//   - A slice of strings representing the screen content in the specified range.
func ScreenContentToStrings(screen tcell.Screen, x1 int, x2 int, y1 int, y2 int) []string {
	var buf bytes.Buffer
	var result []string
	for row := y1; row < y2; row++ {
		prevStyle := tcell.StyleDefault
		for col := x1; col < x2; col++ {
			main, combc, style, width := screen.GetContent(col, row)
			if style != prevStyle {
				if prevStyle != tcell.StyleDefault {
					buf.WriteString(resetStyle)
				}
				prevStyle = style
			}
			styleStr := ToAnsi(style)
			buf.WriteString(styleStr)
			buf.WriteRune(main)
			for _, c := range combc {
				buf.WriteRune(c)
			}
			if width > 1 {
				col += 1
			}
		}
		buf.WriteString(resetStyle)
		buf.WriteRune('\n')
		result = append(result, buf.String())
		buf.Reset()
	}
	return result
}
