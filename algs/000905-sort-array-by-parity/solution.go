package sortarraybyparity

/*
	Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
	Return any array that satisfies this condition.


	Example 1:

	Input: nums = [3,1,2,4]
	Output: [2,4,3,1]
	Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

	Example 2:

	Input: nums = [0]
	Output: [0]
*/

func Solution(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]%2 == 0 {
			left++

			continue
		}

		if nums[right]%2 == 1 {
			right--

			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return nums
}
