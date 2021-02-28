package sort

func countingSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// find max number in nums
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	// generate slice c, length max
	c := make([]int, max+1, max+1)
	for i := 0; i < len(nums); i++ {
		c[nums[i]]++
	}

	// 累加个数，利用个数形成一个有序数组
	// 形成的这个数组，表示了nums中的数据a，只能被限定在c[a]对应的一个范围中
	// 例如原始数据nums，[2, 5, 3, 0, 2, 3, 0, 3]
	// 生成的计数列表c：[2, 2, 4, 7, 7, 8]
	// 那么对于nums中的3来说，所有3的位置都不应该超过7，既c[3]
	// 那么排序的时候，只要把3从第7号位开始安放，并且安放一个之后再把位置减1，表示之后的3的边界又缩窄了一位

	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 为什么逆序遍历呢
	// 因为如果顺序遍历的话，原始数组中的元素会被按照先边界，再往中间的形式安置，就造成了不稳定
	// 例如数组nums: [3a, 3b, 3c, 2, 1, 0], a, b, c表示不同的元素, 计数数组为c: [1, 1, 1, 3] --累加--> [1, 2, 3, 6]
	// 如果顺序遍历，那么第一个3a就会被放到第6-1位，也就是第五位[?, ?, ?, ?, ?, 3a], c = [1, 2, 3, 5]
	// 继续遍历，3b则被放置在第5-1位，既第四位[?, ?, ?, ?, 3b, 3a], c = [1, 2, 3, 4]
	// 这样就看得出来，**顺序遍历会破坏算法的稳定性，而逆序遍历刚好适合原始数组中，相同数据的排列顺序**
	temp := make([]int, len(nums), len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		index := c[nums[i]] - 1
		temp[index] = nums[i]
		c[nums[i]]--
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]

	}
}
