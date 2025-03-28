package flyweight_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "gof-in-go/structural/flyweight"
)

func TestBigChar_String(t *testing.T) {
	bcA := NewBigChar('A', "##\n#.\n##\n")
	assert.Equal(t, "##\n#.\n##\n", bcA.String(), "BigChar 'A' の文字データが正しいこと")

	bc1 := NewBigChar('1', "#\n#\n#\n")
	assert.Equal(t, "#\n#\n#\n", bc1.String(), "BigChar '1' の文字データが正しいこと")
}

func TestBigCharFactory_GetInstance(t *testing.T) {
	ResetBigCharFactory()
	factory1 := GetBigCharFactory(map[rune]string{})
	factory2 := GetBigCharFactory(map[rune]string{})
	assert.Same(t, factory1, factory2, "BigCharFactory はシングルトンであること")
}

func TestBigCharFactory_GetBigChar(t *testing.T) {
	ResetBigCharFactory()
	fontData := map[rune]string{
		'B': "**\n*.\n**\n",
	}
	factory := GetBigCharFactory(fontData)
	bc1 := factory.GetBigChar('B')
	bc2 := factory.GetBigChar('B')
	assert.Same(t, bc1, bc2, "同じ文字に対しては同じ BigChar インスタンスが返されること")
	assert.Equal(t, "**\n*.\n**\n", bc1.String(), "BigChar の内容が正しいこと")
}

func TestBigString_String(t *testing.T) {
	ResetBigCharFactory()
	fontData := map[rune]string{
		'1': "#\n#\n#\n",
		'2': "#-\n#-\n--\n",
	}
	GetBigCharFactory(fontData)
	bs := NewBigString("121")
	expected := "#\n#\n#\n#-\n#-\n--\n#\n#\n#\n"
	assert.Equal(t, expected, bs.String(), "BigString の文字列表現が正しいこと")
}

func TestProcessDigits(t *testing.T) {
	fontData := map[rune]string{
		'1': "#\n#\n#\n",
		'2': "#-\n#-\n--\n",
	}

	expectedUsage := "Usage: flyweight digits\nExample: flyweight 1212123\n"
	actualUsage := ProcessDigits("", fontData)
	assert.Equal(t, expectedUsage, actualUsage, "引数がない場合は Usage メッセージが返されること")

	expectedOutput := "#\n#\n#\n#-\n#-\n--\n#\n#\n#\n"
	actualOutput := ProcessDigits("121", fontData)
	assert.Equal(t, expectedOutput, actualOutput, "引数が数字の場合は対応する大きな文字列表現が返されること")

	actualUnknown := ProcessDigits("3", fontData)
	assert.Contains(t, actualUnknown, "3?", "存在しない文字に対しては '?' が含まれること")
}
