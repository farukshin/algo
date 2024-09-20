# Бинарный поиск

Для поиска в отсортировоанном массве

```
package main

func binarySearch(arr []int, s int) int {
	var l = 0
	var r = len(arr) - 1
	for l <= r {
		var mid = (l + r) / 2
		if arr[mid] == s {
			return mid
		} else if arr[mid] < s {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	var arr = []int{1, 2, 5, 7, 8, 9, 16}
	println(binarySearch(arr, 5))
}

``