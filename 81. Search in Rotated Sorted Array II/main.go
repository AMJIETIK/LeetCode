package main

import(
	"fmt"
	"sort"
)

func main(){
	var list = []int{4,6,8,2,5,8,9,1}
	fmt.Println(search(list, 4))
}

func search(nums []int, target int) bool {
    sort.Ints(nums)

    left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // индекс среднего элемента

		if nums[mid] == target {
			return true // элемент найден
		}
		if nums[mid] < target {
			left = mid + 1 // ищем в правой части
		} else {
			right = mid - 1 // ищем в левой части
		}
	}
    return false
}