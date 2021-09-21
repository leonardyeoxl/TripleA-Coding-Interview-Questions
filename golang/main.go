package main

import (
	"fmt"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func Solution(s string) bool {
	if len(s) == 0{
		return true
	}

	stack := []string{}
	starting_paranthesis := []string{"(", "{", "["}
	closing_paranthesis := []string{")", "}", "]"}
	_map := map[string]string{
		")": "(",
        "}": "{",
        "]": "[",
	}

	_char := ""
	for index := 0; index < len(s); index++ {
		if stringInSlice(s[index:index+1], starting_paranthesis) {
			stack = append(stack, s[index:index+1])
		} else if stringInSlice(s[index:index+1], closing_paranthesis) {
			_char, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if _map[s[index:index+1]] != _char {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	s := "()"
	result := Solution(s)
	fmt.Println(result)

	s = "()[]{}"
	result = Solution(s)
	fmt.Println(result)

	s = "(]"
	result = Solution(s)
	fmt.Println(result)

	s = "([)]"
	result = Solution(s)
	fmt.Println(result)

	s = "{[]}"
	result = Solution(s)
	fmt.Println(result)

	s = "()(){(())"
	result = Solution(s)
	fmt.Println(result)

	s = ""
	result = Solution(s)
	fmt.Println(result)

	s = "([{}])()"
	result = Solution(s)
	fmt.Println(result)
}