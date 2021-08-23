package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicValidParentheses(t *testing.T) {
	res := ValidParentheses("()")

	assert.Equal(t, res, true)
}

func TestRandomValidParentheses(t *testing.T) {
	res := ValidParentheses("((((()()()))))")

	assert.Equal(t, res, true)
}

func TestValidParenthesesFalse(t *testing.T) {
	res := ValidParentheses("())(")
	assert.Equal(t, res, false)
}

func TestSingleValidParenthesesFalse(t *testing.T) {
	res := ValidParentheses(")")
	assert.Equal(t, res, false)
}
