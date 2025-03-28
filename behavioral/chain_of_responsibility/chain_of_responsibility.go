package chainofresponsibility

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) *Trouble {
	return &Trouble{number: number}
}

func (t *Trouble) Number() int {
	return t.number
}

func (t *Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.number)
}

type Support interface {
	SetNext(next Support) Support
	Support(trouble *Trouble) string
}

// baseSupportがSupportインターフェースを満たしていることを確認する
var _ Support = (*baseSupport)(nil)

type baseSupport struct {
	name string
	next Support
}

func NewBaseSupport(name string) baseSupport {
	return baseSupport{name: name}
}

func (s *baseSupport) SetNext(next Support) Support {
	s.next = next
	return next
}

func (s *baseSupport) Support(trouble *Trouble) string {
	if s.next != nil {
		return s.next.Support(trouble)
	}
	return s.fail(trouble)
}

func (s *baseSupport) done(trouble *Trouble) string {
	return fmt.Sprintf("%s is Resolved by [%s].", trouble, s.name)
}

func (s *baseSupport) fail(trouble *Trouble) string {
	return fmt.Sprintf("%s cannot be Resolved.", trouble)
}

// SpecialSupportがSupportインターフェースを満たし、Resolveメソッドを持っていることを確認する
var _ interface {
	Support
	Resolve(trouble *Trouble) bool
} = (*SpecialSupport)(nil)

type SpecialSupport struct {
	baseSupport
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	return &SpecialSupport{
		baseSupport: NewBaseSupport(name),
		number:      number,
	}
}

func (s *SpecialSupport) Support(trouble *Trouble) string {
	if s.Resolve(trouble) {
		return s.done(trouble)
	}
	return s.baseSupport.Support(trouble)
}

func (s *SpecialSupport) Resolve(trouble *Trouble) bool {
	return trouble.Number() == s.number
}

// OddSupportがSupportインターフェースを満たし、Resolveメソッドを持っていることを確認する
var _ interface {
	Support
	Resolve(trouble *Trouble) bool
} = (*OddSupport)(nil)

type OddSupport struct {
	baseSupport
}

func NewOddSupport(name string) *OddSupport {
	return &OddSupport{baseSupport: NewBaseSupport(name)}
}

func (s *OddSupport) Support(trouble *Trouble) string {
	if s.Resolve(trouble) {
		return s.done(trouble)
	}
	return s.baseSupport.Support(trouble)
}

func (s *OddSupport) Resolve(trouble *Trouble) bool {
	return trouble.Number()%2 == 1
}

// NoSupportがSupportインターフェースを満たし、Resolveメソッドを持っていることを確認する
var _ interface {
	Support
	Resolve(trouble *Trouble) bool
} = (*NoSupport)(nil)

type NoSupport struct {
	baseSupport
}

func NewNoSupport(name string) *NoSupport {
	return &NoSupport{baseSupport: NewBaseSupport(name)}
}

func (s *NoSupport) Support(trouble *Trouble) string {
	return s.baseSupport.Support(trouble)
}

func (s *NoSupport) Resolve(trouble *Trouble) bool {
	return false
}

// LimitSupportがSupportインターフェースを満たし、Resolveメソッドを持っていることを確認する
var _ interface {
	Support
	Resolve(trouble *Trouble) bool
} = (*LimitSupport)(nil)

type LimitSupport struct {
	baseSupport
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	return &LimitSupport{
		baseSupport: NewBaseSupport(name),
		limit:       limit,
	}
}

func (s *LimitSupport) Support(trouble *Trouble) string {
	if s.Resolve(trouble) {
		return s.done(trouble)
	}
	return s.baseSupport.Support(trouble)
}

func (s *LimitSupport) Resolve(trouble *Trouble) bool {
	return trouble.Number() < s.limit
}

func SupportChain(troubles []*Trouble, firstSupport Support) []string {
	var results []string
	for _, trouble := range troubles {
		results = append(results, firstSupport.Support(trouble))
	}
	return results
}
