package mediator

type Colleague interface {
	SetMediator(mediator Mediator)
	SetColleagueEnabled(enabled bool)
}

type Mediator interface {
	CreateColleagues()
	ColleagueChanged()

	GetCheckGuest() *ColleagueCheckbox
	GetCheckLogin() *ColleagueCheckbox
	GetTextUser() *ColleagueTextField
	GetTextPass() *ColleagueTextField
	GetButtonOk() *ColleagueButton
	GetButtonCancel() *ColleagueButton
}

// ColleagueCheckboxがColleagueインターフェースを満たしていることを確認
var _ Colleague = (*ColleagueCheckbox)(nil)

type ColleagueCheckbox struct {
	caption  string
	checked  bool
	mediator Mediator
}

func NewColleagueCheckbox(caption string, checked bool) *ColleagueCheckbox {
	return &ColleagueCheckbox{
		caption: caption,
		checked: checked,
	}
}

func (c *ColleagueCheckbox) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueCheckbox) SetColleagueEnabled(enabled bool) {
	if c.checked != enabled {
		c.checked = enabled
		if c.mediator != nil {
			c.mediator.ColleagueChanged()
		}
	}
}

func (c *ColleagueCheckbox) IsChecked() bool {
	return c.checked
}

// ColleagueTextFieldがColleagueインターフェースを満たしていることを確認
var _ Colleague = (*ColleagueTextField)(nil)

type ColleagueTextField struct {
	text     string
	enabled  bool
	mediator Mediator
}

func NewColleagueTextField(text string) *ColleagueTextField {
	return &ColleagueTextField{
		text:    text,
		enabled: true,
	}
}

func (t *ColleagueTextField) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

func (t *ColleagueTextField) SetColleagueEnabled(enabled bool) {
	t.enabled = enabled
}

func (t *ColleagueTextField) SetText(text string) {
	t.text = text
	if t.mediator != nil {
		t.mediator.ColleagueChanged()
	}
}

func (t *ColleagueTextField) GetText() string {
	return t.text
}

func (t *ColleagueTextField) IsEnabled() bool {
	return t.enabled
}

// ColleagueButtonがColleagueインターフェースを満たしていることを確認
var _ Colleague = (*ColleagueButton)(nil)

type ColleagueButton struct {
	caption  string
	enabled  bool
	mediator Mediator
}

func NewColleagueButton(caption string) *ColleagueButton {
	return &ColleagueButton{
		caption: caption,
		enabled: true,
	}
}

func (b *ColleagueButton) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *ColleagueButton) SetColleagueEnabled(enabled bool) {
	b.enabled = enabled
}

func (b *ColleagueButton) IsEnabled() bool {
	return b.enabled
}

// LoginFrameがMediatorインターフェースを満たしていることを確認
var _ Mediator = (*LoginFrame)(nil)

type LoginFrame struct {
	checkGuest   *ColleagueCheckbox
	checkLogin   *ColleagueCheckbox
	textUser     *ColleagueTextField
	textPass     *ColleagueTextField
	buttonOk     *ColleagueButton
	buttonCancel *ColleagueButton
}

func NewLoginFrame() *LoginFrame {
	lf := &LoginFrame{}
	lf.CreateColleagues()
	return lf
}

func (lf *LoginFrame) CreateColleagues() {
	lf.checkGuest = NewColleagueCheckbox("Guest", true)
	lf.checkLogin = NewColleagueCheckbox("Login", false)

	lf.textUser = NewColleagueTextField("")
	lf.textPass = NewColleagueTextField("")

	lf.buttonOk = NewColleagueButton("OK")
	lf.buttonCancel = NewColleagueButton("Cancel")

	lf.checkGuest.SetMediator(lf)
	lf.checkLogin.SetMediator(lf)
	lf.textUser.SetMediator(lf)
	lf.textPass.SetMediator(lf)
	lf.buttonOk.SetMediator(lf)
	lf.buttonCancel.SetMediator(lf)

	lf.ColleagueChanged()
}

func (lf *LoginFrame) GetCheckGuest() *ColleagueCheckbox {
	return lf.checkGuest
}

func (lf *LoginFrame) GetCheckLogin() *ColleagueCheckbox {
	return lf.checkLogin
}

func (lf *LoginFrame) GetTextUser() *ColleagueTextField {
	return lf.textUser
}

func (lf *LoginFrame) GetTextPass() *ColleagueTextField {
	return lf.textPass
}

func (lf *LoginFrame) GetButtonOk() *ColleagueButton {
	return lf.buttonOk
}

func (lf *LoginFrame) GetButtonCancel() *ColleagueButton {
	return lf.buttonCancel
}

func (lf *LoginFrame) ColleagueChanged() {
	if lf.checkGuest.IsChecked() {
		lf.checkLogin.SetColleagueEnabled(false)
		lf.textUser.SetColleagueEnabled(false)
		lf.textPass.SetColleagueEnabled(false)
		lf.buttonOk.SetColleagueEnabled(true)
	} else {
		lf.checkLogin.SetColleagueEnabled(true) // Ensure Login checkbox reflects the state
		lf.textUser.SetColleagueEnabled(true)
		lf.userpassChanged()
	}
}

func (lf *LoginFrame) userpassChanged() {
	if len(lf.textUser.GetText()) > 0 {
		lf.textPass.SetColleagueEnabled(true)
		if len(lf.textPass.GetText()) > 0 {
			lf.buttonOk.SetColleagueEnabled(true)
		} else {
			lf.buttonOk.SetColleagueEnabled(false)
		}
	} else {
		lf.textPass.SetColleagueEnabled(false)
		lf.buttonOk.SetColleagueEnabled(false)
	}
}
