package composite

import (
	"fmt"
	"strings"
)

type Entry interface {
	GetName() string
	GetSize() int
	PrintList(prefix string) string
}

// FileがEntryインターフェースを満たしていることを確認
var _ Entry = (*File)(nil)

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) PrintList(prefix string) string {
	return fmt.Sprintf("%s/%s (%d)", prefix, f.name, f.size)
}

// DirectoryがEntryインターフェースを満たしていることを確認
var _ Entry = (*Directory)(nil)

type Directory struct {
	name      string
	directory []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	size := 0
	for _, entry := range d.directory {
		size += entry.GetSize()
	}
	return size
}

func (d *Directory) PrintList(prefix string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s/%s (%d)\n", prefix, d.name, d.GetSize()))
	for _, entry := range d.directory {
		builder.WriteString(entry.PrintList(prefix+"/"+d.name) + "\n")
	}
	return strings.TrimSpace(builder.String())
}

func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}
