package template_method

import (
	"strings"
)

type Formatter interface {
	Open() string
	Content() string
	Close() string
}

type Template struct {
	formatter Formatter
}

func NewTemplate(formatter Formatter) *Template {
	return &Template{formatter: formatter}
}

func (t *Template) Format() string {
	var sb strings.Builder
	sb.WriteString(t.formatter.Open())
	for range 5 {
		sb.WriteString(t.formatter.Content())
	}
	sb.WriteString(t.formatter.Close())
	return sb.String()
}

// CharがFormatterインターフェースを満たすことを確認
var _ Formatter = (*Char)(nil)

type Char struct {
	ch rune
}

func NewChar(ch rune) *Char {
	return &Char{ch: ch}
}

func (c *Char) Open() string {
	return "<<"
}

func (c *Char) Content() string {
	return string(c.ch)
}

func (c *Char) Close() string {
	return ">>\n"
}

// StrがFormatterインターフェースを満たすことを確認
var _ Formatter = (*Str)(nil)

type Str struct {
	text string
}

func NewStr(text string) *Str {
	return &Str{text: text}
}

func (s *Str) Open() string {
	return s.line()
}

func (s *Str) Content() string {
	return "|" + s.text + "|\n"
}

func (s *Str) Close() string {
	return s.line()
}

func (s *Str) line() string {
	return "+" + strings.Repeat("-", len(s.text)) + "+\n"
}
