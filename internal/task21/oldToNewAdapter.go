package task21

import "fmt"

// OldToNewAdapter - адаптер, который преобразует OldPrinter к NewPrinter
type OldToNewAdapter struct {
	oldPrinter *OldPrinter
}

// PrintNew использует OldPrinter для печати в новом формате
func (a *OldToNewAdapter) PrintNew() {
	fmt.Println("New Format:", a.oldPrinter.message)
}

// NewOldToNewAdapter создает новый экземпляр OldToNewAdapter
func NewOldToNewAdapter(oldPrinter *OldPrinter) *OldToNewAdapter {
	return &OldToNewAdapter{oldPrinter: oldPrinter}
}
