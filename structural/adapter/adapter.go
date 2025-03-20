package adapter

type Banner struct {
	text string
}

func NewBanner(text string) *Banner {
	return &Banner{text: text}
}

func (b *Banner) encloseInParen() string {
	return "(" + b.text + ")"
}

func (b *Banner) encloseInAster() string {
	return "*" + b.text + "*"
}

type Wrap interface {
	WrapWithWeak() string
	WrapWithStrong() string
}

// WrapBannerがWrapインターフェースを満たすことを確認
var _ Wrap = (*WrapBanner)(nil)

type WrapBanner struct {
	banner *Banner
}

func NewWrapBanner(text string) *WrapBanner {
	return &WrapBanner{banner: NewBanner(text)}
}

func (pb *WrapBanner) WrapWithWeak() string {
	return pb.banner.encloseInParen()
}

func (pb *WrapBanner) WrapWithStrong() string {
	return pb.banner.encloseInAster()
}
