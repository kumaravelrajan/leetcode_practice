class Solution:
    def romanToInt(self, s: str) -> int:
        dict_roman_to_integer = {
            "I":             1,
            "V":             5,
            "X":             10,
            "L":             50,
            "C":             100,
            "D":             500,
            "M":             1000,
            "IX":            9,
            "IV":            4,
            "XL":            40,
            "XC":            90,
            "CD":            400,
            "CM":            900
        }

        integer_to_return = 0
        index = 0

        while index < len(s):
            if index + 1 < len(s):
                if(s[index] + s[index+1] in dict_roman_to_integer):
                    integer_to_return += dict_roman_to_integer[s[index] + s[index+1]]
                    index += 2
                else:
                    integer_to_return += dict_roman_to_integer[s[index]]
                    index += 1

            else:
                integer_to_return += dict_roman_to_integer[s[index]]
                index += 1

        print(integer_to_return)


sol = Solution()
sol.romanToInt("MCDLXXVI")