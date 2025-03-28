package proxy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "gof-in-go/structural/proxy"
)

func TestPrinterProxy(t *testing.T) {
	p := NewPrinterProxyWithName("Alice")

	assert.Equal(t, "Alice", p.GetPrinterName(), "初期の名前が正しくありません")

	p.SetPrinterName("Bob")
	assert.Equal(t, "Bob", p.GetPrinterName(), "名前の変更が反映されていません")

	output := p.Print("Hello, world.")
	expectedOutput := "=== Bob ===\nHello, world."
	assert.Equal(t, expectedOutput, output, "出力内容が期待通りではありません")
}
