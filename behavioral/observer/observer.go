package observer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Observer interface {
	Update(generator NumberGenerator) string
}

type DigitObserver struct{}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (d *DigitObserver) Update(generator NumberGenerator) string {
	output := fmt.Sprintf("DigitObserver:%d", generator.GetNumber())
	return output
}

type GraphObserver struct{}

func NewGraphObserver() *GraphObserver {
	return &GraphObserver{}
}

func (g *GraphObserver) Update(generator NumberGenerator) string {
	count := generator.GetNumber()
	graph := "GraphObserver:"
	for range count {
		graph += "*"
	}
	return graph
}

type NumberGenerator interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	NotifyObservers() []string
	GetNumber() int
	Execute()
}

type BaseNumberGenerator struct {
	observers []Observer
	mu        sync.Mutex
}

func (b *BaseNumberGenerator) AddObserver(observer Observer) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.observers = append(b.observers, observer)
}

func (b *BaseNumberGenerator) DeleteObserver(observer Observer) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i, obs := range b.observers {
		if obs == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *BaseNumberGenerator) NotifyObservers() []string {
	return []string{}
}

type RandomNumberGenerator struct {
	BaseNumberGenerator
	random  *rand.Rand
	number  int
	outputs []string
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		random:  rand.New(rand.NewSource(time.Now().UnixNano())),
		outputs: make([]string, 0),
	}
}

func (r *RandomNumberGenerator) GetNumber() int {
	return r.number
}

func (r *RandomNumberGenerator) Execute() {
	for range 20 {
		r.number = r.random.Intn(50)
		r.notifyObserversConcrete()
	}
}

func (r *RandomNumberGenerator) notifyObserversConcrete() {
	r.BaseNumberGenerator.mu.Lock()
	defer r.BaseNumberGenerator.mu.Unlock()
	for _, observer := range r.BaseNumberGenerator.observers {
		r.outputs = append(r.outputs, observer.Update(r))
	}
}

func (r *RandomNumberGenerator) NotifyObservers() []string {
	return r.outputs
}
