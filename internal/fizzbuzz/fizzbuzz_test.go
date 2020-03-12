package fizzbuzz

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
				N1:    3,
				N2:    5,
				Limit: 10,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				N1:    3,
				N2:    5,
				Limit: 15,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expected: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				Limit: 3,
				N1:    2,
				Str1:  "fizz",
			},
			expected: []string{"1", "fizz", "3"},
			success:  true,
		},
		test{
			f: FizzBuzz{
				N1:    4,
				Limit: 4,
				Str1:  "fizz",
			},
			expected: []string{"1", "2", "3", "buzz"},
			success:  false,
		},
		test{
			f: FizzBuzz{
				Limit: 5,
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
		actual := Get(g.f)
		if !reflect.DeepEqual(actual, g.expected) && g.success {
			t.Fatalf("\nTest %d : \nexpected \n%+v, \ngot \n%+v", i+1, g.expected, actual)
		}
	}
}
