package main

import "fmt"

func binarysearch(list []int, key int) int {
	start := 0
	end := len(list) - 1
	mid := 0
	for start <= end {
		mid = (start + end + 1) / 2 // 取前半个 (start + end) / 2 取后半个
		if list[mid] > key {
			end = mid - 1
		} else if list[mid] < key {
			start = mid + 1
		} else {
			return mid
		}
	}
	return mid
}

func main() {
	fmt.Printf("hello, world\n")
	a := []int{1, 3, 5, 7, 9}
	fmt.Println(binarysearch(a, 1))
	fmt.Println(binarysearch(a, 2))
	fmt.Println(binarysearch(a, 8))
	fmt.Println(binarysearch(a, 9))
}
