package main

import (
	"fmt"
	"os"

	"github.com/anurag0608/randnumsplash"
)

func checkAndGenRandFile() {
	_, err := os.Stat("rand.txt")
	if err != nil {
		if err := randnumsplash.GenerateRandFile(1024*1024*20, "", "rand.txt", true); err != nil {
			panic(err)
		}
	}
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered: %w", r)
		}
	}()
	checkAndGenRandFile()
	chunkSize := 1024 * 1024 * 5 // 5Mb
	file, err := os.Open("rand.txt")
	if err != nil {
		panic(err)
	}
	buff := make([]byte, chunkSize)
	itr := 0
	os.Mkdir("output", 0755)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			fmt.Println(err)
			break
		}
		itr += 1
		newFileName := fmt.Sprintf("output/rand%v.txt", itr)
		outFile, err := os.OpenFile(newFileName, os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
			break
		}
		outFile.Write(buff[:bytesRead])
	}
}
