package datafile

import (
	"fmt"
	"io/ioutil"
)

func WriteF(filename string, data string) {
	databyte := []byte(data)
	err := ioutil.WriteFile(filename, databyte, 0o644)
	if err != nil {
		fmt.Println(err)
	}
}
