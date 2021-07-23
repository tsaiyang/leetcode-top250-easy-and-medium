package main

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) < 1 || !isPair(stack[len(stack)-1], s[i]) {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func isPair(b1, b2 byte) bool {
	if (b1 == '(' && b2 == ')') || (b1 == '[' && b2 == ']') || (b1 == '{' && b2 == '}') {
		return true
	}
	return false
}
