package facade

import (
	"errors"
	"strings"
)

type Database struct {
	data string
}

func NewDatabase(data string) Database {
	return Database{data: data}
}

func (d Database) GetProperties() (map[string]string, error) {
	properties := make(map[string]string)
	lines := strings.Split(d.data, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			properties[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return properties, nil
}

type HtmlWriter struct {
	builder strings.Builder
}

func (w *HtmlWriter) Title(title string) {
	w.builder.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<title>" + title + "</title>\n</head>\n<body>\n<h1>" + title + "</h1>\n")
}

func (w *HtmlWriter) Paragraph(msg string) {
	w.builder.WriteString("<p>" + msg + "</p>\n")
}

func (w *HtmlWriter) Link(href, caption string) {
	w.Paragraph("<a href=\"" + href + "\">" + caption + "</a>")
}

func (w *HtmlWriter) Mailto(mailaddr, username string) {
	w.Link("mailto:"+mailaddr, username)
}

func (w *HtmlWriter) Close() string {
	w.builder.WriteString("</body>\n</html>\n")
	return w.builder.String()
}

type PageMaker struct {
	db Database
}

func NewPageMaker(db Database) PageMaker {
	return PageMaker{db: db}
}

func (p PageMaker) MakeWelcomePage(mailaddr string) (string, error) {
	properties, err := p.db.GetProperties()
	if err != nil {
		return "", err
	}
	username, ok := properties[mailaddr]
	if !ok {
		return "", errors.New("email address not found")
	}

	writer := &HtmlWriter{}
	writer.Title(username + "'s web page")
	writer.Paragraph("Welcome to " + username + "'s web page!")
	writer.Paragraph("Nice to meet you!")
	writer.Mailto(mailaddr, username)
	return writer.Close(), nil
}
