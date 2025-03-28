package visitor_test

import (
	"testing"

	. "gof-in-go/behavioral/visitor"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryStructure(t *testing.T) {
	visitor := &ListVisitor{}

	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)
	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))

	result1 := rootDir.Accept(visitor)

	expected1 := `/root (30000)
/root/bin (30000)
/root/bin/vi (10000)
/root/bin/latex (20000)
/root/tmp (0)
/root/usr (0)`

	assert.Equal(t, expected1, result1)

	yukiDir := NewDirectory("yuki")
	hanakoDir := NewDirectory("hanako")
	tomuraDir := NewDirectory("tomura")
	usrDir.Add(yukiDir)
	usrDir.Add(hanakoDir)
	usrDir.Add(tomuraDir)
	yukiDir.Add(NewFile("diary.html", 100))
	yukiDir.Add(NewFile("Composite.java", 200))
	hanakoDir.Add(NewFile("memo.tex", 300))
	tomuraDir.Add(NewFile("game.doc", 400))
	tomuraDir.Add(NewFile("junk.mail", 500))

	result2 := rootDir.Accept(visitor)

	expected2 := `/root (31500)
/root/bin (30000)
/root/bin/vi (10000)
/root/bin/latex (20000)
/root/tmp (0)
/root/usr (1500)
/root/usr/yuki (300)
/root/usr/yuki/diary.html (100)
/root/usr/yuki/Composite.java (200)
/root/usr/hanako (300)
/root/usr/hanako/memo.tex (300)
/root/usr/tomura (900)
/root/usr/tomura/game.doc (400)
/root/usr/tomura/junk.mail (500)`

	assert.Equal(t, expected2, result2)
}
