package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	url := ""
	arch := ""

	// parse arguments
	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]
		reg := regexp.MustCompile(`([\d\D][^=]+)=([\d\D]+)$`)
		kvArray := reg.FindStringSubmatch(strings.TrimSpace(arg))
		fmt.Println(kvArray)
		if len(kvArray) > 2 {
			key := strings.TrimSpace(kvArray[1])
			value := strings.TrimSpace(kvArray[2])
			switch strings.TrimSpace(key) {
			case "-u", "u", "--url", "url":
				{
					url = strings.TrimSpace(value)
					break
				}
			case "-a", "a", "--arch", "arch":
				{
					arch = strings.TrimSpace(value)
					break
				}
			}
		}
	}

	if len(url) <= 0 {
		fmt.Printf("url must be set!\n")
		return
	}
	if len(arch) <= 0 {
		arch = "archive"
	}

	archZip := fmt.Sprintf("%s.zip", arch)

	err := Compress(arch, archZip)
	if err != nil {
		fmt.Println(err)
		return
	}

	md5, err := ComputeMd5(archZip)
	if err != nil {
		fmt.Printf("Err: %v", err)
		return
	}

	key := os.Getenv("PACKAGER_KEY")

	err = Upload(archZip, md5, key, url)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

}
