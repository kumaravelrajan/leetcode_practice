# link - https://leetcode.com/problems/longest-palindromic-substring/

class Solution:
    def longestPalindrome(self, s: str) -> str:
        temp_pal = ""
        longest_pal = ""

        if len(s) == 1:
            return s


        # Finding aba in "babad"
        for index, char in enumerate(s[1:len(s)-1], 1):
            if s[index-1] == s[index+1]:
                temp_pal = s[index-1:index+2]

                i = 2

                while 0 <= index - i < len(s) and 0 <= index + i < len(s) and s[index - i] == s[index + i]:
                    temp_pal = s[index - i] + temp_pal + s[index+i]
                    i += 1
                
                if len(temp_pal) > len(longest_pal):
                    longest_pal = temp_pal
                
            else:
                temp_pal = char
                if len(longest_pal) < len(temp_pal):
                    longest_pal = temp_pal

        # Finding bb in "cbbd"

        for index, char in enumerate(s):
            i = index + 1
            temp_pal = char

            while 0 <= i < len(s):
                if s[i] == char:
                    temp_pal += s[i]
                    i += 1
                else:
                    break

            j = 1
            while True:
                if 0 <= index - j < len(s) and 0 <= i < len(s) and s[index - j] == s[i]:
                    temp_pal = s[index-j] + temp_pal + s[i]
                    j += 1
                    i += 1
                else:
                    break

            if len(temp_pal) > len(longest_pal):
                longest_pal = temp_pal

        print( longest_pal)
            
sol = Solution()

sol.longestPalindrome("lejyqjcpluiggwlmnumqpxljlcwdsirzwlygexejhvojztcztectzrepsvwssiixfmpbzshpilmojehqyqpzdylxptsbvkgatzdlzphohntysrbrcdgeaiypmaaqilthipjbckkfbxtkreohabrjpmelxidlwdajmkndsdbbaypcemrwlhwbwaljacijjmsaqembgtdcskejplifnuztlmvasbqcyzmvczpkimpbbwxdtviptzaenkbddaauyvqppagvqfpednnckooxzcpuudckakutqyknuqrxjgfdtsxsoztjkqvfvelrklforpjnrbvyyvxigjhkjmxcphjzzilvbjbvwiwnnkbmboiqamgoimujtswdqesighoxsprhnsceshotakvmoxqkqjvbpqucvafiuqwmrlfjpjijbctfupywkbawquchbclgvhxbanybret")
