package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	// var data []byte
	// data := []byte{}
	// data := make([]byte, 99999)
	// bytesRead, err := resp.Body.Read(data)
	// fmt.Println(resp)
	// fmt.Println(bytesRead, err)
	// fmt.Println(string(data))

	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
