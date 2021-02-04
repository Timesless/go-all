package main

import (
	"fmt"
	"math"
)

/*
   查找最大的平均值
*/
func main() {
	s := []int{-1, -3, -9, 10, 15, 7}
	fmt.Println(findMaxAverage(s, 2))
}

func findMaxAverage(nums []int, k int) float64 {

	sz := len(nums)
	if sz < 2 {
		return float64(nums[0])
	}
	fk, ln := float64(k), sz-k+1
	f := func(prev, curr int) int {
		return prev + curr
	}
	max := float64(math.MinInt64)
	tmpTotal := Reduce(nums[:k], f)
	tmpAvg := float64(tmpTotal) / fk
	if tmpAvg > max {
		max = tmpAvg
	}
	for i := 1; i < ln; i++ {
		tmpTotal = tmpTotal - nums[i-1] + nums[i+k-1]
		tmpAvg = float64(tmpTotal) / fk
		if tmpAvg > max {
			max = tmpAvg
		}
	}
	return max
}

/*
	自定义reduce函数
*/
func Reduce(nums []int, f func(prev, curr int) int) int {
	rs := 0
	for i, e := range nums {
		if i == 0 {
			rs = f(0, e)
		} else {
			rs = f(rs, e)
		}
	}
	return rs
}
