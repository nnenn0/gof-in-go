package decorator

import "strings"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}

// StringDisplayがDisplayインターフェースを満たしていることを確認
var _ Display = (*StringDisplay)(nil)

type StringDisplay struct {
	string string
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{string: s}
}

func (s *StringDisplay) GetColumns() int {
	return len(s.string)
}

func (s *StringDisplay) GetRows() int {
	return 1
}

func (s *StringDisplay) GetRowText(row int) string {
	if row != 0 {
		panic("Index out of bounds")
	}
	return s.string
}

type Border struct {
	display Display
}

type SideBorder struct {
	Border
	borderChar rune
}

func NewSideBorder(display Display, ch rune) *SideBorder {
	return &SideBorder{Border{display}, ch}
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.display.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.display.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return string(s.borderChar) + s.display.GetRowText(row) + string(s.borderChar)
}

type FullBorder struct {
	Border
}

func NewFullBorder(display Display) *FullBorder {
	return &FullBorder{Border{display}}
}

func (f *FullBorder) GetColumns() int {
	return 1 + f.display.GetColumns() + 1
}

func (f *FullBorder) GetRows() int {
	return 1 + f.display.GetRows() + 1
}

func (f *FullBorder) GetRowText(row int) string {
	if row == 0 || row == f.display.GetRows()+1 {
		return "+" + strings.Repeat("-", f.display.GetColumns()) + "+"
	}
	return "|" + f.display.GetRowText(row-1) + "|"
}
