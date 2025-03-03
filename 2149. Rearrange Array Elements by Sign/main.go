package main

import (
	"fmt"
)

func main() {
	var nums = []int{3,1,-2,-5,2,-4}
	fmt.Println(rearrangeArray(nums))
}

func rearrangeArray(nums []int) []int {
    positive := make([]int,0,5)
    negative := make([]int,0,5)
    for _,num := range nums{
        if num > 0{
            positive = append(positive, num)
        }else{
            negative = append(negative, num)
        }
    }

var positiveI, negativeI int
    for i:=0; i<len(nums)-1;{
        nums[i] = positive[positiveI]
        i++
        positiveI++
        nums[i] = negative[negativeI]
        i++
        negativeI++
    }
    return nums
}