package main

type Difficulty struct {
	Easy, Medium, Hard bool
}

func (d Difficulty) Match(str string) bool {
	if d.Easy == true && d.Medium == true && d.Hard == true {
		return true
	}

	switch str {
	case "Easy":
		return d.Easy
	case "Medium":
		return d.Medium
	case "Hard":
		return d.Hard
	}

	return false
}
