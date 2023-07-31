package main

func main() {

}

func peakIndexInMountainArrayV1(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return i - 1
		}
	}

	return 0
}

// O(log(N)
func peakIndexInMountainArrayV2(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		mid := (l + r) / 2

		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
