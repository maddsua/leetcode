package main

import "regexp"

func isNumber(s string) bool {

	numberRegexes := []*regexp.Regexp{
		regexp.MustCompile("^(\\+|\\-)?((\\d+\\.)|(\\.\\d+)|(\\d+\\.\\d+)|(\\d+))$"),
		regexp.MustCompile("(?i)^(\\+|\\-)?((\\d+(\\.\\d*)?e(\\+|\\-)?\\d+)|(\\.\\d+e(\\+|\\-)?\\d+))$"),
	}

	for _, v := range numberRegexes {
		if v.MatchString(s) {
			return true
		}
	}

	return false
}

func main() {
	println(isNumber("125"))
	println(isNumber("125.75"))
	println(isNumber("2e0"))
}
