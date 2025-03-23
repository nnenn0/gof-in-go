package singleton_test

import (
	"gof-in-go/creational/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	assert.Equal(t, obj1, obj2, "obj1とobj2は同じインスタンスではありません")
}
