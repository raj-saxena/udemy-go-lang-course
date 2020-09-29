package main

import (
	"fmt"
	"io"
	"os"
)

type LameWriter struct{}

func (w LameWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main()  {
	name := os.Args[1]

	f, err := os.Open(name)
	if(err != nil) {
		fmt.Printf("Ah, crap! File %v doesn't exist", name)
		return
	}
	w := LameWriter{}

	io.Copy(w, f)
}