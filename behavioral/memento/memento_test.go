package memento_test

import (
	"testing"
	"time"

	. "gof-in-go/behavioral/memento"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGameSimulation(t *testing.T) {
	initialMoney := 100
	gamer := NewGamer(initialMoney)
	memento := gamer.CreateMemento()

	require.NotNil(t, memento, "Memento should not be nil")

	for range 10 {
		outcome := gamer.Bet()
		_ = outcome

		if gamer.Money() > memento.Money() {
			memento = gamer.CreateMemento()
			require.NotNil(t, memento, "新しいMementoはnilであってはならない")
			assert.GreaterOrEqual(t, gamer.Money(), memento.Money(), "所持金は新しいMementoで大きくなければならない")
		} else if gamer.Money() < memento.Money()/2 {
			gamer.RestoreMemento(memento)
			assert.Equal(t, memento.Money(), gamer.Money(), "所持金は以前の状態に復元されなければならない")
		}
		time.Sleep(10 * time.Millisecond)
	}

	assert.GreaterOrEqual(t, gamer.Money(), 0, "ゲーマーの所持金は0以上でなければならない")
}
