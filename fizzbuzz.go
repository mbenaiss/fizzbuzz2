package main

import "strconv"

type FizzBuzz struct {
	n1, n2, limit int
	str1, str2    string
}

func fizzbuzz(f FizzBuzz) []string {
	result := []string{}
	for i := 1; i <= f.limit; i++ {
		var res string
		s1 := check(i, f.n1, f.str1)
		s2 := check(i, f.n2, f.str2)
		if len(s1) != 0 {
			res += s1
		}
		if len(s2) != 0 {
			res += s2
		}
		if len(res) == 0 {
			res = strconv.Itoa(i)
		}
		result = append(result, res)
	}
	return result
}

func check(i, n int, str string) string {
	if n == 0 {
		return ""
	}
	if i%n == 0 {
		return str
	}
	return ""
}
