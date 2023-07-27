1456. Maximum Number of Vowels in a Substring of Given Length
Medium
2.9K
98
Companies
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

 

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.


Python Submission:
    class Solution:
        def maxVowels(self, s: str, k: int) -> int:
            vowels = {'a', 'e', 'i', 'o', 'u'}
            

            current_vowels = 0

            window_left = 0
            window_right = k - 1

            for c in s[window_left:window_right+1]:
                if c in vowels:
                    current_vowels+=1

            max_vowels = current_vowels

            while window_right < len(s)-1:
                if s[window_right+1] in vowels :
                    if s[window_left] not in vowels:
                        current_vowels+=1
                    else:
                        pass
                else:
                    if s[window_left] in vowels:
                        current_vowels-=1
                    else:
                        pass

                if current_vowels > max_vowels:
                    max_vowels = current_vowels
                
                window_right+=1
                window_left+=1
                
            return max_vowels

Go Submission:
  
    func maxVowels(s string, k int) int {

        vowels := map[byte]bool{
            'a': true,
            'e': true,
            'i': true,
            'o': true,
            'u': true,
        }

        current_vowels := 0

        for i:=0; i< k; i++{
            if vowels[s[i]]{
                current_vowels +=1
            }
        }


        max_vowels := current_vowels

        for i,j:=0,k-1; j< len(s)-1; i,j=i+1,j+1{
        

        if vowels[s[j+1]]{
            if vowels[s[i]]{

            }else{
                    current_vowels+=1
            }
        }else{
            if vowels[s[i]]{
                    current_vowels-=1
            }else{
                
            }
        }
                        

            if current_vowels > max_vowels{
                max_vowels = current_vowels
            }


        }
        return max_vowels
    }




1657. Determine if Two Strings Are Close
Medium
2.5K
123
Companies
Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

 

Example 1:

Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"


Python Submission:

    class Solution:
        def closeStrings(self, word1: str, word2: str) -> bool:
            
            if len(word1)!=len(word2):
                return False

            d1 = {}
            d2 = {}

            for c in word1:
                if c in d1:
                    d1[c]+=1
                else:
                    d1[c]=1
            
            for c in word2:
                if c in d2:
                    d2[c]+=1
                else:
                    d2[c]=1
            
            s1 = set(d1.keys())
            l1 = list(d1.values())
            l1.sort()

            s2 = set(d2.keys())
            l2 = list(d2.values())
            l2.sort()

            return l1==l2 and s1==s2


Go Submission:

    func closeStrings(word1 string, word2 string) bool {
    
        if len(word1)!=len(word2){
            return false
        }

        b1 := map[byte]bool{}
        b2 := map[byte]bool{}

        d1 := map[byte]int{}
        d2 := map[byte]int{}

        for i:=0;i<len(word1);i+=1{
            b1[word1[i]]=true
            d1[word1[i]]+=1
        }

        for j:=0;j<len(word2);j+=1{
            b2[word2[j]]=true
            d2[word2[j]]+=1
        }

        c1 := []int{}
        c2 := []int{}

        for  _, value := range d1 {
            c1 = append(c1, value)
        }

        for  _, value := range d2 {
            c2 = append(c2, value)
        }

        sort.Ints(c1)
        sort.Ints(c2)


        return reflect.DeepEqual(c1, c2) && reflect.DeepEqual(b1, b2)

    }


649. Dota2 Senate
Medium
1.9K
1.4K
Companies
In the world of Dota2, there are two parties: the Radiant and the Dire.

The Dota2 senate consists of senators coming from two parties. Now the Senate wants to decide on a change in the Dota2 game. The voting for this change is a round-based procedure. In each round, each senator can exercise one of the two rights:

Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
Announce the victory: If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and decide on the change in the game.
Given a string senate representing each senator's party belonging. The character 'R' and 'D' represent the Radiant party and the Dire party. Then if there are n senators, the size of the given string will be n.

The round-based procedure starts from the first senator to the last senator in the given order. This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.

Suppose every senator is smart enough and will play the best strategy for his own party. Predict which party will finally announce the victory and change the Dota2 game. The output should be "Radiant" or "Dire".

 

Example 1:

Input: senate = "RD"
Output: "Radiant"
Explanation: 
The first senator comes from Radiant and he can just ban the next senator's right in round 1. 
And the second senator can't exercise any rights anymore since his right has been banned. 
And in round 2, the first senator can just announce the victory since he is the only guy in the senate who can vote.

Python Submission:
    class Solution:
        def predictPartyVictory(self, senate: str) -> str:
            n, R, D = len(senate), [], []

            for i,c in enumerate(senate):
                if c == "R":
                    R.append(i)
                else:
                    D.append(i)

            while R and D:
                r, d = R.pop(0), D.pop(0)

                if r < d:
                    R.append(r+n)
                else:
                    D.append(d+n)
            return "Radiant" if R else "Dire"

Go Submission:
    func predictPartyVictory(senate string) string {
        n, R, D := len(senate), []int{}, []int{}

        for i:=0;i<n;i+=1{
            if senate[i]=='R'{
                R = append(R,i)
            }else{
                D = append(D,i)
            }
        }

        for len(R)>0 && len(D)>0{
            r := R[0]
            d := D[0]
            if r<d{
                D = D[1:]
                popped := R[0]
                R = R[1:]
                R = append(R,popped+n)
            }else{
                R = R[1:]
                popped := D[0]
                D = D[1:]
                D = append(D,popped+n)
            }
        }

        if len(R) > 0{
            return "Radiant"
        }
        return "Dire"


    }