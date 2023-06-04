package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type printer struct {}

func main() {
	//Getting Command line argument thus reading the filename form command line
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File could not open with error")
		log.Fatal(err)
		os.Exit(1)
	}

	p := printer{}

	io.Copy(p, file)
	defer file.Close()

}

func (p printer) Write(bs []byte) (int,error){
	fmt.Println(string(bs))
	return len(bs), nil
}