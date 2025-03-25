package abstruct_factory

import (
	"fmt"
	"strings"
)

type Item interface {
	MakeHTML() string
}

type Link interface {
	Item
	GetURL() string
	GetCaption() string
}

type Tray interface {
	Item
	AddItem(Item)
	GetItems() []Item
	GetCaption() string
}

type Page interface {
	AddItem(Item)
	MakeHTML() string
	GetTitle() string
	GetAuthor() string
	GetContent() []Item
}

// DivLinkがLinkインターフェースを満たすことを確認
var _ Link = (*DivLink)(nil)

type DivLink struct {
	caption string
	url     string
}

func NewDivLink(caption, url string) *DivLink {
	return &DivLink{
		caption: caption,
		url:     url,
	}
}

func (d *DivLink) MakeHTML() string {
	return fmt.Sprintf("<div class=\"LINK\"><a href=\"%s\">%s</a></div>\n", d.url, d.caption)
}

func (d *DivLink) GetURL() string {
	return d.url
}

func (d *DivLink) GetCaption() string {
	return d.caption
}

// DivTrayがTrayインターフェースを満たすことを確認
var _ Tray = (*DivTray)(nil)

type DivTray struct {
	caption string
	items   []Item
}

func NewDivTray(caption string) *DivTray {
	return &DivTray{
		caption: caption,
		items:   []Item{},
	}
}

func (d *DivTray) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<p><b>%s</b></p>\n", d.caption))
	sb.WriteString("<div class=\"TRAY\">")
	for _, item := range d.items {
		sb.WriteString(item.MakeHTML())
	}
	sb.WriteString("</div>\n")
	return sb.String()
}

func (d *DivTray) AddItem(item Item) {
	d.items = append(d.items, item)
}

func (d *DivTray) GetItems() []Item {
	return d.items
}

func (d *DivTray) GetCaption() string {
	return d.caption
}

// DivPageがPageインターフェースを満たすことを確認
var _ Page = (*DivPage)(nil)

type DivPage struct {
	title   string
	author  string
	content []Item
}

func NewDivPage(title, author string) *DivPage {
	return &DivPage{
		title:   title,
		author:  author,
		content: []Item{},
	}
}

func (d *DivPage) AddItem(item Item) {
	d.content = append(d.content, item)
}

func (d *DivPage) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html><head><title>")
	sb.WriteString(d.title)
	sb.WriteString("</title><style>\n")
	sb.WriteString("div.TRAY { padding:0.5em; margin-left:5em; border:1px solid black; }\n")
	sb.WriteString("div.LINK { padding:0.5em; background-color: lightgray; }\n")
	sb.WriteString("</style></head><body>\n")
	sb.WriteString(fmt.Sprintf("<h1>%s</h1>\n", d.title))

	for _, item := range d.content {
		sb.WriteString(item.MakeHTML())
	}

	sb.WriteString("<hr><address>")
	sb.WriteString(d.author)
	sb.WriteString("</address>\n")
	sb.WriteString("</body></html>\n")

	return sb.String()
}

func (d *DivPage) GetTitle() string {
	return d.title
}

func (d *DivPage) GetAuthor() string {
	return d.author
}

func (d *DivPage) GetContent() []Item {
	return d.content
}

type Factory interface {
	CreateLink(caption, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title, author string) Page
}

// DivFactoryがFactoryインターフェースを満たすことを確認
var _ Factory = (*DivFactory)(nil)

type DivFactory struct{}

func NewDivFactory() *DivFactory {
	return &DivFactory{}
}

func (f *DivFactory) CreateLink(caption, url string) Link {
	return NewDivLink(caption, url)
}

func (f *DivFactory) CreateTray(caption string) Tray {
	return NewDivTray(caption)
}

func (f *DivFactory) CreatePage(title, author string) Page {
	return NewDivPage(title, author)
}

func GeneratePageContent(page Page) (string, error) {
	html := page.MakeHTML()

	return html, nil
}

// ListLinkがLinkインターフェースを満たすことを確認
var _ Link = (*ListLink)(nil)

type ListLink struct {
	caption string
	url     string
}

func NewListLink(caption, url string) *ListLink {
	return &ListLink{
		caption: caption,
		url:     url,
	}
}

func (l *ListLink) MakeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", l.url, l.caption)
}

func (l *ListLink) GetURL() string {
	return l.url
}

func (l *ListLink) GetCaption() string {
	return l.caption
}

// ListTrayがTrayインターフェースを満たすことを確認
var _ Tray = (*ListTray)(nil)

type ListTray struct {
	caption string
	items   []Item
}

func NewListTray(caption string) *ListTray {
	return &ListTray{
		caption: caption,
		items:   []Item{},
	}
}

func (l *ListTray) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString("<li>\n")
	sb.WriteString(l.caption)
	sb.WriteString("\n<ul>\n")
	for _, item := range l.items {
		sb.WriteString(item.MakeHTML())
	}
	sb.WriteString("</ul>\n")
	sb.WriteString("</li>\n")
	return sb.String()
}

func (l *ListTray) AddItem(item Item) {
	l.items = append(l.items, item)
}

func (l *ListTray) GetItems() []Item {
	return l.items
}

func (l *ListTray) GetCaption() string {
	return l.caption
}

// ListPageがPageインターフェースを満たすことを確認
var _ Page = (*ListPage)(nil)

type ListPage struct {
	title   string
	author  string
	content []Item
}

func NewListPage(title, author string) *ListPage {
	return &ListPage{
		title:   title,
		author:  author,
		content: []Item{},
	}
}

func (l *ListPage) AddItem(item Item) {
	l.content = append(l.content, item)
}

func (l *ListPage) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html><head><title>")
	sb.WriteString(l.title)
	sb.WriteString("</title></head>\n")
	sb.WriteString("<body>\n")
	sb.WriteString(fmt.Sprintf("<h1>%s</h1>\n", l.title))
	sb.WriteString("<ul>\n")

	for _, item := range l.content {
		sb.WriteString(item.MakeHTML())
	}

	sb.WriteString("</ul>\n")
	sb.WriteString("<hr><address>")
	sb.WriteString(l.author)
	sb.WriteString("</address>\n")
	sb.WriteString("</body></html>\n")

	return sb.String()
}

func (l *ListPage) GetTitle() string {
	return l.title
}

func (l *ListPage) GetAuthor() string {
	return l.author
}

func (l *ListPage) GetContent() []Item {
	return l.content
}

// ListFactoryがFactoryインターフェースを満たすことを確認
var _ Factory = (*ListFactory)(nil)

type ListFactory struct{}

func NewListFactory() *ListFactory {
	return &ListFactory{}
}

func (f *ListFactory) CreateLink(caption, url string) Link {
	return NewListLink(caption, url)
}

func (f *ListFactory) CreateTray(caption string) Tray {
	return NewListTray(caption)
}

func (f *ListFactory) CreatePage(title, author string) Page {
	return NewListPage(title, author)
}
