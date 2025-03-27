package strategy

import (
	"math/rand"
)

type Hand int

const (
	Rock Hand = iota
	Scissors
	Paper
)

var hands = []Hand{Rock, Scissors, Paper}

func GetHand(value int) Hand {
	return hands[value]
}

func (h Hand) IsStrongerThan(other Hand) bool {
	return h.fight(other) == 1
}

func (h Hand) IsWeakerThan(other Hand) bool {
	return h.fight(other) == -1
}

func (h Hand) fight(other Hand) int {
	if h == other {
		return 0
	} else if (h+1)%3 == other {
		return 1
	} else {
		return -1
	}
}

type Strategy interface {
	NextHand() Hand
	Study(win bool)
}

// WinningStrategyがStrategyインターフェースを満たしていることを確認
var _ Strategy = (*WinningStrategy)(nil)

type WinningStrategy struct {
	rand     *rand.Rand
	won      bool
	prevHand Hand
}

func NewWinningStrategy(seed int64) *WinningStrategy {
	return &WinningStrategy{
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (s *WinningStrategy) NextHand() Hand {
	if !s.won {
		for {
			newHand := GetHand(s.rand.Intn(3))
			if s.prevHand != newHand {
				s.prevHand = GetHand(s.rand.Intn(3))
				break
			}
		}
	}
	return s.prevHand
}

func (s *WinningStrategy) Study(win bool) {
	s.won = win
}

type ProbStrategy struct {
	rand             *rand.Rand
	prevHandValue    int
	currentHandValue int
	history          [3][3]int
}

// WinningStrategyがStrategyインターフェースを満たしていることを確認
var _ Strategy = (*ProbStrategy)(nil)

func NewProbStrategy(seed int64) *ProbStrategy {
	return &ProbStrategy{
		rand: rand.New(rand.NewSource(seed)),
		history: [3][3]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (s *ProbStrategy) NextHand() Hand {
	sum := s.getSum(s.currentHandValue)
	bet := s.rand.Intn(sum)
	var handValue int
	if bet < s.history[s.currentHandValue][0] {
		handValue = 0
	} else if bet < s.history[s.currentHandValue][0]+s.history[s.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}

	s.prevHandValue = s.currentHandValue
	s.currentHandValue = handValue

	return GetHand(handValue)
}

func (s *ProbStrategy) getSum(handValue int) int {
	sum := 0
	for i := range 3 {
		sum += s.history[handValue][i]
	}
	return sum
}

func (s *ProbStrategy) Study(win bool) {
	if win {
		s.history[s.prevHandValue][s.currentHandValue]++
	} else {
		s.history[s.prevHandValue][(s.currentHandValue+1)%3]++
		s.history[s.prevHandValue][(s.currentHandValue+2)%3]++
	}
}

type Player struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		name:     name,
		strategy: strategy,
	}
}

func (p *Player) NextHand() Hand {
	return p.strategy.NextHand()
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.gameCount++
}

func (p *Player) GameCount() int {
	return p.gameCount
}
