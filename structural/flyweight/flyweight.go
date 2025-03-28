package flyweight

import (
	"strings"
	"sync"
)

type BigChar struct {
	charname rune
	fontdata string
}

func NewBigChar(charname rune, fontdata string) *BigChar {
	return &BigChar{charname: charname, fontdata: fontdata}
}

func (bc *BigChar) String() string {
	return bc.fontdata
}

type BigCharFactory struct {
	pool     map[rune]*BigChar
	fontData map[rune]string
	mu       sync.Mutex
}

var singleton *BigCharFactory
var once sync.Once

func GetBigCharFactory(initialFontData map[rune]string) *BigCharFactory {
	once.Do(func() {
		singleton = &BigCharFactory{
			pool:     make(map[rune]*BigChar),
			fontData: initialFontData,
		}
	})
	return singleton
}

func ResetBigCharFactory() {
	once = sync.Once{}
	singleton = nil
}

func (f *BigCharFactory) GetBigChar(charname rune) *BigChar {
	f.mu.Lock()
	defer f.mu.Unlock()
	if bc, ok := f.pool[charname]; ok {
		return bc
	}
	font, ok := f.fontData[charname]
	if !ok {
		font = string(charname) + "?"
	}
	bc := NewBigChar(charname, font)
	f.pool[charname] = bc
	return bc
}

type BigString struct {
	bigchars []*BigChar
}

func NewBigString(s string) *BigString {
	factory := GetBigCharFactory(nil)
	bigchars := make([]*BigChar, len(s))
	for i, r := range s {
		bigchars[i] = factory.GetBigChar(r)
	}
	return &BigString{bigchars: bigchars}
}

func (bs *BigString) String() string {
	var sb strings.Builder
	for _, bc := range bs.bigchars {
		sb.WriteString(bc.String())
	}
	return sb.String()
}

func ProcessDigits(digits string, fontData map[rune]string) string {
	ResetBigCharFactory()
	GetBigCharFactory(fontData)
	if len(digits) == 0 {
		return "Usage: flyweight digits\nExample: flyweight 1212123\n"
	}
	bs := NewBigString(digits)
	return bs.String()
}
