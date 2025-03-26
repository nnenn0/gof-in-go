package bridge

import (
	"fmt"
	"strings"
)

type DisplayImpl interface {
	RawOpen() string
	RawPrint() string
	RawClose() string
}

type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{impl: impl}
}

func (d *Display) Open() string {
	return d.impl.RawOpen()
}

func (d *Display) Print() string {
	return d.impl.RawPrint()
}

func (d *Display) Close() string {
	return d.impl.RawClose()
}

func (d *Display) Display() []string {
	var result []string
	result = append(result, d.Open())
	result = append(result, d.Print())
	result = append(result, d.Close())
	return result
}

type CountDisplay struct {
	impl DisplayImpl
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{impl: impl}
}

func (cd *CountDisplay) Open() string {
	return cd.impl.RawOpen()
}

func (cd *CountDisplay) Print() string {
	return cd.impl.RawPrint()
}

func (cd *CountDisplay) Close() string {
	return cd.impl.RawClose()
}
func (cd *CountDisplay) Display() []string {
	var result []string
	result = append(result, cd.Open())
	result = append(result, cd.Print())
	result = append(result, cd.Close())
	return result
}

func (cd *CountDisplay) MultiDisplay(times int) []string {
	var result []string
	result = append(result, cd.Open())
	for i := 0; i < times; i++ {
		result = append(result, cd.Print())
	}
	result = append(result, cd.Close())
	return result
}

// StringDisplayImplがDisplayImplインターフェースを満たしている
var _ DisplayImpl = (*StringDisplayImpl)(nil)

type StringDisplayImpl struct {
	str string
}

func NewStringDisplayImpl(str string) *StringDisplayImpl {
	return &StringDisplayImpl{
		str: str,
	}
}

func (s *StringDisplayImpl) RawOpen() string {
	return s.printLine()
}

func (s *StringDisplayImpl) RawPrint() string {
	return fmt.Sprintf("|%s|", s.str)
}

func (s *StringDisplayImpl) RawClose() string {
	return s.printLine()
}

func (s *StringDisplayImpl) printLine() string {
	return "+" + strings.Repeat("-", len(s.str)) + "+"
}
