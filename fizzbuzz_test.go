package main

import (
	"reflect"
	"testing"
)

type test struct {
	f        FizzBuzz
	expected []string
	success  bool
}

func TestFizzBuzz(t *testing.T) {
	tt := []test{
		test{
			f: FizzBuzz{
				n1:    3,
				n2:    5,
				limit: 10,
				str1:  "fizz",
				str2:  "buzz",
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				n1:    3,
				n2:    5,
				limit: 15,
				str1:  "fizz",
				str2:  "buzz",
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				limit: 3,
				n1:    2,
				str1:  "fizz",
			},
			expected: []string{"1", "fizz", "3"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				n1:    4,
				limit: 4,
				str1:  "fizz",
			},
			expected: []string{"1", "2", "3", "buzz"},
			success:  false,
		},
		test{
			f: FizzBuzz{
				limit: 5,
			},
			expected: []string{"1", "2", "3", "4", "5"},
			success:  true,
		},
		test{
			expected: []string{},
			success:  true,
		},
	}
	for i, g := range tt {
		actual := fizzbuzz(g.f)
		if !reflect.DeepEqual(actual, g.expected) && g.success {
			t.Fatalf("\nTest %d : \nexpected \n%+v, \ngot \n%+v", i+1, g.expected, actual)
		}
	}
}
