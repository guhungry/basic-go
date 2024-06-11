package money

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatMoney(value float64) string {
	p := message.NewPrinter(language.Thai)
	return p.Sprintf("%0.2f\n", value)
}
