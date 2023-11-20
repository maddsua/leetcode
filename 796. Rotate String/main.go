package main

func rotateString(s string, goal string) bool {

	rotated := s

	for i := 0; i < len(s); i++ {

		if rotated == goal {
			return true
		}

		//	rotate
		temp := string(rotated[len(rotated)-1])
		rotated = temp + rotated[:len(rotated)-1]
	}

	return false
}

func main() {
	result := rotateString("abcde", "cdeab")
	println(result)
}
