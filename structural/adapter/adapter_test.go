package adapter_test

import (
	"testing"

	. "gof-in-go/structural/adapter"

	"github.com/stretchr/testify/assert"
)

func TestBanner(t *testing.T) {
	b := NewBanner("test")

	assert.Equal(t, "(test)", b.EncloseInParen(), "EncloseInParenの戻り値が期待と異なります")
	assert.Equal(t, "*test*", b.EncloseInAster(), "EncloseInAsterの戻り値が期待と異なります")
}

func TestWrapBanner(t *testing.T) {
	wb := NewWrapBanner("test")

	assert.Equal(t, "(test)", wb.WrapWithWeak(), "WrapWithWeakの戻り値が期待と異なります")
	assert.Equal(t, "*test*", wb.WrapWithStrong(), "WrapWithStrongの戻り値が期待と異なります")
}
