package main

import "testing"

type TestCase struct {
	Input    string
	Expected string
}

var testCases = []TestCase{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
	{"qwe\\4\\5", "qwe45"},
	{"qwe\\45", "qwe44444"},
	{"qwe\\\\5", "qwe\\\\\\\\\\"},
}

func TestUnpackString(t *testing.T) {
	for _, tc := range testCases {
		result := Unpack(tc.Input)
		if result != tc.Expected {
			t.Errorf("Для входных данных: %s, ожидается: %s, получено: %s", tc.Input, tc.Expected, result)
		}
	}
}
