# TripleA-Coding-Interview-Questions

## Question 1

- Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
- An input string is valid if:
    - Open brackets must be closed by the same type of brackets.
    - Open brackets must be closed in the correct order.

### Example 1:
```
Input: s = "()"
Output: true
```
### Example 2:
```
Input: s = "()[]{}"
Output: true
```
### Example 3:
```
Input: s = "(]"
Output: false
```
### Example 4:
```
Input: s = "([)]"
Output: false
```
### Example 5:
```
Input: s = "{[]}"
Output: true
```

### Example 6
```
Input: s = "((()))"
Output: True
```

### Example 7
```
Input: s = "[()]{}"
Output: True
```
### Example 8
```
Input: s = "({[)]"
Output: False
```

### Constraints:
- 1 <= s.length <= 10^4
- s consists of parentheses only '()[]{}'.

## References

- https://leetcode.com/problems/valid-parentheses/