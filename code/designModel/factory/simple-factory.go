package main

import "fmt"

// Printer 简单工厂要返回的接口类型
type Printer interface {
	Print(name string) string
}

// CnPrinter 是 Printer 接口的实现，它说中文
type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好, %s", name)
}

// EnPrinter 是 Printer 接口的实现，它说中文
type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

type PayMethod interface {
	Pay(method string, amount float64) error
}

type WxPay struct {
}

func (*WxPay) Pay(method string, amount float64) error {
	fmt.Println(method, amount)
	return nil
}

type AliPay struct {
}

func (*AliPay) Pay(method string, amount float64) error {
	fmt.Println(method, amount)
	return nil
}

func NewPayMethod(method string) PayMethod {
	switch method {
	case "wxPay":
		return new(WxPay)
	case "aliPay":
		return new(AliPay)
	default:
		return new(AliPay)
	}
}

func main() {
	cnPrinter := NewPrinter("cn")
	fmt.Println(cnPrinter.Print("小鱼"))
	enPrinter := NewPrinter("en")
	fmt.Println(enPrinter.Print("xiao yu"))

	wxPay := NewPayMethod("wxPay")
	fmt.Println(wxPay.Pay("wx", 12))
	alipay := NewPayMethod("aliPay")
	fmt.Println(alipay.Pay("ali", 123))
}
