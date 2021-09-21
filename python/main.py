class Solution:
    def isValid(self, s):
        # Fill this in.
        if len(s) == 0:
            return True

        stack = []
        starting_paranthesis = ['(', '{', '[']
        closing_paranthesis = [')', '}', ']']
        _map = {
            ')': '(',
            '}': '{',
            ']': '['
        }

        for char in s:
            if char in starting_paranthesis:
                stack.append(char)
            elif char in closing_paranthesis:
                _char = stack.pop()
                if _map[char] != _char:
                    return False
        
        if len(stack) > 0:
            return False
        return True


# Test Program
s = "()" 
# should return True
print(Solution().isValid(s))

s = "()[]{}" 
# should return True
print(Solution().isValid(s))

s = "(]" 
# should return False
print(Solution().isValid(s))

s = "([)]" 
# should return False
print(Solution().isValid(s))

s = "{[]}"
# should return True
print(Solution().isValid(s))

s = "()(){(())" 
# should return False
print(Solution().isValid(s))

s = ""
# should return True
print(Solution().isValid(s))

s = "([{}])()"
# should return True
print(Solution().isValid(s))