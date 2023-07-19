package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
lets read sample.txt by divinding it in "n" parts
How to run :-
go run main.go [n] [useBuffer]
n: No of parts/chunks
useBuffer: Boolean (true=> use Buffered technique)

Create a dummy file first.. Name it sample.txt
*/
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	args := os.Args
	var n int
	var useBuffer bool
	if len(args) > 1 {
		parts, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		if len(args) == 3 {
			if strings.Compare(args[2], "true") == 0 {
				useBuffer = true
			} else {
				// defaulting to false
				useBuffer = false
			}
		}
		n = parts
	} else {
		// defaulting to n=5
		n = 5
	}
	filePath := "C:\\Users\\hhxgdoihgopxhop\\Desktop\\all-about-go\\buffered-io\\part_1\\sample.txt"
	statFile, err := os.OpenFile("stat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer statFile.Close()

	startStat(statFile, filePath, n, useBuffer)
}
func startStat(statFile *os.File, filePath string, n int, useBuffer bool) {
	var fileSize int64 // not defined yet
	// get the size of target file
	if fileInfo, err := os.Stat(filePath); err != nil {
		panic(err)
	} else {
		fileSize = fileInfo.Size()
	}
	fmt.Printf("FileSize: %0.3fMb, No of parts to be read in: %d\n", convertBytesToMb(fileSize), n)
	var totalTime time.Duration
	var avgHeapUsage float64
	if useBuffer {
		fmt.Println("\n----------Using Buffer-------------")
		totalTime, avgHeapUsage = readFileBufferedStat(filePath, fileSize, n)
	} else {
		fmt.Println("\n----------Using Non Buffer------------")
		totalTime, avgHeapUsage = readFileNonBufferedStat(filePath)
	}
	statFile.WriteString(fmt.Sprintf("%v|%0.4f|%0.5f\n", n, totalTime.Seconds(), avgHeapUsage))

}
func readFileBufferedStat(filePath string, fileSize int64, n int) (time.Duration, float64) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(-1)
	}
	fileSizeInMB := convertBytesToMb(fileSize)
	buffSize := (int(math.Ceil(fileSizeInMB)) * 1024 * 1024) / n
	fmt.Printf("Buff Size: %d bytes|%0.3fKb|%0.3fMb\n", buffSize, convertBytesToKb(int64(buffSize)), convertBytesToMb(int64(buffSize)))
	buff := make([]byte, buffSize)
	fmt.Println("--------------------------------------")
	var counter int
	var totalTime time.Duration
	var totalHeapUsage float64
	for {
		start := time.Now()
		byteRead, err := file.Read(buff)
		if err != nil {
			fmt.Println("File read complete...")
			break
		}
		counter++
		fmt.Printf("Chunk-%v\n", counter)
		end := time.Now()
		totalTime += showByteRead(byteRead, start, end)
		totalHeapUsage += showHeapUsage()
		fmt.Println("--------------------------------------")
	}
	fmt.Printf("Total time taken(seconds): %0.3fs\n", totalTime.Seconds())
	fmt.Printf("Average Heap usage: %0.5fMb\n", totalHeapUsage/float64(counter))
	return totalTime, totalHeapUsage / float64(counter)
}
func readFileNonBufferedStat(filePath string) (time.Duration, float64) {
	start := time.Now()
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	totalTime := showByteRead(len(data), start, end)
	totalHeapUsage := showHeapUsage()
	return totalTime, totalHeapUsage
}
func convertBytesToKb(b int64) float64 {
	return float64(b) / 1024
}
func convertBytesToMb(b int64) float64 {
	return float64(b) / (1024 * 1024)
}
func showHeapUsage() float64 {
	fmt.Println("\n############################")
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Heap memory usage: %0.3f Mb\n", float64(memStats.HeapAlloc)/(1024*1024))
	fmt.Println("############################")
	return float64(memStats.HeapAlloc) / (1024 * 1024)
}
func showByteRead(byteRead int, start time.Time, end time.Time) time.Duration {
	elapsed := end.Sub(start)
	fmt.Printf("Bytes Read: %v\n", byteRead)
	fmt.Printf("Read:%0.2fKb, Took: %0.5fs\n", convertBytesToKb(int64(byteRead)), elapsed.Seconds())
	return elapsed
}
