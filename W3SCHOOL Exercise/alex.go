package main

import (
	"fmt"
	"strings"
)

func Lower(w string) string {
	s := strings.Split(w, " ")

	for i := 0; i < len(s); i++ {
		if s[i] == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}

		if s[i] == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}

		if s[i] == "(cap)" {
			s[i-1] = strings.ToUpper(s[i-1][:1]) + strings.ToLower(s[i-1][1:])
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return strings.Join(s, " ")
}
func main() {
	fmt.Println(Lower("learning (up) go is (cap) a POWERFUL (low) thing"))
}
