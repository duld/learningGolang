package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// write some data using copy //
	nfc, err := os.Create("copy.txt")
	if err != nil {
		log.Fatal("There was a problem creating file copy.txt", err)
	}
	defer nfc.Close()

	myData := strings.NewReader("Some exciting data...")
	_, err = io.Copy(nfc, myData)

	// write some data using writefile
	nfw, err := os.Create("write.txt")
	if err != nil {
		log.Fatal("Ther was a problem creating write.txt", err)
	}
	defer nfw.Close()

	someMoreData := "Super exciting data"
	_, err = nfw.WriteString(someMoreData)
	if err != nil {
		fmt.Println(err)
	}
}
