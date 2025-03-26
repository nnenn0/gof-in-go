package bridge_test

import (
	"testing"

	. "gof-in-go/structural/bridge"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	impl1 := NewStringDisplayImpl("Hello, Japan.")
	d1 := NewDisplay(impl1)
	result1 := d1.Display()

	assert.Len(t, result1, 3)
	expect := []string{
		"+-------------+",
		"|Hello, Japan.|",
		"+-------------+",
	}
	assert.Equal(t, expect, result1)

	impl2 := NewStringDisplayImpl("Hello, World.")
	d2 := NewCountDisplay(impl2)

	result2 := d2.Display()

	assert.Len(t, result2, 3)
	expect = []string{
		"+-------------+",
		"|Hello, World.|",
		"+-------------+",
	}
	assert.Equal(t, expect, result2)

	result3 := d2.MultiDisplay(3)

	assert.Len(t, result3, 5)
	expect = []string{
		"+-------------+",
		"|Hello, World.|",
		"|Hello, World.|",
		"|Hello, World.|",
		"+-------------+",
	}
	assert.Equal(t, expect, result3)
}
