package prototype_test

import (
	"testing"

	. "gof-in-go/creational/prototype"

	"github.com/stretchr/testify/assert"
)

func TestUnderlinePen(t *testing.T) {
	upen := NewUnderlinePen('-')
	expected := `Hello, world.
-------------`
	assert.Equal(t, expected, upen.Use("Hello, world."))
}

func TestMessageBox(t *testing.T) {
	mbox := NewMessageBox('*')
	expected := `***************
*Hello, world.*
***************`
	assert.Equal(t, expected, mbox.Use("Hello, world."))
}

func TestPrototypePattern(t *testing.T) {
	manager := NewManager()
	upen := NewUnderlinePen('-')
	mbox := NewMessageBox('*')
	sbox := NewMessageBox('/')

	manager.Register("strong message", upen)
	manager.Register("warning box", mbox)
	manager.Register("slash box", sbox)

	p1, err := manager.Create("strong message")
	assert.NoError(t, err)
	assert.Equal(t, upen.Use("Hello, world."), p1.Use("Hello, world."))

	p2, err := manager.Create("warning box")
	assert.NoError(t, err)
	assert.Equal(t, mbox.Use("Hello, world."), p2.Use("Hello, world."))

	p3, err := manager.Create("slash box")
	assert.NoError(t, err)
	assert.Equal(t, sbox.Use("Hello, world."), p3.Use("Hello, world."))

	_, err = manager.Create("unknown")
	assert.Error(t, err)
	assert.Equal(t, "prototype not found", err.Error())
}
