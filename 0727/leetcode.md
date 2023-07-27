841. Keys and Rooms
Medium
5.5K
244
Companies
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0. Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.

When you visit a room, you may find a set of distinct keys in it. Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.

Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i, return true if you can visit all the rooms, or false otherwise.

 

Example 1:

Input: rooms = [[1],[2],[3],[]]
Output: true
Explanation: 
We visit room 0 and pick up key 1.
We then visit room 1 and pick up key 2.
We then visit room 2 and pick up key 3.
We then visit room 3.
Since we were able to visit every room, we return true.

Python Submission:
    class Solution:
        def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
            open_rooms = set({0})
            new_keys = set(rooms[0])


            while new_keys:

                keys = new_keys
                new_keys = set({})

                for key in keys:
                    if key in open_rooms:
                        continue
                    open_rooms.add(key)

                    for room_key in rooms[key]:
                        if room_key in open_rooms:
                            continue
                        new_keys.add(room_key)
            return len(open_rooms)==len(rooms)



2300. Successful Pairs of Spells and Potions
Medium
2.2K
49
Companies
You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.

Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

 

Example 1:

Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
Explanation:
- 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
- 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
- 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
Thus, [4,0,3] is returned.


Python Submission:
    class Solution:
        def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
            potions.sort()

            ret = []
            for i in range(len(spells)) :
                
                left = 0
                right = len(potions)-1
                
                if spells[i] * potions[right] < success:
                    ret.append(0)
                elif spells[i] * potions[left] >= success:
                    ret.append( len(potions))
                else:
                    while left<right-1:
                        mid = int((left+right)/2)
                        if potions[mid] * spells[i] >= success:
                            right = mid
                        else:
                            left = mid
                
                    ret.append( len(potions)-right)
            return ret



215. Kth Largest Element in an Array
Medium
14.6K
711
Companies
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Python Submission:
    class Solution:
        def findKthLargest(self, nums: List[int], k: int) -> int:
            nums.sort()
            return nums[-k]




1926. Nearest Exit from Entrance in Maze
Medium
1.9K
67
Companies
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+'). You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze. Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

 

Example 1:


Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
Output: 1
Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
Initially, you are at the entrance cell [1,2].
- You can reach [1,0] by moving 2 steps left.
- You can reach [0,2] by moving 1 step up.
It is impossible to reach [2,3] from the entrance.
Thus, the nearest exit is [0,2], which is 1 step away.


Python Submission:

    class Solution:

        def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
            rows=len(maze)
            cols=len(maze[0])
            
            sx,sy=entrance
            maze[sx][sy]='s'

            for x in range(rows):
                if maze[x][0]=='.':
                    maze[x][0]='e'
                if maze[x][-1]=='.':
                    maze[x][-1]='e'
            for y in range(cols):
                if maze[0][y]=='.':
                    maze[0][y]='e'
                if maze[-1][y]=='.':
                    maze[-1][y]='e'
            
            dir=[(0,1),(1,0),(-1,0),(0,-1)]
            
            done=[[False]*cols for _ in range(rows)]

            queue=collections.deque()
            queue.append((0,sx,sy))
            done[sx][sy]=True

            while len(queue)>0:
                d,x,y=queue.popleft()
                for dx,dy in dir:
                    nx,ny=x+dx,y+dy
                    if 0<=nx<rows and 0<=ny < cols and not done[nx][ny]:
                        done[nx][ny] = True
                        if maze[nx][ny]==".":
                            queue.append((d+1,nx,ny))
                        elif maze[nx][ny] =="e":
                            return d+1


            return -1





450. Delete Node in a BST
Medium
7.9K
207
Companies
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.



 
Python Submission:

    def getMinimumKey(curr):
        while curr.left:
            curr = curr.left
        return curr
    
    

    
    def deleteNode(root, key):
    
        parent = None
    
        curr = root
    
        while curr and curr.val != key:
    
            parent = curr
    
            if key < curr.val:
                curr = curr.left
            else:
                curr = curr.right
    
        if curr is None:
            return root
    
        # 案例1：待刪除節點無子節點，即為葉子節點
        if curr.left is None and curr.right is None:
    
            if curr != root:
                if parent.left == curr:
                    parent.left = None
                else:
                    parent.right = None
    
            else:
                root = None
    
        # 案例2：要刪除的節點有兩個孩子
        elif curr.left and curr.right:
    
            #找到它的中序後繼節點
            successor = getMinimumKey(curr.right)
    
            val = successor.val
    
            deleteNode(root, successor.val)
    
            curr.val = val
    
        # 案例3：要刪除的節點只有一個孩子
        else:
            if curr.left:
                child = curr.left
            else:
                child = curr.right
    
            if curr != root:
                if curr == parent.left:
                    parent.left = child
                else:
                    parent.right = child
    
            else:
                root = child
    
        return root




    class Solution:
        def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
            root = deleteNode(root, key)
            return root