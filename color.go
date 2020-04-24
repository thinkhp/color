package color

import "strconv"

const escape = "\x1b"

// Base attributes
const (
	Reset int = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack int = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack int = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack int = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack int = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

type Attr []string

func New(attr ...int) Attr {
	a := make([]string, len(attr))
	for i, v := range attr {
		a[i] = strconv.Itoa(v)
	}
	return a
}

func (a Attr)Sprint(str string) string {
	s := "\033["
	for _, v := range a {
		s += v+";"
	}
	s = s[:len(s)-1]+"m"
	return s+str+"\033[0m"
}
