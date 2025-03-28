// command.go
package command

import (
	"errors"
)

type Command interface {
	Execute() string
}

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) Execute() string {
	output := ""
	for _, cmd := range m.commands {
		output += cmd.Execute()
	}
	return output
}

func (m *MacroCommand) Append(cmd Command) error {
	if cmd == m {
		return errors.New("infinite loop caused by append")
	}
	m.commands = append(m.commands, cmd)
	return nil
}

func (m *MacroCommand) Undo() {
	if len(m.commands) > 0 {
		m.commands = m.commands[:len(m.commands)-1]
	}
}

func (m *MacroCommand) Clear() {
	m.commands = []Command{}
}

type Drawable interface {
	Draw(x int, y int) string
}

type DrawCommand struct {
	drawable Drawable
	position Position
}

type Position struct {
	X int
	Y int
}

func NewDrawCommand(drawable Drawable, x int, y int) *DrawCommand {
	return &DrawCommand{
		drawable: drawable,
		position: Position{X: x, Y: y},
	}
}

func (d *DrawCommand) Execute() string {
	return d.drawable.Draw(d.position.X, d.position.Y)
}

type DrawCanvas struct {
	width   int
	height  int
	history *MacroCommand
	points  []Position
}

func NewDrawCanvas(width int, height int, history *MacroCommand) *DrawCanvas {
	return &DrawCanvas{
		width:   width,
		height:  height,
		history: history,
		points:  []Position{},
	}
}

func (d *DrawCanvas) Draw(x int, y int) string {
	d.points = append(d.points, Position{X: x, Y: y})
	return "描画: x=" + string(rune(x+'0')) + ", y=" + string(rune(y+'0'))
}

func (d *DrawCanvas) GetPoints() []Position {
	return d.points
}

func (d *DrawCanvas) ClearPoints() {
	d.points = []Position{}
}
