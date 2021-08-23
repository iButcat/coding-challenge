package main

/*
func ValidParentheses(parens string) bool {
	var counter int
	for i := 0; i < len(parens); i++ {
		if string(parens[i]) == "(" {
			counter++
		} else {
			counter--
		}
	}
	if counter == 0 {
		return true
	} else {
		return false
	}
}
*/

func ValidParentheses(parens string) bool {
	var stack []rune

	if len(parens)%2 != 0 {
		return false
	}

	for i := 0; i < len(parens); i++ {
		if parens[i] == '(' {
			stack = append(stack, rune(parens[i]))
		}

		if parens[i] == ')' && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
