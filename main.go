package main

import "fmt"

func main() {
	rows := 4 // Количество строк в пирамиде

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
