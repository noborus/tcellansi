package tcellansi

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestStyleToES(t *testing.T) {
	tests := []struct {
		name  string
		style tcell.Style
		want  string
	}{
		{
			name:  "default style",
			style: tcell.StyleDefault,
			want:  "",
		},
		{
			name:  "foreground color",
			style: tcell.StyleDefault.Foreground(tcell.ColorRed),
			want:  "\x1b[38;5;9m", // palette color processing
		},
		{
			name:  "foreground color2",
			style: tcell.StyleDefault.Foreground(tcell.GetColor("#ff0000")),
			want:  "\x1b[38;2;255;0;0m", // RGB color processing
		},
		{
			name:  "background color",
			style: tcell.StyleDefault.Background(tcell.ColorBlue),
			want:  "\x1b[48;5;12m", // palette color processing
		},
		{
			name:  "background color2",
			style: tcell.StyleDefault.Background(tcell.GetColor("#0000ff")),
			want:  "\x1b[48;2;0;0;255m", // RGB color processing
		},
		{
			name:  "italic attribute",
			style: tcell.StyleDefault.Italic(true),
			want:  "\x1b[3m",
		},
		{
			name:  "underline attribute",
			style: tcell.StyleDefault.Underline(true),
			want:  "\x1b[4m",
		},
		{
			name:  "combined attributes",
			style: tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorYellow).Bold(true).Underline(true),
			want:  "\x1b[38;5;2m\x1b[48;5;11m\x1b[1m\x1b[4m", // palette color, palette color, bold, underline
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StyleToANSI(tt.style); got != tt.want {
				t.Errorf("StyleToES() = %v, want %v", got, tt.want)
			}
		})
	}
}
