package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() Book
}

type Book struct {
	Name string
}

type BookShelf struct {
	books []Book
}

func NewBookShelf() *BookShelf {
	return &BookShelf{books: []Book{}}
}

func (bs *BookShelf) AppendBook(b Book) {
	bs.books = append(bs.books, b)
}

func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{bookShelf: bs}
}

// BookshelfがAggregateインターフェースを満たすことを確認
var _ Aggregate = (*BookShelf)(nil)

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (it *BookShelfIterator) HasNext() bool {
	return it.index < len(it.bookShelf.books)
}

func (it *BookShelfIterator) Next() Book {
	b := it.bookShelf.books[it.index]
	it.index++
	return b
}
