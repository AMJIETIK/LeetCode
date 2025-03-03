package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Print(nums)
}

func sortColors(nums []int)  {
    list := map[int]int{
        0: 0,
        1: 0,
        2: 0,
    }

    for _,v := range nums{
        list[v]++
    }

    index := 0
    
    //appending 0
    for i:=0; i<list[0]; i++{
        nums[index] = 0
        index++
    }
    //appending 1
    for i:=0; i<list[1]; i++{
        nums[index] = 1
        index++
    }
    //appending 2
    for i:=0; i<list[2]; i++{
        nums[index] = 2
        index++
    }
}