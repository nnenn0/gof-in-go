package builder_test

import (
	"testing"

	. "gof-in-go/creational/builder"

	"github.com/stretchr/testify/assert"
)

func TestTextBuilder(t *testing.T) {
	t.Run("MakeTitle", func(t *testing.T) {
		textBuilder := &TextBuilder{}
		textBuilder.MakeTitle("Greeting")
		result := textBuilder.Result()
		expected := `==============================
『Greeting』

`
		assert.Equal(t, expected, result)
	})

	t.Run("MakeString", func(t *testing.T) {
		textBuilder := &TextBuilder{}
		textBuilder.MakeString("一般的なあいさつ")
		result := textBuilder.Result()
		expected := `■一般的なあいさつ

`
		assert.Equal(t, expected, result)

		textBuilder.MakeString("時間帯に応じたあいさつ")
		result = textBuilder.Result()
		expected = `■一般的なあいさつ

■時間帯に応じたあいさつ

`
		assert.Equal(t, expected, result)
	})

	t.Run("MakeItems", func(t *testing.T) {

		textBuilder := &TextBuilder{}
		textBuilder.MakeItems([]string{"How are you?", "Hello.", "Hi."})
		result := textBuilder.Result()
		expected := `	・How are you?
	・Hello.
	・Hi.

`
		assert.Equal(t, expected, result)

		textBuilder.MakeItems([]string{"Good morning.", "Good afternoon.", "Good evening."})
		result = textBuilder.Result()
		expected = `	・How are you?
	・Hello.
	・Hi.

	・Good morning.
	・Good afternoon.
	・Good evening.

`
		assert.Equal(t, expected, result)
	})

	t.Run("Close", func(t *testing.T) {

		textBuilder := &TextBuilder{}
		textBuilder.Close()
		result := textBuilder.Result()
		expected := `==============================
`
		assert.Equal(t, expected, result)
	})
}

func TestHTMLBuilder(t *testing.T) {
	t.Run("MakeTitle", func(t *testing.T) {
		htmlBuilder := &HTMLBuilder{}
		htmlBuilder.MakeTitle("Greeting")
		result := htmlBuilder.Result()
		expected := `<!DOCTYPE html>
<html>
<head><title>Greeting</title></head>
<body>
<h1>Greeting</h1>

`
		assert.Equal(t, expected, result)
	})

	t.Run("MakeString", func(t *testing.T) {
		htmlBuilder := &HTMLBuilder{}
		htmlBuilder.MakeString("一般的なあいさつ")
		result := htmlBuilder.Result()
		expected := `<p>一般的なあいさつ</p>

`
		assert.Equal(t, expected, result)

		htmlBuilder.MakeString("時間帯に応じたあいさつ")
		result = htmlBuilder.Result()
		expected = `<p>一般的なあいさつ</p>

<p>時間帯に応じたあいさつ</p>

`

		assert.Equal(t, expected, result)
	})

	t.Run("MakeItems", func(t *testing.T) {

		htmlBuilder := &HTMLBuilder{}
		htmlBuilder.MakeItems([]string{"How are you?", "Hello.", "Hi."})
		result := htmlBuilder.Result()
		expected := `<ul>
<li>How are you?</li>
<li>Hello.</li>
<li>Hi.</li>
</ul>

`

		assert.Equal(t, expected, result)

		htmlBuilder.MakeItems([]string{"Good morning.", "Good afternoon.", "Good evening."})
		result = htmlBuilder.Result()
		expected = `<ul>
<li>How are you?</li>
<li>Hello.</li>
<li>Hi.</li>
</ul>

<ul>
<li>Good morning.</li>
<li>Good afternoon.</li>
<li>Good evening.</li>
</ul>

`
		assert.Equal(t, expected, result)
	})

	t.Run("Close", func(t *testing.T) {

		htmlBuilder := &HTMLBuilder{}
		htmlBuilder.Close()
		result := htmlBuilder.Result()
		expected := `</body></html>
`
		assert.Equal(t, expected, result)
	})
}

func TestDirector(t *testing.T) {
	t.Run("TextBuilder", func(t *testing.T) {
		textBuilder := &TextBuilder{}
		director := NewDirector(textBuilder)
		director.Construct()
		result := textBuilder.Result()
		expected := `==============================
『Greeting』

■一般的なあいさつ

	・How are you?
	・Hello.
	・Hi.

■時間帯に応じたあいさつ

	・Good morning.
	・Good afternoon.
	・Good evening.

==============================
`
		assert.Equal(t, expected, result)
	})

	t.Run("HTMLBuilder", func(t *testing.T) {
		htmlBuilder := &HTMLBuilder{}
		director := NewDirector(htmlBuilder)
		director.Construct()
		result := htmlBuilder.Result()
		expected := `<!DOCTYPE html>
<html>
<head><title>Greeting</title></head>
<body>
<h1>Greeting</h1>

<p>一般的なあいさつ</p>

<ul>
<li>How are you?</li>
<li>Hello.</li>
<li>Hi.</li>
</ul>

<p>時間帯に応じたあいさつ</p>

<ul>
<li>Good morning.</li>
<li>Good afternoon.</li>
<li>Good evening.</li>
</ul>

</body></html>
`
		assert.Equal(t, expected, result)
	})
}
