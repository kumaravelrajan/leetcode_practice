from typing import List
from collections import defaultdict

class Solution:

    def get_pairs_prefix_or_suffix_with_both(pos1or2_lookup, both_pos_lookup):
        score = 0

        

    def score(self, cards: List[str], x: str) -> int:
        x_in_pos1_lookup = defaultdict(int)
        x_in_pos2_lookup = defaultdict(int)
        x_in_both_pos_lookup = defaultdict(int)
        score = 0
        balance = None

        for card in cards:
            if card[0] == x and card[1] == x:
                x_in_both_pos_lookup[x] += 1
            
            elif card[0]==x and card[1] != x:
                x_in_pos1_lookup[card[1]] += 1

            elif card[0] != x and card[1] == x:
                x_in_pos2_lookup[card[0]] += 1

        for key, value in x_in_pos1_lookup.items():
            if balance is None:
                balance = value
            else:
                score += min(balance, value)
                balance = abs(balance - value)

        # If there are no cards such as "xy" to load into pos1_lookup initialize balance to 0
        if balance is None:
            balance = 0

        # If there are still cards remaining s.t. 'xy' for some y then check if there exist any cards in x_in_both_pos_lookup so that these can be paired up. 
        if balance > 0:
            if x_in_both_pos_lookup:
                if balance > x_in_both_pos_lookup[x]:
                    score += x_in_both_pos_lookup[x]
                    x_in_both_pos_lookup[x] = 0
                
                else:
                    score += balance
                    x_in_both_pos_lookup[x] = x_in_both_pos_lookup[x] - balance

        # Now focus on pos2_lookup
        balance = None # Because balance of pos1_lookup worthless in pos2_lookup since no card in pos1_lookup can be compatible with any card in pos2_lookup. 

        for key, value in x_in_pos2_lookup.items():
            if balance is None:
                balance = value
            else:
                score += min(balance, value)
                balance = abs(balance - value)

        if balance is None: 
            # If there was no entry in pos2_lookup, balance never changed from None. Hence initialize it to 0.
            balance = 0

        if balance > 0:
            if x_in_both_pos_lookup[x] > 0:
                # This means there are still unused cards in both_pos_lookup which can be paired with leftovers from pos2_lookup

                score += min(balance, x_in_both_pos_lookup[x])

        return score

sol = Solution()
print(sol.score(["aa", "ab", "ba", "ac"], "a"))

print(sol.score(["aa","ab","ba"], "a"))

print(sol.score(["aa","ab","ba","ac"], "b"))

print(sol.score(["ab","bb"], "b"))
