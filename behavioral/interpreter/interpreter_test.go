package interpreter_test

import (
	"testing"

	. "gof-in-go/behavioral/interpreter"

	"github.com/stretchr/testify/assert"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr string
	}{
		{
			name:     "valid program",
			input:    "program go end",
			expected: "[program [go]]",
		},
		{
			name:     "valid program with repeat",
			input:    "program repeat 3 go right end end",
			expected: "[program [[repeat 3 [go, right]]]]",
		},
		{
			name:        "missing program keyword",
			input:       "go end",
			expectedErr: "Error: 'program' が期待されましたが、'go' が見つかりました。",
		},
		{
			name:        "missing end keyword",
			input:       "program go",
			expectedErr: "Error: 'end' が見つかりません。",
		},
		{
			name:        "invalid primitive command",
			input:       "program walk end",
			expectedErr: "Error: 未知の <primitive command>: 'walk'",
		},
		{
			name:        "repeat without number",
			input:       "program repeat go end end",
			expectedErr: "Error: go",
		},
		{
			name:        "repeat with non-number",
			input:       "program repeat abc go end end",
			expectedErr: "Error: strconv.Atoi: parsing \"abc\": invalid syntax",
		},
		{
			name:        "extra tokens at the end",
			input:       "program go end extra",
			expectedErr: "Error: 予期しないトークン: 'extra'",
		},
		{
			name:     "nested repeat",
			input:    "program repeat 2 repeat 3 left go end end end",
			expected: "[program [[repeat 2 [[repeat 3 [left, go]]]]]]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			node, err := ParseProgram(test.input)
			if test.expectedErr != "" {
				assert.Error(t, err, test.expectedErr, "エラーメッセージが期待通りではありません: "+test.name)
			} else {
				assert.NoError(t, err, "エラーが発生しました: "+test.name)
				assert.Equal(t, test.expected, node.String(), "解析結果が期待通りではありません: "+test.name)
			}
		})
	}
}

func TestContext(t *testing.T) {
	text := "program repeat 3 go end"
	ctx := NewContext(text)

	assert.Equal(t, "program", ctx.CurrentToken(), "最初のトークンが正しくありません")
	assert.Equal(t, "repeat", ctx.NextToken(), "次のトークンが正しくありません")
	assert.Equal(t, "repeat", ctx.CurrentToken(), "現在のトークンが正しくありません")

	err := ctx.SkipToken("repeat")
	assert.NoError(t, err, "SkipTokenでエラーが発生しました")
	assert.Equal(t, "3", ctx.CurrentToken(), "スキップ後のトークンが正しくありません")

	num, err := ctx.CurrentNumber()
	assert.NoError(t, err, "CurrentNumberでエラーが発生しました")
	assert.Equal(t, 3, num, "数値の変換が正しくありません")
	ctx.NextToken()

	err = ctx.SkipToken("go")
	assert.NoError(t, err, "SkipTokenでエラーが発生しました")
	assert.Equal(t, "end", ctx.CurrentToken(), "スキップ後のトークンが正しくありません")

	err = ctx.SkipToken("end")
	assert.NoError(t, err, "SkipTokenでエラーが発生しました")
	assert.Equal(t, "", ctx.CurrentToken(), "最後のトークン後の状態が正しくありません")

	err = ctx.SkipToken("extra")
	assert.Error(t, err, "存在しないトークンをスキップしてもエラーが発生しませんでした")
	assert.Contains(t, err.Error(), "'extra' が期待されましたが、トークンが見つかりません。", "エラーメッセージが期待通りではありません")
}

func TestPrimitiveCommandNodeParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr string
	}{
		{
			name:     "valid go",
			input:    "go",
			expected: "go",
		},
		{
			name:     "valid right",
			input:    "right",
			expected: "right",
		},
		{
			name:     "valid left",
			input:    "left",
			expected: "left",
		},
		{
			name:        "invalid command",
			input:       "walk",
			expectedErr: "Error: 未知の <primitive command>: 'walk'",
		},
		{
			name:        "empty input",
			input:       "",
			expectedErr: "Error: <primitive command> が見つかりません。",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := NewContext(test.input)
			node := NewPrimitiveCommandNode()
			err := node.Parse(ctx)
			if test.expectedErr != "" {
				assert.Error(t, err, test.expectedErr, "エラーメッセージが期待通りではありません: "+test.name)
			} else {
				assert.NoError(t, err, "エラーが発生しました: "+test.name)
				assert.Equal(t, test.expected, node.String(), "解析結果が期待通りではありません: "+test.name)
			}
		})
	}
}

func TestRepeatCommandNodeParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr string
	}{
		{
			name:     "valid repeat",
			input:    "repeat 3 go end",
			expected: "[repeat 3 [go]]",
		},
		{
			name:        "missing number",
			input:       "repeat go end",
			expectedErr: "Error: go",
		},
		{
			name:        "missing command list",
			input:       "repeat 3",
			expectedErr: "Error: 'end' が見つかりません。",
		},
		{
			name:        "non-number",
			input:       "repeat abc go end",
			expectedErr: "Error: strconv.Atoi: parsing \"abc\": invalid syntax",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := NewContext(test.input)
			node := NewRepeatCommandNode()
			err := node.Parse(ctx)
			if test.expectedErr != "" {
				assert.Error(t, err, test.expectedErr, "エラーメッセージが期待通りではありません: "+test.name)
			} else {
				assert.NoError(t, err, "エラーが発生しました: "+test.name)
				assert.Equal(t, test.expected, node.String(), "解析結果が期待通りではありません: "+test.name)
			}
		})
	}
}

func TestCommandListNodeParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr string
	}{
		{
			name:     "single command",
			input:    "go end",
			expected: "[go]",
		},
		{
			name:     "multiple commands",
			input:    "go right left end",
			expected: "[go, right, left]",
		},
		{
			name:     "repeat command",
			input:    "repeat 2 go end end",
			expected: "[[repeat 2 [go]]]",
		},
		{
			name:        "missing end",
			input:       "go",
			expectedErr: "Error: 'end' が見つかりません。",
		},
		{
			name:     "empty list",
			input:    "end",
			expected: "[]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := NewContext(test.input)
			node := NewCommandListNode()
			err := node.Parse(ctx)
			if test.expectedErr != "" {
				assert.Error(t, err, test.expectedErr, "エラーメッセージが期待通りではありません: "+test.name)
			} else {
				assert.NoError(t, err, "エラーが発生しました: "+test.name)
				assert.Equal(t, test.expected, node.String(), "解析結果が期待通りではありません: "+test.name)
			}
		})
	}
}

func TestCommandNodeParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "primitive command",
			input:    "go",
			expected: "go",
		},
		{
			name:     "repeat command",
			input:    "repeat 3 left end",
			expected: "[repeat 3 [left]]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := NewContext(test.input)
			node := NewCommandNode()
			err := node.Parse(ctx)
			assert.NoError(t, err, "エラーが発生しました: "+test.name)
			assert.Equal(t, test.expected, node.String(), "解析結果が期待通りではありません: "+test.name)
		})
	}
}
