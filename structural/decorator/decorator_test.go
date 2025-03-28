package decorator_test

import (
	"testing"

	. "gof-in-go/structural/decorator"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	d1 := NewStringDisplay("Hello, world.")
	assert.Equal(t, 13, d1.GetColumns())
	assert.Equal(t, 1, d1.GetRows())
	assert.Equal(t, "Hello, world.", d1.GetRowText(0))

	d2 := NewSideBorder(d1, '#')
	assert.Equal(t, 15, d2.GetColumns())
	assert.Equal(t, 1, d2.GetRows())
	assert.Equal(t, "#Hello, world.#", d2.GetRowText(0))

	d3 := NewFullBorder(d2)
	assert.Equal(t, 17, d3.GetColumns())
	assert.Equal(t, 3, d3.GetRows())
	assert.Equal(t, "+---------------+", d3.GetRowText(0))
	assert.Equal(t, "|#Hello, world.#|", d3.GetRowText(1))
	assert.Equal(t, "+---------------+", d3.GetRowText(2))

	d4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewFullBorder(
						NewStringDisplay("Hello, world."),
					), '*',
				),
			),
		), '/',
	)
	assert.Equal(t, 23, d4.GetColumns())
	assert.Equal(t, 7, d4.GetRows())
	assert.Equal(t, "/+-------------------+/", d4.GetRowText(0))
	assert.Equal(t, "/|+-----------------+|/", d4.GetRowText(1))
	assert.Equal(t, "/||*+-------------+*||/", d4.GetRowText(2))
	assert.Equal(t, "/||*|Hello, world.|*||/", d4.GetRowText(3))
	assert.Equal(t, "/||*+-------------+*||/", d4.GetRowText(4))
	assert.Equal(t, "/|+-----------------+|/", d4.GetRowText(5))
	assert.Equal(t, "/+-------------------+/", d4.GetRowText(6))
}
