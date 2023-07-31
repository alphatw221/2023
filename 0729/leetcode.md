735. Asteroid Collision
Medium
6.7K
685
Companies
We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

 

Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Python Submission:
    class Solution:
        def asteroidCollision(self, asteroids: List[int]) -> List[int]:

            stack = []


            for asteroid in asteroids:

                if asteroid > 0:
                    stack.append(asteroid)
                else:
                    add = True

                    while stack and stack[-1]>0:

                        if stack[-1]+asteroid < 0:
                            stack.pop(-1)
                            add = True
                        elif stack[-1]+asteroid == 0:
                            stack.pop(-1)
                            add = False
                            break
                        else:
                            add = False
                            break

                    if add:
                        stack.append(asteroid)

            return stack


        


394. Decode String
Medium
11.4K
515
Companies
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.

 

Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"

Python Submission:
    class Solution:
        def decodeString(self, s: str) -> str:

            num_stack = []
            string_stack = []

            still_num = False

            for i in range(len(s)):


                c = s[i]

                if c.isnumeric():
                    if not still_num:
                        num_stack.append('|')
                    num_stack.append(c)
                    still_num = True

                elif c == ']':

                    inner = ''
                    while string_stack:
                        p = string_stack.pop(-1)
                        if p == '[':
                            break
                        inner = p + inner

                    num = ''
                    while num_stack:
                        p = num_stack.pop(-1)
                        if p == '|':
                            break
                        num = p + num
                    num = int(num)


                    string_stack.append(inner*num)
                else:
                    string_stack.append(c)
                    still_num = False
            
            return ''.join(string_stack)




2095. Delete the Middle Node of a Linked List
Medium
3.3K
61
Companies
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
 

Example 1:


Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node. 


Python Submission:
    class Solution:
        def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
            ptr = head

            l = [ptr]
            while ptr.next:
                ptr= ptr.next
                l.append(ptr)

            mid = math.floor(len(l)/2)

            if len(l) == 2:
                head.next = None
            elif len(l)==1:
                return None
            else:
                l[mid-1].next = l[mid+1]
            return head