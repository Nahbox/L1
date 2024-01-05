package task13

import "fmt"

func Task13() {
	a := 1
	b := 2
	fmt.Printf("До 1-го способа: a=%d, b=%d	", a, b)

	// 1 способ
	a, b = b, a
	fmt.Printf("После 1-го способа: a=%d, b=%d\n", a, b)

	a = 1
	b = 2
	fmt.Printf("До 2-го способа: a=%d, b=%d	", a, b)

	// 2 способ
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("После 2-го способа: a=%d, b=%d\n", a, b)
}
