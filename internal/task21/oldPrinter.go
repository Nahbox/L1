package task21

import "fmt"

// OldPrinter - структура, которая печатает данные в старом формате
type OldPrinter struct {
	message string
}

// PrintOld печатает данные в старом формате
func (op *OldPrinter) PrintOld() {
	fmt.Println("Old Format:", op.message)
}
