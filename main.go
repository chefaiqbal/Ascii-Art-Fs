package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To check the arguments is  exactly the same format :Otherwise check the error
	bannerType := "standard"
	if len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	} else if len(os.Args) >= 3 {
		if os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("EX: go run . something standard")
			return
		}
		bannerType = os.Args[2]
	}

	banner := bannerType + ".txt"

	data, _ := os.ReadFile(banner) //Read the data from the given file

	//split the input into words and the .txt file into strings
	words := strings.Split(os.Args[1], "\\n")
	lines := strings.Split(string(data), "\n")

	for index := range words {
		if words[index] == "" { // break if there is no word
			break

		}

		letters := strings.Split(words[index], "") //slice the words into symbols
		var ascii []int
		for i := range letters {
			l := int([]rune(letters[i])[0]) //store the rune value of the letter in l
			ascii = append(ascii, l)        //append l value to ascii untill all letters are appended
		}
		for j := 1; j < 9; j++ { //loop 8 times because a row is 8 lines long
			str := ""
			for k := range ascii { //loop through all runes in ascii
				line := (ascii[k] - 32) * 9 //this finds the row before the ascii symbol in the txt file
				str += lines[line+j]        //gives the string all the lines into the string
			}
			fmt.Println(str)

		}
	}
}
