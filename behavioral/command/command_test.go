package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "gof-in-go/behavioral/command"
)

type MockCommand struct {
	output string
}

func (m *MockCommand) Execute() string {
	return m.output
}

type MockDrawable struct {
	drawnPoints []Position
}

func (m *MockDrawable) Draw(x int, y int) string {
	m.drawnPoints = append(m.drawnPoints, Position{X: x, Y: y})
	return ""
}

func (m *MockDrawable) GetDrawnPoints() []Position {
	return m.drawnPoints
}

func TestMacroCommand(t *testing.T) {
	macroCmd := &MacroCommand{}
	assert.NotNil(t, macroCmd, "MacroCommand が作成できること")
	assert.Empty(t, macroCmd.Execute(), "初期状態では実行しても何も起きないこと")

	mockCmd1 := &MockCommand{output: "コマンド1実行"}
	mockCmd2 := &MockCommand{output: "コマンド2実行"}

	err := macroCmd.Append(mockCmd1)
	assert.NoError(t, err, "コマンドの追加が成功すること")
	err = macroCmd.Append(mockCmd2)
	assert.NoError(t, err, "コマンドの追加が成功すること")

	output := macroCmd.Execute()
	assert.Equal(t, "コマンド1実行コマンド2実行", output, "追加したコマンドが順番に実行されること")

	macroCmd.Undo()
	output = macroCmd.Execute()
	assert.Equal(t, "コマンド1実行", output, "Undo で最後のコマンドが取り消されること")

	macroCmd.Clear()
	assert.Empty(t, macroCmd.Execute(), "Clear で全てのコマンドが削除されること")

	err = macroCmd.Append(macroCmd)
	assert.Error(t, err, "自身を追加しようとするとエラーになること")
}

func TestDrawCommand(t *testing.T) {
	mockDrawable := &MockDrawable{}

	drawCmd := NewDrawCommand(mockDrawable, 10, 20)
	assert.NotNil(t, drawCmd, "DrawCommand が作成できること")

	drawCmd.Execute()
	drawnPoints := mockDrawable.GetDrawnPoints()
	assert.Len(t, drawnPoints, 1, "Drawable の Draw メソッドが呼び出されること")
	assert.Equal(t, 10, drawnPoints[0].X, "正しい X 座標で Draw メソッドが呼び出されること")
	assert.Equal(t, 20, drawnPoints[0].Y, "正しい Y 座標で Draw メソッドが呼び出されること")
}

func TestMainLogic(t *testing.T) {
	history := &MacroCommand{}
	canvas := NewDrawCanvas(400, 400, history)

	cmd1 := NewDrawCommand(canvas, 10, 20)
	history.Append(cmd1)
	cmd1.Execute()
	assert.Len(t, canvas.GetPoints(), 1, "1回目のドラッグで点が描画されること")
	assert.Equal(t, Position{X: 10, Y: 20}, canvas.GetPoints()[0], "1回目のドラッグで正しい座標に点が描画されること")

	cmd2 := NewDrawCommand(canvas, 30, 40)
	history.Append(cmd2)
	cmd2.Execute()
	assert.Len(t, canvas.GetPoints(), 2, "2回目のドラッグで点が描画されること")
	assert.Equal(t, Position{X: 30, Y: 40}, canvas.GetPoints()[1], "2回目のドラッグで正しい座標に点が描画されること")

	history.Clear()
	canvas.ClearPoints()
	assert.Empty(t, history.Execute(), "履歴がクリアされること")
	assert.Empty(t, canvas.GetPoints(), "描画領域もクリアされること")
}
