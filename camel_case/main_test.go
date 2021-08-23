package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithuUnderscore(t *testing.T) {
	res := ToCamelCase("the_Stealth_Warrior")

	assert.Equal(t, res, "theStealthWarrior")
}

func TestWithUpperAndUnderscore(t *testing.T) {
	res := ToCamelCase("The_Stealth_Warrior")

	assert.Equal(t, res, "TheStealthWarrior")
}

func TestWithUpperScore(t *testing.T) {
	res := ToCamelCase("The-Stealth-Warrior")

	assert.Equal(t, res, "TheStealthWarrior")
}

func TestWithscore(t *testing.T) {
	res := ToCamelCase("the-Stealth-Warrior")

	assert.Equal(t, res, "theStealthWarrior")
}

func TestWithUpperSpace(t *testing.T) {
	res := ToCamelCase("The Stealth Warrior")

	assert.Equal(t, res, "TheStealthWarrior")
}

func TestWithSpace(t *testing.T) {
	res := ToCamelCase("the Stealth Warrior")

	assert.Equal(t, res, "theStealthWarrior")
}
