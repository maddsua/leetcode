package main

type Fragment struct {
	repeats int
	token   string
}

func decodeString(s string) string {

	fragments := make([]Fragment, 0)

	var next *Fragment

	for _, v := range s {

	}

}

func main() {

	println(decodeString("3[a]2[bc]"))

}
