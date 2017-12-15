package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	codeFilePath := os.Args[1]
	dat, err := ioutil.ReadFile(codeFilePath)
	if err != nil {
		fmt.Println("Failed to Read file")
		return
	}
	//fmt.Print(string(dat))
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(dat); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(str)
}
