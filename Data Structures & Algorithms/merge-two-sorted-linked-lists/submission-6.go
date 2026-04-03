/**
 * Definition for singly-linked list.
  * type ListNode struct {
   *     Val int
    *     Next *ListNode
	 * }
	  */

	  func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	      if list1 == nil {return list2}
		  	if list2 == nil {return list1}

				var newHead *ListNode
					if list1.Val < list2.Val {
							tmp := list1.Next
									list1.Next = mergeTwoLists(tmp, list2)
											
													return list1
														} else {
																tmp := list2.Next
																		list2.Next = mergeTwoLists(list1, tmp) 
																				
																						return list2
																							}

																								return newHead
																								}
																								