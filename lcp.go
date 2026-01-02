package main

import (
	"fmt"
	"strings"
	"unicode"
)

// import "fmt"

// func longestCommonPrefix(strs []string) string {

// 	for index, firstRune := range strs[0] {

// 		for _, value := range strs {

// 			if index == len(value) || value[index] != byte(firstRune) {
// 				return string(strs[0][:index])
// 			}

// 		}
// 	}
// 	return strs[0]

// }

// func fizzBuzz(n int) []string {
//     result := []string{}
//     for i := 1; i <= n; i++{
//         if n % 3 == 0 && n % 5 == 0{
//             result = append(result, "FizzBuzz")
//         } else if n % 3 == 0 {
//             append(result, "Fizz")
//         } else if n % 5 == 0{
//             append(result, "Buzz")
//         } else {
//             append(result, string(i))
//         }
//     }
//     return result
// }

// func main() {
// 	paragraph := "Bob"
// 	result := mostCommonWord(paragraph, []string{})
// 	fmt.Println("\nMax word", result)
// }

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]bool)

	for _, b := range banned {
		bannedMap[strings.ToLower(b)] = true
	}
	cache := make(map[string]int)
	word := ""
	maxWord := ""
	maxCount := 0

	paragraph += " "

	for _, char := range paragraph {

		if unicode.IsLetter(char) {
			word += string(unicode.ToLower(char))
		} else {
			if len(word) > 0 {
				if !bannedMap[word] {
					cache[word]++
					if cache[word] >= maxCount {
						maxCount = cache[word]
						maxWord = word
					}
				}
				word = ""
			}
		}

	}
	fmt.Println("cache", cache)
	return maxWord
}
