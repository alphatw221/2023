437. Path Sum III
Medium
9.9K
471
Companies
Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

 

Example 1:


Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.


Python Submission:

    class Solution:

        paths = 0

        def test(self, node, nums, target):

            _nums = []
            for num in nums:
                _nums.append(num-node.val)
                if num == node.val:
                    self.paths+=1
            _nums.append(target)


            if node.left:
                self.test(node.left, _nums, target)
            
            if node.right:
                self.test(node.right, _nums, target)


        def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:

            
            if not root:
                return 0

            self.test(root,[targetSum], targetSum)

            return self.paths


        

547. Number of Provinces
Medium
8.7K
319
Companies
There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

 

Example 1:


Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Python Submission:
    class Solution:
        def findCircleNum(self, isConnected: List[List[int]]) -> int:

            def go(x, done):

                for k in range(len(isConnected[x])):
                    
                    if not done[k] and isConnected[x][k]==1:
                        done[k]=True
                        go(k,done)


            x= len(isConnected)

            done = [False]*x
            provinces = 0
            for i in range(x):
                    if not done[i]:
                        go(i, done)
                        provinces+=1
            return provinces





2336. Smallest Number in Infinite Set
Medium
1.4K
144
Companies
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

Implement the SmallestInfiniteSet class:

SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
int popSmallest() Removes and returns the smallest integer contained in the infinite set.
void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
 

Example 1:

Input
["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
[[], [2], [], [], [], [1], [], [], []]
Output
[null, null, 1, 2, 3, null, 1, 4, 5]

Explanation
SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
                                   // is the smallest number, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.




Python Submission:
    class SmallestInfiniteSet:

        def __init__(self):
            self.heap = []
            self.set = set({})
            self.min = 1

        def popSmallest(self) -> int:
            if self.heap:
                num = heapq.heappop(self.heap)
                self.set.remove(num)
                return num
            self.min+=1
            return self.min-1

        def addBack(self, num: int) -> None:
            
            if self.min> num and num not in self.set:
                self.set.add(num)
                heapq.heappush(self.heap, num)