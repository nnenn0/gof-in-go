package proxy

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(s string) string
}

// PrinterがPrintableインターフェースを満たしていることを確認
var _ Printable = (*Printer)(nil)

type Printer struct {
	name string
}

func NewPrinter() *Printer {
	p := &Printer{}
	p.heavyJob("Printerのインスタンスを生成中")
	return p
}

func NewPrinterWithName(name string) *Printer {
	p := &Printer{name: name}
	p.heavyJob("Printerのインスタンス(" + name + ")を生成中")
	return p
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) Print(s string) string {
	return "=== " + p.name + " ===\n" + s
}

func (p *Printer) heavyJob(msg string) string {
	result := msg
	for range 5 {
		result += "."
	}
	result += "完了。\n"
	return result
}

// PrinterProxyがPrintableインターフェースを満たしていることを確認
var _ Printable = (*PrinterProxy)(nil)

type PrinterProxy struct {
	name string
	real *Printer
}

func NewPrinterProxy() *PrinterProxy {
	return &PrinterProxy{name: "No Name"}
}

func NewPrinterProxyWithName(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

func (p *PrinterProxy) SetPrinterName(name string) {
	p.name = name
	if p.real != nil {
		p.real.SetPrinterName(name)
	}
}

func (p *PrinterProxy) GetPrinterName() string {
	return p.name
}

func (p *PrinterProxy) Print(s string) string {
	p.realize()
	return p.real.Print(s)
}

func (p *PrinterProxy) realize() {
	if p.real == nil {
		p.real = NewPrinterWithName(p.name)
	}
}
