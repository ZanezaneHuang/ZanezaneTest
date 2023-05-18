package main

import "fmt"

//写快排

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	begin, end := left, right
	idx := nums[begin]
	for begin < end {
		for begin < end && nums[end] >= idx {
			end--
		}
		for begin < end && nums[begin] <= idx {
			begin++
		}
		if begin < end {
			nums[begin], nums[end] = nums[end], nums[begin]
		}
	}
	nums[begin], nums[left] = nums[left], nums[begin]
	QuickSort(nums, left, begin-1)
	QuickSort(nums, begin+1, right)
}
func main() {
	sums := []int{1, 3, 4, 5, 2}
	QuickSort(sums, 0, len(sums)-1)
	fmt.Println(sums)
}
