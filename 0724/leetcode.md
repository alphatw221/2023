151. Reverse Words in a String

Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

 

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"

Python Submission:
    class Solution(object):
        def reverseWords(self, s):
            """
            :type s: str
            :rtype: str
            """
            ret = [w for w in s.split(' ') if w]
            ret.reverse()
            return ' '.join(ret)

Go Submission:
    func reverseWords(s string) string {

        str1 := strings.Split(s, " ")

        ret := ""

        for i := len(str1)-1; i >= 0; i-- {

            if str1[i] == "" {
                continue
            }

            ret += " " + str1[i] 
        }

        return ret[1:]


    }



11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.


Python Submission:
    class Solution:
        def maxArea(self, height: List[int]) -> int:

            l=0
            r = len(height)-1
            max_water=0
            while l<r:
                wide = r-l
                h = min(height[l],height[r])
                if max_water<wide*h:
                    max_water=wide*h
                if height[l]<=height[r]:
                    l+=1
                else:
                    r-=1
            return max_water



Go Submission:
    func maxArea(height []int) int {
        l:=0
        r:=len(height)-1

        max_water:=0
        for l<r{
            wide:=r-l
            h :=0

            if height[l] > height[r]{
                h=height[r]
            }else{
                h=height[l]
            }

            if max_water<wide*h{
                max_water=wide*h
            }
            if height[l]<=height[r]{
                l+=1
            }else{
                r-=1
            }
        }
        return max_water
    }







2390. Removing Stars From a String

You are given a string s, which contains stars *.

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.


Python Submission:
    

    class Solution:
        def removeStars(self, s: str) -> str:
            stack = []

            for c in s:
                if c!='*':
                    stack.append(c)
                else:
                    stack.pop()
            
            return ''.join(stack)
        


Go Submission:
    func removeStars(s string) string {

        var stack []byte

        for i, _ := range s {
            //Do something with index and character

            if s[i]!='*'{

                stack = append(stack, s[i])
            }else{
                stack = stack[:len(stack)-1]
            }
        }
        return string(stack)

    }