package task21

// Точка входа в Task21
func Task21() {
	oldPrinter := &OldPrinter{message: "Hello, world!"}

	// Используем адаптер для вызова нового метода
	newPrinter := NewOldToNewAdapter(oldPrinter)
	newPrinter.PrintNew()
}
