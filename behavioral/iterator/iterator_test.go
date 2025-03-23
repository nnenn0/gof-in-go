package iterator_test

import (
	"testing"

	. "gof-in-go/behavioral/iterator"

	"github.com/stretchr/testify/assert"
)

func TestBookShelfLen(t *testing.T) {
	book1 := Book{Name: "book1"}
	book2 := Book{Name: "book2"}
	book3 := Book{Name: "book3"}

	bookshelf := NewBookShelf()
	bookshelf.AppendBook(book1)
	bookshelf.AppendBook(book2)
	bookshelf.AppendBook(book3)

	assert.Equal(t, 3, bookshelf.Len(), "本棚のサイズが正しくありません")
}

func TestBookShelfIterator(t *testing.T) {
	book1 := Book{Name: "book1"}
	book2 := Book{Name: "book2"}
	book3 := Book{Name: "book3"}

	bookshelf := NewBookShelf()
	bookshelf.AppendBook(book1)
	bookshelf.AppendBook(book2)
	bookshelf.AppendBook(book3)

	it := bookshelf.Iterator()
	bookNames := []string{book1.Name, book2.Name, book3.Name}
	index := 0

	for it.HasNext() {
		b := it.Next()
		assert.Equal(t, bookNames[index], b.Name, "予期しない本の名前が含まれています")
		index++
	}

	assert.Equal(t, len(bookNames), index, "イテレーターがすべての本を巡回していません")
}

func TestEmptyBookShelfIterator(t *testing.T) {
	bookshelf := NewBookShelf()
	it := bookshelf.Iterator()
	assert.False(t, it.HasNext(), "空の本棚でHasNext()がtrueを返しました")
}

func TestIteratorExhaustion(t *testing.T) {
	bookshelf := NewBookShelf()
	bookshelf.AppendBook(Book{Name: "book1"})
	bookshelf.AppendBook(Book{Name: "book2"})

	it := bookshelf.Iterator()
	_ = it.Next()
	_ = it.Next()
	assert.False(t, it.HasNext(), "すべての本を取り出したのに HasNext() が true を返しました")
}
