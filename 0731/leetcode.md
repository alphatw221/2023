790. Domino and Tromino Tiling
Medium
3K
943
Companies
You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.


Given an integer n, return the number of ways to tile an 2 x n board. Since the answer may be very large, return it modulo 109 + 7.

In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.

 

Example 1:


Input: n = 3
Output: 5
Explanation: The five different ways are show above.


Python Submission:
    class Solution:
        def numTilings(self,n:int)->int:                            # Example:  n = 5

            prev, curr, tri = 1, 1, 0                               #  n   prev  curr  tri 
                                                                    # ---  ----  ----  ----
            for i in range(1,n):                                    #   1     1     1     0
                                                                    #   2     1     2     1
                prev, curr, tri = curr, prev+curr+2*tri, prev+tri   #   3     2     5     2
                                                                    #   4     5    11     4
            return curr%1000000007                                  #   5    11    24     9
                                                                    #              / 
                                                                    #           return

Python Submission:


Python Submission:
