package main

import (
	"testing"

	"github.com/mbenaiss/fizzbuzz/internal/fizzbuzz"
)

func TestParseQuery(t *testing.T) {
	tt := []struct {
		query    map[string][]string
		expected fizzbuzz.FizzBuzz
	}{
		{
			query:    map[string][]string{},
			expected: fizzbuzz.FizzBuzz{},
		},
		{
			query: map[string][]string{
				"n1":    []string{"3"},
				"n2":    []string{"5"},
				"limit": []string{"15"},
				"str1":  []string{"fizz"},
				"str2":  []string{"buzz"},
			},
			expected: fizzbuzz.FizzBuzz{
				N1:    3,
				N2:    5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
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
