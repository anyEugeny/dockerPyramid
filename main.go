package main

import (
	"fmt"
	"os"
)

func main() {
	rows := 5 // Количество строк в пирамиде
	// Читайте аргументы командной строки
	rows = os.Args[1:]
	for i := 1; i <= rows; i++ {
		// Печать пробелов перед звездочками
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// Печать звездочек
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		// Переход на новую строку после каждой строки
		fmt.Println()
	}
}
