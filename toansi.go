package tcellansi

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gdamore/tcell/v3"
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
	fg := style.GetForeground()
	bg := style.GetBackground()

	// Foreground color
	if fg != tcell.ColorDefault {
		ansi.WriteString("\x1b[")
		ansi.WriteString(foregroundColorToAnsi(fg))
		ansi.WriteString("m")
	}
	// Background color
	if bg != tcell.ColorDefault {
		ansi.WriteString("\x1b[")
		ansi.WriteString(backgroundColorToAnsi(bg))
		ansi.WriteString("m")

	}
	if style.HasBold() {
		ansi.WriteString("\x1b[1m")
	}
	if style.HasDim() {
		ansi.WriteString("\x1b[2m")
	}
	if style.HasItalic() {
		ansi.WriteString("\x1b[3m")
	}
	if style.HasUnderline() {
		ansi.WriteString(underlineToAnsi(style))
	}
	if style.HasBlink() {
		ansi.WriteString("\x1b[5m")
	}
	if style.HasReverse() {
		ansi.WriteString("\x1b[7m")
	}
	if style.HasStrikeThrough() {
		ansi.WriteString("\x1b[9m")
	}
	return ansi.String()
}

// foregroundColorToAnsi converts the foreground color to an ANSI escape sequence.
func foregroundColorToAnsi(fg tcell.Color) string {
	if fg > tcell.ColorWhite {
		return "38;" + colorToAnsi(fg, ";")
	}
	if (fg - tcell.ColorValid) < 8 {
		return strconv.Itoa(int(30 + (fg - tcell.ColorValid)))
	}
	return strconv.Itoa(int(82 + (fg - tcell.ColorValid)))
}

// backgroundColorToAnsi converts the background color to an ANSI escape sequence.
func backgroundColorToAnsi(bg tcell.Color) string {
	if bg > tcell.ColorWhite {
		return "48;" + colorToAnsi(bg, ";")
	}
	if (bg - tcell.ColorValid) < 8 {
		return strconv.Itoa(int(40 + (bg - tcell.ColorValid)))
	}
	return strconv.Itoa(int(92 + (bg - tcell.ColorValid)))
}

// colorToAnsi converts the tcell.Color to an ANSI escape sequence.
func colorToAnsi(color tcell.Color, delm string) string {
	if color&tcell.ColorIsRGB == 0 {
		return fmt.Sprintf("5%s%d", delm, color&^tcell.ColorValid)
	}
	r, g, b := color.RGB()
	return fmt.Sprintf("2%s%d%s%d%s%d", delm, r, delm, g, delm, b)
}

// underlineToAnsi converts the underline style and color to an ANSI escape sequence.
func underlineToAnsi(style tcell.Style) string {
	var ansi bytes.Buffer
	ansi.WriteString("\x1b[")
	us := getUnderlineStyle(style)
	ansi.WriteString(underlineStyleToAnsi(us))
	uc := getUnderlineColor(style)
	if uc != tcell.ColorDefault {
		ansi.WriteString("\x1b[58:")
		ansi.WriteString(colorToAnsi(uc, ":"))
		ansi.WriteString("m")
	}
	return ansi.String()
}

// getUnderlineStyle returns the underline style of the given style.
// This is a temporary function to retrieve the underline style using reflection.
// It assumes that the tcell.Style type has a method named "UnderlineStyle".
// Once tcell officially supports UnderlineStyle, this function should be replaced.
func getUnderlineStyle(style tcell.Style) tcell.UnderlineStyle {
	v := reflect.ValueOf(style)
	m := v.MethodByName("GetUnderlineStyle")
	if m.IsValid() {
		results := m.Call(nil)
		if len(results) == 1 {
			if us, ok := results[0].Interface().(tcell.UnderlineStyle); ok {
				return us
			}
		}
	}
	return tcell.UnderlineStyleSolid
}

// getUnderlineColor returns the underline color of the given style.
// This is a temporary function to retrieve the underline color using reflection.
// It assumes that the tcell.Style type has a method named "UnderlineColor".
// Once tcell officially supports UnderlineColor, this function should be replaced.
func getUnderlineColor(style tcell.Style) tcell.Color {
	v := reflect.ValueOf(style)
	m := v.MethodByName("GetUnderlineColor")
	if m.IsValid() {
		results := m.Call(nil)
		if len(results) == 1 {
			if uc, ok := results[0].Interface().(tcell.Color); ok {
				return uc
			}
		}
	}
	return tcell.ColorDefault
}

// underlineStyleToAnsi converts the tcell.UnderlineStyle to an ANSI escape sequence.
// It converts the underline style to the corresponding ANSI escape sequence.
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
			str, style, width := screen.Get(col, row)
			if width > 1 {
				col++
				if col >= x2 {
					break
				}
			}
			if style != prevStyle {
				if prevStyle != tcell.StyleDefault {
					buf.WriteString(resetStyle)
				}
				prevStyle = style
				styleStr := ToAnsi(style)
				buf.WriteString(styleStr)
			}
			buf.WriteString(str)
		}
		if prevStyle != tcell.StyleDefault {
			buf.WriteString(resetStyle)
		}
		buf.WriteRune('\n')
		result = append(result, buf.String())
		buf.Reset()
	}
	return result
}

// TrimRightSpaces trims trailing spaces from each line of the given screen content strings.
// ANSI escape sequences are preserved.
// TrimRightSpaces removes trailing spaces from each line in the given slice of strings.
// If a line ends with a newline character, it trims spaces before the newline.
// If the line ends with a specific reset style sequence before the newline, the original lines are returned unmodified.
// The function preserves the newline character at the end of each line.
func TrimRightSpaces(lines []string) []string {
	trimmed := make([]string, len(lines))
	for i, line := range lines {
		n := len(line)
		end := n
		if n > 0 && line[n-1] == '\n' {
			end--
		}
		if end >= len(resetStyle) && line[end-len(resetStyle):end] == resetStyle {
			trimmed[i] = line
			continue
		}

		trimmedLine := line[:end]
		for len(trimmedLine) > 0 && trimmedLine[len(trimmedLine)-1] == ' ' {
			trimmedLine = trimmedLine[:len(trimmedLine)-1]
		}
		trimmed[i] = trimmedLine + "\n"
	}
	return trimmed
}
