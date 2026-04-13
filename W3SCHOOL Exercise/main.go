// package main

// func main() {
// arry := [3]int {0:12,1:13}
// var cars = [6]string {"benz", "honda", "mazda", "gis", "toyota", "inoson"}
// goLang := []string{}

// cars [3] = "cybertruck"

// fmt.Println(cars)
// fmt.Println(cars[5])
// fmt.Println(cars[3])
// fmt.Println(arry)
// fmt.Println(len(cars))
// fmt.Println(len(arry))
// fmt.Println(len(goLang))
// fmt.Println(cap(goLang))

// price := []int{12,5,8}
// price1 := []int{67,68,75}
// price2 := append(price, price1...)

// fmt.Printf("the slice will be = %v\n", price2)
// fmt.Printf("the length of the new slice is = %d\n", len(price2))
// fmt.Printf("the capacity of the new slice is = %d\n", cap(price2))

// 	newArray := [13]int{1,2,3,4,4,5,5,6,6,6,6,2,4}
// 	newSlice := newArray[4:12]

// 	fmt.Printf("new sclice = %v\n", newSlice)
// 	fmt.Printf("length of new slice = %d\n", len(newSlice))
// 	fmt.Printf("capacity of the new slice = %d\n", cap(newSlice))
// fmt.Println()
// newSlice = newArray[1:4] // Change length by re-slicing the array
//   	fmt.Printf("the re sliced = %v\n", newSlice)
//   	fmt.Printf("length = %d\n", len(newSlice))
//   	fmt.Printf("capacity = %d\n", cap(newSlice))
// fmt.Println()
// newSlice = append(newSlice, 20, 21, 22, 23) // Change length by appending items
//   fmt.Printf("myslice1 = %v\n", newSlice)
//   fmt.Printf("length = %d\n", len(newSlice))
//   fmt.Printf("capacity = %d\n", cap(newSlice))

// familyMembers := 11

// switch familyMembers {
// case 1:
// 	fmt.Println ("Godswill is the first Son")
// case 2:
// 	fmt.Println ("Alexander is the Second Son")
// case 3:
// 	fmt.Println ("Edwin is the third Son")
// case 4:
// 	fmt.Println ("Ruth is the First Daughter")
// case 5:
// 	fmt.Println ("Helen is the second daughter")
// case 6:
// 	fmt.Println ("Innocent is the fourth Son")
// case 7:
// 	fmt.Println ("Gideon is the fifth Son")
// case 8:
// 	fmt.Println("Aaron is the sixth Son")
// case 9:
// fmt.Println("Blessing is the third Daughter")
// case 10:
// 	fmt.Println("Ezekiel is the seventh Son")
// default:
// 	fmt.Println("Is not a member of the family")
// }

//working with multiple switch cases
// familyMembers := 10

// switch familyMembers {
// case 1, 2, 3, 6, 7, 8, 10:
// 	fmt.Println("One of the boys in the family")
// case 4, 5, 9:
// 	fmt.Println("One the girls in the family")
// default:
// 	fmt.Println("Is not a member of the family")
// }

// working with loops
// for i := 0; i < 10; i++ {
// 	fmt.Printf("the value of i is %d\n", i)
// }

// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		break
// 	}
// 	fmt.Printf("the value of i is %d\n", i)
// }

// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		continue
// 	}
// 	fmt.Printf("the value of i is %d\n", i)
// }

//the break statement: is used to exit a loop when a certain condition is met, while the continue statement is used to skip the current iteration of a loop and move on to the next one. In the example above, the break statement will stop the loop when i equals 5, while the continue statement will skip the iteration when i equals 5 and continue with the next iteration.
// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Printf("the value of i is %d\n", i)
// }

// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		fmt.Printf("the value of i is %d\n", i)
// 	}
// }

// for i := 0; i < 10; i++ {
// 	if i%2 != 0 {
// 		fmt.Printf("the value of i is %d\n", i)
// 	}
// }
// for i := 0; i < 10; i++ {
// 	if i%2 == 0 {
// 		fmt.Printf("the value of i is %d\n", i)
// 	}
// }

// for i := 0; i < 10; i++ {
// 	if i%2 != 0 {
// 		fmt.Printf("the value of i is %d\n", i)
// 	}
// }

// for i := 0; i < 10; i++ {
// 	if i == 3 {
// 		fmt.Printf("the value of i is %d\n", i)
// 	} else {
// 		fmt.Printf("the value of i is %d\n", i)
// 	}
// }

//func ToUpperLastWord(words []string) []string
// var words = []string{"hello", "world", "golang", "programming"}

// learning about functions
// func ToUpperLastWord(words []string) []string {
// 	for i := 0; i < len(words); i++ {
// 		if i == len(words)-1 {
// 			words[i] = strings.ToUpper(words[i])
// 		}
// 	}
// 	return words

package main

import (
	"fmt"
	"strings"
)

func main() {
	// text := "Learning Go is powerful"
	// words := strings.Split(text, " ")

	// fmt.Println("Words in text:")
	// for _, word := range words {
	// 	fmt.Println(word)
	// }

	// joined := strings.Join(words, "_")
	// fmt.Println("Joined with underscores:", joined)

	// Splitting the sentence into words
	sentence := "Go makes coding fun"
	words := strings.Split((sentence), " ")

	fmt.Println("Original words in the sentence provided:")
	for _, words := range words {
		fmt.Println(words)
	}
}
