328. Odd Even Linked List
Medium
8.8K
476
Companies
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

 

Example 1:


Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]

Python Submission:
   class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:

        if not head: return head                                
        n1, n2, head2 = head, head.next, head.next            
                                                               
        while n2 and n2.next:                                  
            n1.next, n2.next = n1.next.next, n2.next.next      
            n1, n2 = n1.next, n2.next                          
            
        n1.next = head2                                         
        return head