package memento

import (
	"fmt"
	"math/rand"
	"time"
)

type Gamer struct {
	money      int
	fruits     []string
	rand       *rand.Rand
	fruitsName []string
}

func NewGamer(money int) *Gamer {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return &Gamer{
		money:      money,
		fruits:     []string{},
		rand:       r,
		fruitsName: []string{"リンゴ", "ぶどう", "バナナ", "みかん"},
	}
}

func (g *Gamer) Money() int {
	return g.money
}

func (g *Gamer) Bet() string {
	dice := g.rand.Intn(6) + 1
	switch dice {
	case 1:
		g.money += 100
		return "所持金が増えました。"
	case 2:
		g.money /= 2
		return "所持金が半分になりました。"
	case 6:
		fruit := g.getFruit()
		g.fruits = append(g.fruits, fruit)
		return fmt.Sprintf("フルーツ(%s)をもらいました。", fruit)
	default:
		return "何も起こりませんでした。"
	}
}

func (g *Gamer) CreateMemento() *Memento {
	m := NewMemento(g.money)
	for _, f := range g.fruits {
		if len(f) > 6 && f[:6] == "おいしい" {
			m.AddFruit(f)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.Money()
	g.fruits = memento.Fruits()
}

func (g *Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.money, g.fruits)
}

func (g *Gamer) getFruit() string {
	fruitName := g.fruitsName[g.rand.Intn(len(g.fruitsName))]
	if g.rand.Float64() < 0.5 {
		return "おいしい" + fruitName
	}
	return fruitName
}

type Memento struct {
	money  int
	fruits []string
}

func NewMemento(money int) *Memento {
	return &Memento{
		money:  money,
		fruits: []string{},
	}
}

func (m *Memento) Money() int {
	return m.money
}

func (m *Memento) AddFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) Fruits() []string {
	fruits := make([]string, len(m.fruits))
	copy(fruits, m.fruits)
	return fruits
}
