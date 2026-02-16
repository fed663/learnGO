package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//io.Copy(os.Stdout, file)
	//fmt.Println("ebat")
	fmt.Fprint(file, "Сегодня ")
	fmt.Fprintln(file, "хорошая погода")
}
