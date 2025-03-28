package mediator_test

import (
	"testing"

	. "gof-in-go/behavioral/mediator"

	"github.com/stretchr/testify/assert"
)

func TestLoginFrameLogic(t *testing.T) {
	lf := NewLoginFrame()

	assert.True(t, lf.GetCheckGuest().IsChecked())
	assert.False(t, lf.GetTextUser().IsEnabled())
	assert.False(t, lf.GetTextPass().IsEnabled())
	assert.True(t, lf.GetButtonOk().IsEnabled())

	lf.GetCheckGuest().SetColleagueEnabled(false)
	assert.False(t, lf.GetCheckGuest().IsChecked())
	assert.True(t, lf.GetCheckLogin().IsChecked())
	assert.True(t, lf.GetTextUser().IsEnabled())
	assert.False(t, lf.GetTextPass().IsEnabled())
	assert.False(t, lf.GetButtonOk().IsEnabled())

	lf.GetTextUser().SetText("testuser")
	assert.True(t, lf.GetTextUser().IsEnabled())
	assert.True(t, lf.GetTextPass().IsEnabled())
	assert.False(t, lf.GetButtonOk().IsEnabled())

	lf.GetTextPass().SetText("testpass")
	assert.True(t, lf.GetButtonOk().IsEnabled())

	lf.GetTextUser().SetText("")
	assert.False(t, lf.GetTextPass().IsEnabled())
	assert.False(t, lf.GetButtonOk().IsEnabled())
}
