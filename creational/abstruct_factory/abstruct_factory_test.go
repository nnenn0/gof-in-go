package abstruct_factory_test

import (
	"testing"

	. "gof-in-go/creational/abstruct_factory"

	"github.com/stretchr/testify/assert"
)

func TestDivFactoryPageGeneration(t *testing.T) {
	factory := NewDivFactory()

	blog1 := factory.CreateLink("Blog 1", "https://example.com/blog1")
	blog2 := factory.CreateLink("Blog 2", "https://example.com/blog2")
	blogTray := factory.CreateTray("Blog Site")
	blogTray.AddItem(blog1)
	blogTray.AddItem(blog2)

	page := factory.CreatePage("Blog Site", "Tech Writer")
	page.AddItem(blogTray)

	content, err := GeneratePageContent(page)
	assert.NoError(t, err)

	expected := `<!DOCTYPE html>
<html><head><title>Blog Site</title><style>
div.TRAY { padding:0.5em; margin-left:5em; border:1px solid black; }
div.LINK { padding:0.5em; background-color: lightgray; }
</style></head><body>
<h1>Blog Site</h1>
<p><b>Blog Site</b></p>
<div class="TRAY"><div class="LINK"><a href="https://example.com/blog1">Blog 1</a></div>
<div class="LINK"><a href="https://example.com/blog2">Blog 2</a></div>
</div>
<hr><address>Tech Writer</address>
</body></html>
`

	assert.Equal(t, expected, content)
}

func TestLinkFactoryPageGeneration(t *testing.T) {
	factory := NewListFactory()

	blog1 := factory.CreateLink("Blog 1", "https://example.com/blog1")
	blog2 := factory.CreateLink("Blog 2", "https://example.com/blog2")
	blogTray := factory.CreateTray("Blog Site")
	blogTray.AddItem(blog1)
	blogTray.AddItem(blog2)

	page := factory.CreatePage("Blog Site", "Tech Writer")
	page.AddItem(blogTray)

	content, err := GeneratePageContent(page)
	assert.NoError(t, err)

	expected := `<!DOCTYPE html>
<html><head><title>Blog Site</title></head>
<body>
<h1>Blog Site</h1>
<ul>
<li>
Blog Site
<ul>
  <li><a href="https://example.com/blog1">Blog 1</a></li>
  <li><a href="https://example.com/blog2">Blog 2</a></li>
</ul>
</li>
</ul>
<hr><address>Tech Writer</address>
</body></html>
`

	assert.Equal(t, expected, content)
}

func TestMainFlowSimulation(t *testing.T) {
	factory := NewListFactory()

	blog1 := factory.CreateLink("Blog 1", "https://example.com/blog1")
	blog2 := factory.CreateLink("Blog 2", "https://example.com/blog2")
	blog3 := factory.CreateLink("Blog 3", "https://example.com/blog3")
	blogTray := factory.CreateTray("Blog Site")
	blogTray.AddItem(blog1)
	blogTray.AddItem(blog2)
	blogTray.AddItem(blog3)

	news1 := factory.CreateLink("News 1", "https://example.com/news1")
	news2 := factory.CreateLink("News 2", "https://example.com/news2")
	news3 := factory.CreateTray("News 3")
	news3.AddItem(factory.CreateLink("News 3 (US)", "https://example.com/news3us"))
	news3.AddItem(factory.CreateLink("News 3 (Japan)", "https://example.com/news3jp"))
	newsTray := factory.CreateTray("News Site")
	newsTray.AddItem(news1)
	newsTray.AddItem(news2)
	newsTray.AddItem(news3)

	page := factory.CreatePage("Blog and News", "Alice")
	page.AddItem(blogTray)
	page.AddItem(newsTray)

	content, err := GeneratePageContent(page)
	assert.NoError(t, err)

	expected := `<!DOCTYPE html>
<html><head><title>Blog and News</title></head>
<body>
<h1>Blog and News</h1>
<ul>
<li>
Blog Site
<ul>
  <li><a href="https://example.com/blog1">Blog 1</a></li>
  <li><a href="https://example.com/blog2">Blog 2</a></li>
  <li><a href="https://example.com/blog3">Blog 3</a></li>
</ul>
</li>
<li>
News Site
<ul>
  <li><a href="https://example.com/news1">News 1</a></li>
  <li><a href="https://example.com/news2">News 2</a></li>
<li>
News 3
<ul>
  <li><a href="https://example.com/news3us">News 3 (US)</a></li>
  <li><a href="https://example.com/news3jp">News 3 (Japan)</a></li>
</ul>
</li>
</ul>
</li>
</ul>
<hr><address>Alice</address>
</body></html>
`

	assert.Equal(t, expected, content)
}
