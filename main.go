package main

import "fmt"

// find maximum subarray problem
// using divide and conquer algorithm

func main() {
	problem := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	answer := maxSubArray(problem)

	fmt.Println(answer)
}

// divide big problem to small problem until find base case

// define the base case - find until 1 element left
// max sum is in the 3 cases
// - left haft
// - right haft
// - middle cross

func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

// divide and conquer part
func findMaxSubArray(nums []int, left int, right int) int {
	fmt.Printf("L %v R %v\n", left, right)
	// base case
	if left == right {
		fmt.Println("..base case..", nums[left])
		return nums[left]
	}

	mid := (left + right) / 2
	fmt.Printf("start left... L %v M %v R %v\n", left, mid, right)

	leftMax := findMaxSubArray(nums, left, mid)
	fmt.Printf("start right... L %v M %v R %v\n", left, mid, right)

	rightMax := findMaxSubArray(nums, mid+1, right)
	crossMax := maxCrossing(nums, left, right, mid)

	return max(leftMax, rightMax, crossMax)

}

func maxCrossing(nums []int, left int, right int, mid int) int {

	// start with lowest possible number for int32
	leftSum := -1 << 31
	rightSum := -1 << 31

	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(values ...int) int {

	maxVal := values[0]

	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
