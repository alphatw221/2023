package main

import (
	"strings"
)

func reverseWords(s string) string {

	str1 := strings.Split(s, " ")

	ret := ""

	for i := len(str1) - 1; i >= 0; i-- {

		if str1[i] == "" {
			continue
		}

		ret += " " + str1[i]
	}

	return ret[1:]

}

//
