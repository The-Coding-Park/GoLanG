package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writefiledemo() {
	write1()
}
func write1() {
	b := []byte("how did you like the review?\n")
	ioutil.WriteFile("f.txt", b, 0644)
	fmt.Println("done")
}
func write() {
	f, _ := os.OpenFile("f.txt", os.O_APPEND, os.O_WRONLY, 0777)
	defer f.Close()
	f.Write(b)
	fmt.Println(f.Stat())
}
