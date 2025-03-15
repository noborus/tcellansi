package tcellansi

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestToAnsi(t *testing.T) {
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
			want:  "\x1b[91m", // palette color processing
		},
		{
			name:  "foreground color256",
			style: tcell.StyleDefault.Foreground(tcell.Color250),
			want:  "\x1b[38;5;250m", // updated palette color processing
		},
		{
			name:  "foreground colorRGB",
			style: tcell.StyleDefault.Foreground(tcell.GetColor("#ff0000")),
			want:  "\x1b[38;2;255;0;0m", // RGB color processing
		},
		{
			name:  "background color",
			style: tcell.StyleDefault.Background(tcell.ColorBlue),
			want:  "\x1b[104m", // palette color processing
		},
		{
			name:  "background colorRGB",
			style: tcell.StyleDefault.Background(tcell.GetColor("#0000ff")),
			want:  "\x1b[48;2;0;0;255m", // RGB color processing
		},
		{
			name:  "italic attribute",
			style: tcell.StyleDefault.Italic(true),
			want:  "\x1b[3m",
		},
		{
			name:  "bold attribute",
			style: tcell.StyleDefault.Bold(true),
			want:  "\x1b[1m",
		},
		{
			name:  "dim attribute",
			style: tcell.StyleDefault.Dim(true),
			want:  "\x1b[2m",
		},
		{
			name:  "underline attribute",
			style: tcell.StyleDefault.Underline(true),
			want:  "\x1b[4m",
		},
		{
			name:  "blink attribute",
			style: tcell.StyleDefault.Blink(true),
			want:  "\x1b[5m",
		},
		{
			name:  "reverse attribute",
			style: tcell.StyleDefault.Reverse(true),
			want:  "\x1b[7m",
		},
		{
			name:  "strike through attribute",
			style: tcell.StyleDefault.StrikeThrough(true),
			want:  "\x1b[9m",
		},
		{
			name:  "combined attributes",
			style: tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorYellow).Bold(true).Underline(true),
			want:  "\x1b[32m\x1b[103m\x1b[1m\x1b[4m", // palette color, palette color, bold, underline
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToAnsi(tt.style); got != tt.want {
				t.Errorf("StyleToES() = %v, want %v", got, tt.want)
			}
		})
	}
}

// SetLineContent sets a line of content on the screen at the specified row with the given style.
func SetLineContent(screen tcell.Screen, row int, content string, style tcell.Style) {
	for col, r := range content {
		screen.SetContent(col, row, r, nil, style)
	}
}
func TestScreenContentToStrings(t *testing.T) {
	tests := []struct {
		name   string
		screen tcell.Screen
		x1, x2 int
		y1, y2 int
		want   []string
	}{
		{
			name: "empty screen",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				return s
			}(),
			x1: 0, x2: 10, y1: 0, y2: 10,
			want: []string{
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n",
				"          \x1b[0m\n"},
		},
		{
			name: "single cell",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				s.SetContent(0, 0, 'A', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
				return s
			}(),
			x1: 0, x2: 1, y1: 0, y2: 1,
			want: []string{"\x1b[91mA\x1b[0m\n"},
		},
		{
			name: "multiple cells",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				s.SetContent(0, 0, 'A', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
				s.SetContent(1, 0, 'B', nil, tcell.StyleDefault.Background(tcell.ColorBlue))
				s.SetContent(0, 1, 'C', nil, tcell.StyleDefault.Bold(true))
				return s
			}(),
			x1: 0, x2: 2, y1: 0, y2: 2,
			want: []string{
				"\x1b[91mA\x1b[0m\x1b[104mB\x1b[0m\n",
				"\x1b[1mC\x1b[0m \x1b[0m\n",
			},
		},
		{
			name: "wide character",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				s.SetContent(0, 0, '亜', nil, tcell.StyleDefault)
				return s
			}(),
			x1: 0, x2: 2, y1: 0, y2: 1,
			want: []string{"亜\x1b[0m\n"},
		},
		{
			name: "combc characters",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				s.SetContent(0, 0, 'A', []rune{'\u0301'}, tcell.StyleDefault.Foreground(tcell.ColorRed))
				return s
			}(),
			x1: 0, x2: 2, y1: 0, y2: 1,
			want: []string{"\x1b[91mÁ\x1b[0m \x1b[0m\n"},
		},
		{
			name: "underlineStyle",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				s.SetContent(0, 0, 'A', nil, tcell.StyleDefault.Underline(true).Underline(tcell.UnderlineStyleCurly))
				return s
			}(),
			x1: 0, x2: 1, y1: 0, y2: 1,
			want: []string{"\x1b[4mA\x1b[0m\n"},
		},
		{
			name: "single line",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				SetLineContent(s, 0, "Hello, World!", tcell.StyleDefault.Foreground(tcell.ColorRed))
				return s
			}(),
			x1: 0, x2: 13, y1: 0, y2: 1,
			want: []string{"\x1b[91mHello, World!\x1b[0m\n"},
		},
		{
			name: "over line",
			screen: func() tcell.Screen {
				s := tcell.NewSimulationScreen("")
				s.Init()
				SetLineContent(s, 0, "1234567890123456789012345678901234567890123456789012345678901234567890123456789あ", tcell.StyleDefault.Foreground(tcell.ColorRed))
				return s
			}(),
			x1: 0, x2: 80, y1: 0, y2: 1,
			want: []string{"\x1b[91m1234567890123456789012345678901234567890123456789012345678901234567890123456789\x1b[0m\n"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScreenContentToStrings(tt.screen, tt.x1, tt.x2, tt.y1, tt.y2)
			if len(got) != len(tt.want) {
				t.Errorf("ScreenContentToStrings() = \n%v, want \n%v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ScreenContentToStrings() = \n%v, want \n%v", got, tt.want)
				}
			}
		})
	}
}
