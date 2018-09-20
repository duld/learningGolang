package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	outDir := os.Args[1]
	srcFile := os.Args[2]

	str := fmt.Sprint("Target Director to export files:\t", outDir, "\nSource file to read from: \t", srcFile)
	fmt.Println(str)
}
