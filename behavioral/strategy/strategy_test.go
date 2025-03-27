package strategy_test

import (
	"testing"

	. "gof-in-go/behavioral/strategy"

	"github.com/stretchr/testify/assert"
)

func TestStrategies(t *testing.T) {
	player1 := NewPlayer("Taro", NewWinningStrategy(314))
	player2 := NewPlayer("Hana", NewProbStrategy(15))

	for range 10000 {
		hand1 := player1.NextHand()
		hand2 := player2.NextHand()

		if hand1.IsStrongerThan(hand2) {
			player1.Win()
			player2.Lose()
		} else if hand2.IsStrongerThan(hand1) {
			player1.Lose()
			player2.Win()
		} else {
			player1.Even()
			player2.Even()
		}
	}

	assert.True(t, player1.GameCount() == 10000)
	assert.True(t, player2.GameCount() == 10000)
}

func TestHandComparison(t *testing.T) {
	assert.True(t, Rock.IsStrongerThan(Scissors))
	assert.True(t, Scissors.IsStrongerThan(Paper))
	assert.True(t, Paper.IsStrongerThan(Rock))
	assert.False(t, Rock.IsStrongerThan(Paper))
	assert.False(t, Scissors.IsStrongerThan(Rock))
	assert.False(t, Paper.IsStrongerThan(Scissors))
	assert.True(t, Rock.IsWeakerThan(Paper))
	assert.True(t, Scissors.IsWeakerThan(Rock))
	assert.True(t, Paper.IsWeakerThan(Scissors))
}

func TestWinningStrategy(t *testing.T) {
	strategy := NewWinningStrategy(42)
	hand1 := strategy.NextHand()
	strategy.Study(true)
	hand2 := strategy.NextHand()
	assert.Equal(t, hand1, hand2)
	strategy.Study(false)
	hand3 := strategy.NextHand()
	assert.NotEqual(t, hand2, hand3)
}

func TestProbStrategy(t *testing.T) {
	strategy := NewProbStrategy(42)
	prevHand := strategy.NextHand()
	strategy.Study(true)
	newHand := strategy.NextHand()
	assert.NotNil(t, newHand)
	assert.NotEqual(t, -1, newHand)
	assert.Equal(t, prevHand, newHand)
}
