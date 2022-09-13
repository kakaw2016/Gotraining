package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	g, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer g.Close()
	r := strings.NewReader("James1 Bond")
	io.Copy(f, r)
	bs, err := ioutil.ReadAll(g)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

}
