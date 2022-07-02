package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//*os.file has a Readfunction, thus implementing the interface of Reader(io.Copy)
	io.Copy(os.Stdout, f)
	//io.Copy(dst, res)
}
