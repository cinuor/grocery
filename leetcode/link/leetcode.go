package link

type ListNode struct {
	Val  int
	Next *ListNode
}

// 经典的链表翻转之递归
func ReverseR(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		newHead := ReverseR(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}
}

// 经典的链表翻转之迭代
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	head.Next = nil
	return pre
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Val = carry % 10
		carry /= 10

	}

	return dummy.Next
}

// 19. 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//add sentinel
	s := &ListNode{}
	s.Next = head
	pre := s
	i := 1
	for head != nil {
		if i > n {
			pre = pre.Next
		}
		head = head.Next
		i++
	}

	pre.Next = pre.Next.Next
	return s.Next
}

// 21. 合并两个链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	if l1 != nil {
		prehead.Next = l1
	}

	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	cur, pre := head, head.Next
	cur.Next = swapPairs(pre.Next)
	pre.Next = cur
	return pre
}

// 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if k == 0 {
		return head
	}

	cur := head
	count := 1
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	cur.Next = head

	for i := 0; i < count-(k%count); i++ {
		cur = cur.Next
	}

	head = cur.Next
	cur.Next = nil
	return head
}

// 删除重复节点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	a := dummy
	b := head

	for b != nil && b.Next != nil {
		if a.Next.Val != b.Next.Val {
			a = a.Next
			b = b.Next
		} else {
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		}
	}
	return dummy.Next
}

// 23. 合并K个升序链表
func mergeKList(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[i+k])
		}
		n = k
	}

	return lists[0]

}

// 25. k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a := head
	b := head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a *ListNode, b *ListNode) *ListNode {
	pre := new(ListNode)
	cur := a

	for cur != b {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 判断链表是否有环
func checkCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	a, b := head, head
	for b != nil && b.Next != nil {
		a = a.Next
		b = b.Next.Next
		if a == b {
			return true
		}
	}
	return false

}
