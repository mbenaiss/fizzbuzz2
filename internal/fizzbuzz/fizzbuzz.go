package fizzbuzz

import "strconv"

type FizzBuzz struct {
	N1, N2, Limit int
	Str1, Str2    string
}

func Get(f FizzBuzz) []string {
	result := []string{}
	for i := 1; i <= f.Limit; i++ {
		var res string
		s1 := check(i, f.N1, f.Str1)
		s2 := check(i, f.N2, f.Str2)
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