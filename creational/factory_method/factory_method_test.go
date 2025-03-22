package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDCardFactory(t *testing.T) {
	var factory Factory = &IDCardFactory{}

	card1 := factory.Create("Alice")
	card2 := factory.Create("Bob")
	card3 := factory.Create("Carol")

	assert.Equal(t, "Aliceのカードを使います。", card1.Use())
	assert.Equal(t, "Bobのカードを使います。", card2.Use())
	assert.Equal(t, "Carolのカードを使います。", card3.Use())
}
