package factory_method

import "fmt"

type Product interface {
	Use() string
}

type Factory interface {
	Create(owner string) Product
}

// IDCardがProductインターフェースを満たすことを確認
var _ Product = (*IDCard)(nil)

type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	return &IDCard{owner: owner}
}

func (c *IDCard) Use() string {
	return fmt.Sprintf("%sのカードを使います。", c.owner)
}

// IDCardFactoryがFactoryインターフェースを満たすことを確認
var _ Factory = (*IDCardFactory)(nil)

type IDCardFactory struct{}

func (f *IDCardFactory) Create(owner string) Product {
	card := NewIDCard(owner)
	return card
}
