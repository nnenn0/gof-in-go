package state_test

import (
	"testing"

	. "gof-in-go/behavioral/state"

	"github.com/stretchr/testify/assert"
)

func TestMainFunctionality(t *testing.T) {
	safeFrame := NewSafeController("State Sample Test")

	for hour := range 24 {
		safeFrame.SetClock(hour)

		if hour == 9 {
			assert.Equal(t, "[昼間]", safeFrame.GetCurrentState().String(), "9時で昼間の状態に変化すること")
		} else if hour == 17 {
			assert.Equal(t, "[夜間]", safeFrame.GetCurrentState().String(), "17時で夜間の状態に変化すること")
		}

		if hour == 8 {
			safeFrame.Use()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "非常：夜間の金庫使用！", "8時に夜間の金庫使用の呼び出しが行われること")
			safeFrame.Alarm()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "非常ベル(夜間)", "8時に非常ベルの呼び出しが行われること")
			safeFrame.Phone()
			assert.Contains(t, safeFrame.GetLogRecords(), "夜間の通話録音", "8時に夜間の通話録音のログが出力されること")
		} else if hour == 10 {
			safeFrame.Use()
			assert.Contains(t, safeFrame.GetLogRecords(), "金庫使用(昼間)", "10時に金庫使用のログが出力されること")
			safeFrame.Alarm()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "非常ベル(昼間)", "10時に非常ベルの呼び出しが行われること")
			safeFrame.Phone()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "通常の通話(昼間)", "10時に通常の通話の呼び出しが行われること")
		} else if hour == 20 {
			safeFrame.Use()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "非常：夜間の金庫使用！", "20時に夜間の金庫使用の呼び出しが行われること")
			safeFrame.Alarm()
			assert.Contains(t, safeFrame.GetSecurityCenterCalls(), "非常ベル(夜間)", "20時に非常ベルの呼び出しが行われること")
			safeFrame.Phone()
			assert.Contains(t, safeFrame.GetLogRecords(), "夜間の通話録音", "20時に夜間の通話録音のログが出力されること")
		}
	}
}
