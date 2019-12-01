package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestInitRead(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		log.Println(path)
		var filename string
		if runtime.GOOS == "windows" {
			filename = path + "\\s1.txt"
		} else {
			filename = path + "/s1.txt"
		}
		fmt.Println(filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
		}
		file.Close()
	}
}
