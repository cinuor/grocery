package binarysearch

// 基础二分查找
func BasicBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 明白搜索边界的概念，[left, right)
// 此时的搜索边界是：左闭右开，那么mid一定比平常情况下往右一位
// 这是再让right等于mid，让mid继续往左推进，知道左边第一个出现的中值

// 这种情况对应的是：1，2， 3，3，3，3，3，3， 5
// 当target = 3时， 找到的一定时最左边那个3
func LeftBoundaryBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 关键的操作
			// 找到目标之后，并不直接停止搜索，而是继续缩小搜索边界的right
			// 这样，知道找到最左边的那个target时，
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left >= len(nums) {
		return -1
	}

	return left
}

// 当我们明白了“搜索区间”之后，现在用“左闭右闭”的搜索区间，同样找到最左边的target
// 关键就是当第一次找到target之后，不要直接返回，而是继续把mid往左边推进
func LeftBoundaryBinarySearch2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right { // 退出条件时left > right，更具体一点的就是left = right + 1
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不直接返回, 而是再把mid往做推进
			right = mid - 1
		}
	}

	// 判断left的边界
	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}
