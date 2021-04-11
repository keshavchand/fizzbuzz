package main

import "strings"

type nos_and_str struct {
	num int64
	str string
}

func fizzbuzz(numbers_and_strings []nos_and_str, number int64) string {
	var return_string strings.Builder
	for _, idx := range numbers_and_strings {
		if number%idx.num == 0 {
			return_string.Write([]byte(idx.str))
		}
	}

	return return_string.String()
}

func main() {
	conditionals := make([]nos_and_str, 0)
	conditionals = append(conditionals, nos_and_str{3, "fizz"})
	conditionals = append(conditionals, nos_and_str{5, "buzz"})
	conditionals = append(conditionals, nos_and_str{7, "damn"})
	println(fizzbuzz(conditionals, 35))
	println(fizzbuzz(conditionals, 15))

}
