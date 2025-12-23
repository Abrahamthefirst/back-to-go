package main

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
// 	strs := []string{"flower", "flow", "flight"}
// 	result := longestCommonPrefix(strs)
// 	fmt.Println(result)
// }
