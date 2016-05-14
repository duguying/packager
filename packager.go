package main

import (
	"fmt"
)

func main() {

	if err := Compress(`md5.go`, `md5.go.zip`); err != nil {
		fmt.Println(err)
	}

	if b, err := ComputeMd5("md5.go.zip"); err != nil {
		fmt.Printf("Err: %v", err)
	} else {
		fmt.Printf("md5.go.zip md5 checksum is: %s\n", b)
	}

}
