1318. Minimum Flips to Make a OR b Equal to c
Medium
1.8K
93
Companies
Given 3 positives numbers a, b and c. Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

 

Example 1:



Input: a = 2, b = 6, c = 5
Output: 3
Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)

Python Submission:
    class Solution:
    def minFlips(self, a: int, b: int, c: int) -> int:
        a,b,c = format(a, '032b'),format(b, '032b'),format(c, '032b')

        require_flip = 0
        for i in range(len(a)):
            if a[i]==b[i]:
                if a[i]=='1' and c[i]=='0':
                    require_flip+=2
                elif a[i]=='0' and c[i]=='1':
                    require_flip+=1
            else:
                if c[i]=='0':
                    require_flip+=1
        return require_flip


1448. Count Good Nodes in Binary Tree
Medium
5.1K
133
Companies
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

 

Example 1:
Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.


Python Submission:
    def test(node, path_max):


        include_self = False
        if node.val>=path_max:
            include_self = True
            path_max = node.val

        left_good = test(node.left, path_max) if node.left else 0
        right_good = test(node.right, path_max) if node.right else 0

        if include_self:
            return left_good+right_good+1
        return left_good+right_good


    class Solution:



        def goodNodes(self, root: TreeNode) -> int:
            return test(root,root.val)


199. Binary Tree Right Side View
Medium
10.6K
657
Companies
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Python Submission:

    def go(node, level, l):

        if level<len(l) and l[level]:
            l[level].append(node.val)
        else:
            l.append([node.val])

        
        if node.right:
            go(node.right, level+1, l)
        if node.left:
            go(node.left, level+1, l)

    class Solution:
        def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
            l = []
            if root:
                go(root, 0, l)
            ret = []
            for _level in l:
                ret.append(_level[0])
            return ret

62. Unique Paths
Medium
14.5K
397
Companies
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

 

Example 1:


Input: m = 3, n = 7
Output: 28

Python Submission:
    import numpy as np

class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        path_map = np.zeros((m,n))

        for i in range(m):


            for j in range(n):


                if i==0 or j ==0:
                    path_map[i][j]=1
                else:
                    path_map[i][j]=path_map[i-1][j]+path_map[i][j-1]
        
        return int(path_map[m-1][n-1])




198. House Robber
Medium
18.6K
345
Companies
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

 

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Python Submission:

    class Solution:
        def rob(self, nums: List[int]) -> int:

            if len(nums)<3:
                if len(nums)==2:
                    return max(nums[0],nums[1])
                if len(nums)==1:
                    return nums[0]
                if len(nums)==0:
                    return 0

            d = [nums[0], nums[1], nums[0]+nums[2]]


            for j in range(3,len(nums)):
                d.append(max(d[j-2]+nums[j],d[j-3]+nums[j])) 


            return max(d[-1],d[-2])