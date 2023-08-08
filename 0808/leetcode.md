1679. Max Number of K-Sum Pairs
Medium
2.5K
55
Companies
You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.

 

Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.


Python Submission:
    class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:

        nums.sort()



        i = 0
        r = 0

        l = len(nums)
        j_cache = l-1
        while i<l:

            if i == j_cache:
                break



            for j in range(j_cache,i,-1):
                j_cache = j

                if nums[j]+nums[i]<k:
                    break
                elif nums[j]+nums[i]==k:

                    r+=1
                    j_cache-=1
                    break   

            i+=1

        return r