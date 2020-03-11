package main

import "testing"

func TestParseQuery(t *testing.T) {
	tt := []struct {
		query    map[string][]string
		expected FizzBuzz
	}{
		{
			query:    map[string][]string{},
			expected: FizzBuzz{},
		},
		{
			query: map[string][]string{
				"n1":    []string{"3"},
				"n2":    []string{"5"},
				"limit": []string{"15"},
				"str1":  []string{"fizz"},
				"str2":  []string{"buzz"},
			},
			expected: FizzBuzz{
				n1:    3,
				n2:    5,
				limit: 15,
				str1:  "fizz",
				str2:  "buzz",
			},
		},
	}

	for i, tc := range tt {
		actual := parseQuery(tc.query)
		if actual != tc.expected {
			t.Fatalf("\nTest %d : \nexpected %+v \n   got      %+v", i, tc.expected, actual)
		}
	}
}
