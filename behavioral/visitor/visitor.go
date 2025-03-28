package visitor

import "fmt"

type Element interface {
	Accept(visitor Visitor) string
}

type Entry interface {
	Element
	GetName() string
	GetSize() int
}

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

func (d *Directory) Add(entry Entry) {
	d.directory = append(d.directory, entry)
}

func (d *Directory) Accept(visitor Visitor) string {
	return visitor.VisitDirectory(d)
}

func (d *Directory) Iterator() []Entry {
	return d.directory
}

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

func (f *File) Accept(visitor Visitor) string {
	return visitor.VisitFile(f)
}

type Visitor interface {
	VisitFile(file *File) string
	VisitDirectory(directory *Directory) string
}

type ListVisitor struct {
	currentDir string
}

func (lv *ListVisitor) VisitFile(file *File) string {
	return lv.currentDir + "/" + file.GetName() + " (" + fmt.Sprint(file.GetSize()) + ")"
}

func (lv *ListVisitor) VisitDirectory(directory *Directory) string {
	result := lv.currentDir + "/" + directory.GetName() + " (" + fmt.Sprint(directory.GetSize()) + ")"
	savedDir := lv.currentDir
	lv.currentDir = lv.currentDir + "/" + directory.GetName()
	for _, entry := range directory.Iterator() {
		result += "\n" + entry.Accept(lv)
	}
	lv.currentDir = savedDir
	return result
}
