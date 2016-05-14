package main

import (
	"fmt"
)

func main() {

	if err := Compress(`md5.go`, `md5.go.zip`); err != nil {
		fmt.Println(err)
	}

	md5, err := ComputeMd5("md5.go.zip")
	if err != nil {
		fmt.Printf("Err: %v", err)
	} else {
		fmt.Printf("md5.go.zip md5 checksum is: %s\n", md5)
	}

	err = Upload("md5.go.zip", md5, "http://127.0.0.1")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
