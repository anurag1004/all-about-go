package main

import (
	"os"
	"testing"
)

const filePath = "C:\\Users\\hhxgdoihgopxhop\\Desktop\\all-about-go\\buffered-io\\part_1\\sample.txt"

func TestStartStatBuffered(t *testing.T) {
	var n int = 1
	os.Remove("stat_test.log")
	statFile, err := os.OpenFile("stat_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer statFile.Close()
	for {
		if n > 200 {
			break
		}
		startStat(statFile, filePath, n, true)
		n++
	}
}
func TestStartStatNonBuffered(t *testing.T) {
	var iter int = 1
	os.Remove("stat_test.log")
	statFile, err := os.OpenFile("stat_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer statFile.Close()
	for {
		if iter > 20 {
			break
		}
		startStat(statFile, filePath, iter, false)
		iter++
	}
}
