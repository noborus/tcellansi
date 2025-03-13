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
		ansi.WriteString("\x1b[38;")
		ansi.WriteString(colorToAnsi(fg, ";"))
	}

	// Background color
	if bg != tcell.ColorDefault {
		ansi.WriteString("\x1b[48;")
		ansi.WriteString(colorToAnsi(bg, ";"))
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
		/*
			ansi.WriteString("\x1b[")
			us := style.UnderlineStyle()
			ansi.WriteString(underlineStyleToAnsi(us))
			uc := style.UnderlineColor()
			if uc != tcell.ColorDefault {
				ansi.WriteString("\x1b[58:")
				ansi.WriteString(colorToAnsi(uc, ":"))
			}
		*/
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

func colorToAnsi(color tcell.Color, delm string) string {
	if color&tcell.ColorIsRGB != 0 {
		r, g, b := color.RGB()
		return fmt.Sprintf("2%s%d%s%d%s%dm", delm, r, delm, g, delm, b)
	}
	return fmt.Sprintf("5%s%dm", delm, color&^tcell.ColorValid)
}

func underlineStyleToAnsi(style tcell.UnderlineStyle) string {
	switch style {
	case tcell.UnderlineStyleSolid:
		return "4m"
	case tcell.UnderlineStyleDouble:
		return "4:2m"
	case tcell.UnderlineStyleCurly:
		return "4:3m"
	case tcell.UnderlineStyleDotted:
		return "4:4m"
	case tcell.UnderlineStyleDashed:
		return "4:5m"
	default:
		return "4m"
	}
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
