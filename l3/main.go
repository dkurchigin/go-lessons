package main

import (
	"errors"
	"fmt"
	"strconv"
	s "strings"
)

func ConvertStrings(string_ string) (string, error) {
	numbers_ := "1234567890"
	out_string := ""
	prev_ := ""
	prev_is_number := false

	for _, value := range string_ {
		if prev_ == "" {
			prev_ = string(value)
			if s.Contains(numbers_, prev_) {
				return "", errors.New("Not valid")
			}
			out_string += prev_
		} else {
			if prev_is_number && s.Contains(numbers_, string(value)) {
				return "", errors.New("Not valid")
			}
			if s.Contains(numbers_, string(value)) {
				number_, err := strconv.Atoi(string(value))
				if err != nil {
					fmt.Println("Something fucked up")
				} else {
					out_string += s.Repeat(prev_, number_ - 1)
				}
				prev_is_number = true
			} else {
				prev_is_number = false
				prev_ = string(value)
				out_string += string(value)
			}
		}
	}
	return out_string, nil
}

func main() {
	string_ := "a45e"
	str_, err := ConvertStrings(string_)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str_)
	}
}
