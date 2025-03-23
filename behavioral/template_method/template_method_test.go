package template_method_test

import (
	"testing"

	. "gof-in-go/behavioral/template_method"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	charTemplate := NewTemplate(NewChar('H'))
	assert.Equal(t, "<<HHHHH>>\n", charTemplate.Format(), "charが正しくフォーマットされていません")

	strTemplate := NewTemplate(NewStr("Hello, world."))
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
