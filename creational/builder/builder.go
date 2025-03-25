package builder

import "strings"

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
	Result() string
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("一般的なあいさつ")
	d.builder.MakeItems([]string{"How are you?", "Hello.", "Hi."})
	d.builder.MakeString("時間帯に応じたあいさつ")
	d.builder.MakeItems([]string{"Good morning.", "Good afternoon.", "Good evening."})
	d.builder.Close()
}

// TextBuilderがBuilderインターフェースを満たすことを確認
var _ Builder = (*TextBuilder)(nil)

type TextBuilder struct {
	strings.Builder
}

func (t *TextBuilder) MakeTitle(title string) {
	t.WriteString("==============================\n")
	t.WriteString("『" + title + "』\n\n")
}

func (t *TextBuilder) MakeString(str string) {
	t.WriteString("■" + str + "\n\n")
}

func (t *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		t.WriteString("\t・" + item + "\n")
	}
	t.WriteString("\n")
}

func (t *TextBuilder) Close() {
	t.WriteString("==============================\n")
}

func (t *TextBuilder) Result() string {
	return t.String()
}

// HTMLBuilderがBuilderインターフェースを満たすことを確認
var _ Builder = (*TextBuilder)(nil)

type HTMLBuilder struct {
	filename string
	strings.Builder
}

func (h *HTMLBuilder) MakeTitle(title string) {
	h.filename = title + ".html"
	h.WriteString("<!DOCTYPE html>\n<html>\n<head><title>" + title + "</title></head>\n<body>\n<h1>" + title + "</h1>\n\n")
}

func (h *HTMLBuilder) MakeString(str string) {
	h.WriteString("<p>" + str + "</p>\n\n")
}

func (h *HTMLBuilder) MakeItems(items []string) {
	h.WriteString("<ul>\n")
	for _, item := range items {
		h.WriteString("<li>" + item + "</li>\n")
	}
	h.WriteString("</ul>\n\n")
}

func (h *HTMLBuilder) Close() {
	h.WriteString("</body></html>\n")
}

func (h *HTMLBuilder) Result() string {
	return h.String()
}
