package composite_test

import (
	"testing"

	. "gof-in-go/structural/composite"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryStructure(t *testing.T) {
	root := NewDirectory("root")
	bin := NewDirectory("bin")
	tmp := NewDirectory("tmp")
	usr := NewDirectory("usr")

	root.Add(bin)
	root.Add(tmp)
	root.Add(usr)

	bin.Add(NewFile("vi", 10000))
	bin.Add(NewFile("latex", 20000))

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	usr.Add(yuki)
	usr.Add(hanako)
	usr.Add(tomura)

	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.java", 200))
	hanako.Add(NewFile("memo.tex", 300))
	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))

	expectedOutput := `/root (31500)
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

	assert.Equal(t, expectedOutput, root.PrintList(""))
}
