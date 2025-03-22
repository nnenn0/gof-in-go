package templatemethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	charTemplate := &Template{formatter: &Char{ch: 'H'}}
	assert.Equal(t, "<<HHHHH>>\n", charTemplate.Format(), "charが正しくフォーマットされていません")

	strTemplate := &Template{formatter: NewStr("Hello, world.")}
	expected := `+-------------+
|Hello, world.|
|Hello, world.|
|Hello, world.|
|Hello, world.|
|Hello, world.|
+-------------+
`
	assert.Equal(t, expected, strTemplate.Format(), "strが正しくフォーマットされていません")
}
