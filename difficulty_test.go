package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifficulty_Match_WhenEasy(t *testing.T) {
	d := Difficulty{
		Easy: true,
	}

	assert.Equal(t, true, d.Match("Easy"), "wrong match Easy")
	assert.Equal(t, false, d.Match("Medium"), "wrong match Medium")
	assert.Equal(t, false, d.Match("Hard"), "wrong match Hard")
}

func TestDifficulty_Match_WhenMedium(t *testing.T) {
	d := Difficulty{
		Medium: true,
	}

	assert.Equal(t, false, d.Match("Easy"), "wrong match Easy")
	assert.Equal(t, true, d.Match("Medium"), "wrong match Medium")
	assert.Equal(t, false, d.Match("Hard"), "wrong match Hard")
}

func TestDifficulty_Match_WhenHard(t *testing.T) {
	d := Difficulty{
		Hard: true,
	}

	assert.Equal(t, false, d.Match("Easy"), "wrong match Easy")
	assert.Equal(t, false, d.Match("Medium"), "wrong match Medium")
	assert.Equal(t, true, d.Match("Hard"), "wrong match Hard")
}
