// functions with defer keyword
package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("dealers.txt")
	defer closeFile(file)
	writeFile(file, "A1 Auto")
}

func createFile(path string) *os.File {
	fmt.Println("creating file")
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File, dealerName string) {
	fmt.Println("writing to file")
	fmt.Fprintln(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("closing file")
	file.Close()
}