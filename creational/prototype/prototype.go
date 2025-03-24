package prototype

import (
	"errors"
	"strings"
)

type Product interface {
	Use(s string) string
	CreateCopy() (Product, error)
}

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{showcase: make(map[string]Product)}
}

func (m *Manager) Register(name string, prototype Product) {
	m.showcase[name] = prototype
}

func (m *Manager) Create(prototypeName string) (Product, error) {
	p, exists := m.showcase[prototypeName]
	if !exists {
		return nil, errors.New("prototype not found")
	}
	return p.CreateCopy()
}

// MessageBoxがProductインターフェースを満たすことを確認
var _ Product = (*MessageBox)(nil)

type MessageBox struct {
	decochar rune
}

func NewMessageBox(decochar rune) *MessageBox {
	return &MessageBox{decochar: decochar}
}

func (m *MessageBox) Use(s string) string {
	decolen := len(s) + 2
	border := strings.Repeat(string(m.decochar), decolen)
	return border + "\n" + string(m.decochar) + s + string(m.decochar) + "\n" + border
}

func (m *MessageBox) CreateCopy() (Product, error) {
	return &MessageBox{decochar: m.decochar}, nil
}

// UnderlinePenがProductインターフェースを満たすことを確認
var _ Product = (*UnderlinePen)(nil)

type UnderlinePen struct {
	ulchar rune
}

func NewUnderlinePen(ulchar rune) *UnderlinePen {
	return &UnderlinePen{ulchar: ulchar}
}

func (u *UnderlinePen) Use(s string) string {
	underline := strings.Repeat(string(u.ulchar), len(s))
	return s + "\n" + underline
}

func (u *UnderlinePen) CreateCopy() (Product, error) {
	return &UnderlinePen{ulchar: u.ulchar}, nil
}
