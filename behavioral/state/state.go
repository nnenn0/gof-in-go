package state

import "fmt"

type Context interface {
	SetClock(hour int) string
	ChangeState(state State) string
	CallSecurityCenter(msg string) string
	RecordLog(msg string) string
	GetCurrentState() State
	GetSecurityCenterCalls() []string
	GetLogRecords() []string
}

type State interface {
	DoClock(context Context, hour int) string
	DoUse(context Context) string
	DoAlarm(context Context) string
	DoPhone(context Context) string
	String() string
}

// DayStateがStateインターフェースを満たしていることを確認
var _ State = (*DayState)(nil)

type DayState struct{}

var dayState *DayState

func NewDayState() *DayState {
	if dayState == nil {
		dayState = &DayState{}
	}
	return dayState
}

func (d *DayState) DoClock(context Context, hour int) string {
	if hour < 9 || 17 <= hour {
		return context.ChangeState(NewNightState())
	}
	return ""
}

func (d *DayState) DoUse(context Context) string {
	return context.RecordLog("金庫使用(昼間)")
}

func (d *DayState) DoAlarm(context Context) string {
	return context.CallSecurityCenter("非常ベル(昼間)")
}

func (d *DayState) DoPhone(context Context) string {
	return context.CallSecurityCenter("通常の通話(昼間)")
}

func (d *DayState) String() string {
	return "[昼間]"
}

// NightStateがStateインターフェースを満たしていることを確認
var _ State = (*NightState)(nil)

type NightState struct{}

var nightState *NightState

func NewNightState() *NightState {
	if nightState == nil {
		nightState = &NightState{}
	}
	return nightState
}

func (n *NightState) DoClock(context Context, hour int) string {
	if 9 <= hour && hour < 17 {
		return context.ChangeState(NewDayState())
	}
	return ""
}

func (n *NightState) DoUse(context Context) string {
	return context.CallSecurityCenter("非常：夜間の金庫使用！")
}

func (n *NightState) DoAlarm(context Context) string {
	return context.CallSecurityCenter("非常ベル(夜間)")
}

func (n *NightState) DoPhone(context Context) string {
	return context.RecordLog("夜間の通話録音")
}

func (n *NightState) String() string {
	return "[夜間]"
}

type SafeController struct {
	clockString    string
	screenLog      string
	currentState   State
	securityCenter []string
	logRecords     []string
}

func NewSafeController(title string) *SafeController {
	return &SafeController{
		currentState: NewDayState(),
	}
}

func (s *SafeController) SetClock(hour int) string {
	clockString := fmt.Sprintf("現在時刻は%02d:00", hour)
	s.clockString = clockString
	s.currentState.DoClock(s, hour)
	return clockString
}

func (s *SafeController) ChangeState(state State) string {
	msg := fmt.Sprintf("%sから%sへ状態が変化しました。", s.currentState, state)
	s.screenLog += msg + "\n"
	s.currentState = state
	return msg
}

func (s *SafeController) CallSecurityCenter(msg string) string {
	callMsg := "call! " + msg
	s.screenLog += callMsg + "\n"
	s.securityCenter = append(s.securityCenter, msg)
	return callMsg
}

func (s *SafeController) RecordLog(msg string) string {
	recordMsg := "record ... " + msg
	s.screenLog += recordMsg + "\n"
	s.logRecords = append(s.logRecords, msg)
	return recordMsg
}

func (s *SafeController) GetCurrentState() State {
	return s.currentState
}

func (s *SafeController) GetSecurityCenterCalls() []string {
	return s.securityCenter
}

func (s *SafeController) GetLogRecords() []string {
	return s.logRecords
}

func (s *SafeController) Use() {
	s.currentState.DoUse(s)
}

func (s *SafeController) Alarm() {
	s.currentState.DoAlarm(s)
}

func (s *SafeController) Phone() {
	s.currentState.DoPhone(s)
}
