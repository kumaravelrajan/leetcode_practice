class Solution:
    def myAtoi(self, s: str) -> int:

        if not s:
            return 0

        s = s.strip() # Strip whitespaces

        acceptable_chars_lookup = {
            '0': True,
            '1': True,
            '2': True,
            '3': True,
            '4': True,
            '5': True,
            '6': True,
            '7': True,
            '8': True,
            '9': True
        }

        target_string = ""
        sign = ""

        # Finding sign
        if len(s) > 0:
            if s[0] == '-':
                sign += '-'
                s = s[1:]
            elif s[0] == '+':
                sign += '+'
                s = s[1:]

        if not s:
            return 0

        # Strip leading zeroes
        # index = 0
        # for index, char in enumerate(s):
        #     if char == '0':
        #         continue
        #     else:
        #         break

        
        # if s[index] not in acceptable_chars_lookup:
        #     return 0

        # s = s[index:]            

        # strip leading zeroes
        is_leading_zeroes_removed = False
        index = 0
        for index, char in enumerate(s):
            if (char != '0') and (not is_leading_zeroes_removed):

                if char in acceptable_chars_lookup:
                    is_leading_zeroes_removed = True
                    target_string += char
                else:
                    # Char not 0, leading zeroes have not been cleared yet and char is not even in acceptable_chars_lookup
                    return 0

            elif (char == '0') and not (is_leading_zeroes_removed):
                continue

            elif (is_leading_zeroes_removed):

                # leading zeroes have been removed. If char is not acceptable char, simply break                
                if char in acceptable_chars_lookup:
                    target_string += char
                else:
                    break

        if not len(target_string) > 0:
            return 0
        
        target_number = int(sign + target_string)

        if target_number < (-2**31):
            return -2**31
        elif target_number > (2**31 - 1):
            return 2**31 - 1

        return target_number


sol = Solution()
print(sol.myAtoi("+"))
print(sol.myAtoi("42"))
print(sol.myAtoi("-042"))
print(sol.myAtoi("1337c0d3"))
print(sol.myAtoi("0-1"))
print(sol.myAtoi("000"))

        
            