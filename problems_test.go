package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblems_problems(t *testing.T) {
	info := problems()
	assert.Equal(t, 771, info[0].Number, "wrong first problem number")
	assert.Equal(t, "Easy", info[0].Difficulty, "wrong first problem difficulty")
}

func TestProblems_next_WhenNoFilter(t *testing.T) {
	info := problems()
	currentProblem := 771
	diff := Difficulty{
		Easy:   true,
		Medium: true,
		Hard:   true,
	}
	idx, err := next(info, &currentProblem, &diff)

	assert.Nil(t, err, "wrong error")
	assert.Equal(t, 1, idx, "wrong next problem")
}

func TestProblems_next_WhenFilterMedium(t *testing.T) {
	info := problems()
	currentProblem := 771
	diff := Difficulty{
		Easy:   false,
		Medium: true,
		Hard:   true,
	}

	idx, err := next(info, &currentProblem, &diff)
	assert.Nil(t, err, "wrong error")
	assert.Equal(t, 1265, info[idx].Number, "wrong next medium problem")
}

func TestProblems_next_WhenFilterHard(t *testing.T) {
	info := problems()
	currentProblem := 102
	diff := Difficulty{Hard: true}

	idx, err := next(info, &currentProblem, &diff)
	assert.Nil(t, err, "wrong error")
	assert.Equal(t, 431, info[idx].Number, "wrong next hard problem")
}
