type Node struct {
	Val int
	Next *Node
}

func findDuplicate(nums []int) int {
    slow := nums[0]
	fast := nums[0]

	slow = nums[slow]
	fast = nums[nums[fast]]
	
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
		fmt.Println(slow, fast)
	}
	
	slow = nums[0] 
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		fmt.Println(slow, fast)
	}

	return slow
}
