package main

import "fmt"

func main() {
	ind := binarySearch([]int{1, 4, 6, 7, 8, 9, 11}, 7)
	fmt.Println(ind)
}

func binarySearch(nums []int, elem int) int {
	low, high := 0, len(nums)

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		if guess == elem {
			return mid
		} else if guess < elem {
			low = mid + 1
		} else if guess > elem {
			high = mid - 1
		}
	}

	return -1
}
