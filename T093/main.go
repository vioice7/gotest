package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 10, 11, 12}
	find := binarySearch(slice, 7, 0, len(slice)-1)
	fmt.Println("The find is (rec):", find)
	find = iterBinarySearch(slice, 7, 0, len(slice)-1)
	fmt.Println("The find is (iter):", find)
}

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex <= endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
