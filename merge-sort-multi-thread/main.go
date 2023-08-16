package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const N = 19999999

var wg sync.WaitGroup
var NUM_MAX_CPU int
var arr []int

func init() {
	// seed
	fmt.Println("Seeding rand lib...")
	rand.Seed(time.Now().UnixNano())
	// create an array of size n with random numbers
	arr = make([]int, N)
	fmt.Printf("Filling arr of size: %d\n", N)
	fill(arr)
	// get total num of available cpus
	NUM_MAX_CPU = runtime.NumCPU() + 2 // adding 2 extra
	fmt.Printf("Found %d num of available cpus\n", NUM_MAX_CPU)
}
func getRandomNumber() int {
	return rand.Intn(100)
}
func fill(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = getRandomNumber()
	}
}

// this will be handled by a single thread
func mergeSortParallel(arr []int, i int, j int, i_org int, j_org int, subArrChan chan []int) []int {
	if i < j {
		mid := i + (j-i)/2
		mergeSortParallel(arr, i, mid, i_org, j_org, subArrChan)
		mergeSortParallel(arr, mid+1, j, i_org, j_org, subArrChan)
		merge(arr, i, mid, j)
		if i == i_org && j == j_org {
			wg.Done()
			subArr := make([]int, j_org-i_org)
			copy(subArr, arr[i:j])
			subArrChan <- subArr
			return subArr
		}
	}
	return make([]int, 0)
}
func mergeSort(arr []int, i int, j int) {
	if i < j {
		mid := i + (j-i)/2
		mergeSort(arr, i, mid)
		mergeSort(arr, mid+1, j)
		merge(arr, i, mid, j)
	}
}
func merge(arr []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q

	l1 := make([]int, n1)
	l2 := make([]int, n2)

	for i := 0; i < n1; i++ {
		l1[i] = arr[p+i]
	}
	for i := 0; i < n2; i++ {
		l2[i] = arr[q+1+i]
	}
	var i, j, k int
	k = p
	i = 0
	j = 0
	for i < n1 && j < n2 {
		if l1[i] < l2[j] {
			arr[k] = l1[i]
			i++
		} else {
			arr[k] = l2[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = l1[i]
		k++
		i++
	}
	for j < n2 {
		arr[k] = l2[j]
		k++
		j++
	}
}
func main() {
	single_stat()
	parallel_stats()
}
func parallel_stats() {
	startTime := time.Now()
	num_cpu_local := int(NUM_MAX_CPU)
	chunk_size := int(N / num_cpu_local)
	start, end := 0, chunk_size
	fmt.Printf("num_cpu_local:%d,n_local:%d\n", num_cpu_local, chunk_size)
	wg.Add(num_cpu_local)
	subArrList := make([]([]int), num_cpu_local)
	subArrChan := make(chan []int)
	for {
		if end > N {
			break
		}
		go mergeSortParallel(arr, start, end-1, start, end-1, subArrChan)
		start = end
		end += chunk_size
	}
	wg.Wait()
	i := 0
	for {
		select {
		case subArr := <-subArrChan:
			subArrList[i] = subArr
			i += 1
		default:
			// fmt.Println(subArrList)
			endTime := time.Now()
			diff := endTime.Sub(startTime)
			fmt.Printf("Parallel - Time taken:%d ms\n", diff.Milliseconds())
			return
		}
	}
}
func single_stat() {
	// single thread
	startTime := time.Now()
	mergeSort(arr, 0, N-1)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("Single - Time taken:%d ms\n", diff.Milliseconds())
}
