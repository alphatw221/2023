208. Implement Trie (Prefix Tree)
Medium
10.6K
119
Companies
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
 

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

Python Submission:
    class Trie:

    def __init__(self):
        self.d ={}

    def insert(self, word: str) -> None:
        node = self.d
        for c in word:
            if c not in node:
                node[c] = {}
            node = node[c]
        node['+'] = True

    def search(self, word: str) -> bool:
        node = self.d
        for c in word:
            if c not in node:
                return False
            node = node[c]
        return node.get('+', False)

    def startsWith(self, prefix: str) -> bool:
        node = self.d
        for c in prefix:
            if c not in node:
                return False
            node = node[c]
        return True



    # Your Trie object will be instantiated and called as such:
    # obj = Trie()
    # obj.insert(word)
    # param_2 = obj.search(word)
    # param_3 = obj.startsWith(prefix)





435. Non-overlapping Intervals
Medium
7.3K
189
Companies
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

 

Example 1:

Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.


Python submission:
    import numpy as np
    class Solution:
        def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
            s_list = []
            for s,e in intervals:
                s_list.append(s)

            x = np.array(s_list)
            indexs = np.argsort(x)

            remove = 0
            current = intervals[indexs[0]]
            for index in indexs[1:]:
                next_interval = intervals[index]


                if current[1]<=next_interval[0]:
                    current = next_interval
                elif current[1]>=next_interval[1]:
                    current = next_interval
                    remove+=1
                else:
                    remove+=1
            
            return remove


739. Daily Temperatures
Medium
11.2K
246
Companies
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

 

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Python Submission:
    class Solution:
        def dailyTemperatures(self, temperatures: List[int]) -> List[int]:

            stack = []
            n = len(temperatures)
            res = [0] * n
            for i in range(n):
                t = temperatures[i]
                while stack != [] and temperatures[stack[-1]] < t:
                    less_index = stack.pop()
                    res[less_index] = i - less_index
                stack.append(i)
            return res