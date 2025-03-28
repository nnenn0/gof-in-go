package chainofresponsibility_test

import (
	"testing"

	. "gof-in-go/behavioral/chain_of_responsibility"

	"github.com/stretchr/testify/assert"
)

func TestSupportChain(t *testing.T) {
	// サポートチェーンの作成
	alice := NewNoSupport("Alice")
	bob := NewLimitSupport("Bob", 100)
	charlie := NewSpecialSupport("Charlie", 429)
	diana := NewLimitSupport("Diana", 200)
	elmo := NewOddSupport("Elmo")
	fred := NewLimitSupport("Fred", 300)

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	var troubles []*Trouble
	for i := 0; i < 500; i += 33 {
		troubles = append(troubles, NewTrouble(i))
	}

	results := SupportChain(troubles, alice)

	expectedResults := []string{
		"[Trouble 0] is Resolved by [Bob].",
		"[Trouble 33] is Resolved by [Bob].",
		"[Trouble 66] is Resolved by [Bob].",
		"[Trouble 99] is Resolved by [Bob].",
		"[Trouble 132] is Resolved by [Diana].",
		"[Trouble 165] is Resolved by [Diana].",
		"[Trouble 198] is Resolved by [Diana].",
		"[Trouble 231] is Resolved by [Elmo].",
		"[Trouble 264] is Resolved by [Fred].",
		"[Trouble 297] is Resolved by [Elmo].",
		"[Trouble 330] cannot be Resolved.",
		"[Trouble 363] is Resolved by [Elmo].",
		"[Trouble 396] cannot be Resolved.",
		"[Trouble 429] is Resolved by [Charlie].",
		"[Trouble 462] cannot be Resolved.",
		"[Trouble 495] is Resolved by [Elmo].",
	}

	assert.Equal(t, expectedResults, results)
}

func TestTrouble(t *testing.T) {
	trouble := NewTrouble(42)
	assert.Equal(t, 42, trouble.Number())
	assert.Equal(t, "[Trouble 42]", trouble.String())
}

func TestSpecialSupport(t *testing.T) {
	support := NewSpecialSupport("Test", 100)
	assert.True(t, support.Resolve(NewTrouble(100)))
	assert.False(t, support.Resolve(NewTrouble(99)))
}

func TestOddSupport(t *testing.T) {
	support := NewOddSupport("Test")
	assert.True(t, support.Resolve(NewTrouble(1)))
	assert.True(t, support.Resolve(NewTrouble(99)))
	assert.False(t, support.Resolve(NewTrouble(2)))
	assert.False(t, support.Resolve(NewTrouble(100)))
}

func TestNoSupport(t *testing.T) {
	support := NewNoSupport("Test")
	assert.False(t, support.Resolve(NewTrouble(42)))
}

func TestLimitSupport(t *testing.T) {
	support := NewLimitSupport("Test", 100)
	assert.True(t, support.Resolve(NewTrouble(0)))
	assert.True(t, support.Resolve(NewTrouble(99)))
	assert.False(t, support.Resolve(NewTrouble(100)))
	assert.False(t, support.Resolve(NewTrouble(101)))
}
