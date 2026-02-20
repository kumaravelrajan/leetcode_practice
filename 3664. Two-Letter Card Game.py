from typing import List
from collections import defaultdict

class Solution:

    def score(self, cards: List[str], x: str) -> int:

        set1 = defaultdict(int)
        tset1 = 0
        freqset1 = 0
        set2 = defaultdict(int)
        tset2 = 0
        freqset2 = 0
        tbothx = 0
        answer = 0
        used = 0

        for card in cards:
            if card[0] == x and card[1] != x:
                tset1 += 1
                set1[card[1]] += 1

                if freqset1 < set1[card[1]]:
                    freqset1 = set1[card[1]]
            
            elif card[0] != x and card[1] == x:
                tset2 += 1
                set2[card[0]] += 1

                if freqset2 < set2[card[0]]:
                    freqset2 = set2[card[0]]
            
            elif card[0] == x and card[1] == x:
                tbothx += 1


        # Pairing cards in same set
        if freqset1 > tset1 // 2:
            answer += tset1 - freqset1
            tset1 = tset1 - (2 * (tset1 - freqset1))

        elif freqset1 <= tset1 // 2:
            answer += tset1 // 2
            tset1 = tset1 % 2

        if freqset2 > tset2 // 2:
            answer += tset2 - freqset2
            tset2 = tset2 - (2 * (tset2 - freqset2))

        elif freqset2 <= tset2:
            answer += tset2 // 2
            tset2 = tset2 % 2

        #Pairing leftover cards in set1 (and set2) with cards in double_x.

        while tset1 != 0 and tbothx != 0:
            tset1 -= 1
            tbothx -= 1
            used += 1

        while tset2 != 0 and tbothx != 0:
            tset2 -= 1
            tbothx -= 1
            used += 1

        # Spitting pairs in answer if pairs exist in tbothx so that one additional pair can be created

        while answer > 0 and tbothx > 1:
            answer -= 1
            tbothx -= 2
            used += 2

        return answer + used

        

sol = Solution()
print(sol.score(["aa", "ab", "ba", "ac"], "a"))

print(sol.score(["aa","ab","ba"], "a"))

print(sol.score(["aa","ab","ba","ac"], "b"))

print(sol.score(["ab","bb"], "b"))
