package main

import (
	"fmt"
	"strings"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительное задание: поддержка escape - последовательности

qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
*/

var m = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func Unpack(s string) string {
	var (
		result []string = make([]string, 0)
		prev   string
		escape bool
	)
	if len(s) == 0 {
		return ""
	}

	for _, r := range s {
		if string(r) == "\\" && !escape {
			escape = true
			result = append(result, prev)
			continue
		}
		if escape {
			prev = string(r)
			escape = false
			continue
		}
		count, ok := m[string(r)]
		if ok {
			if prev == "" {
				return ""
			}
			result = append(
				result,
				strings.Repeat(prev, count))
			prev = ""
			continue
		}
		if prev != "" {
			result = append(result, prev)
		}
		prev = string(r)
	}
	if prev != "" {
		result = append(result, prev)
	}

	return strings.Join(result, "")
}

// не забываем, что / это спец символ, поэтому для того, чтобы прочитать его, его нужно экранировать,
// итого в настроящей строке кол-во слешей будет в 2 раза больше, чем в той, которую видит программа

func main() {
	fmt.Println(Unpack("a4bc2d5e") == "aaaabccddddde")
	fmt.Println(Unpack("abcd") == "abcd")
	fmt.Println(Unpack("45") == "")
	fmt.Println(Unpack("") == "")
	fmt.Println(Unpack("qwe\\4\\5") == "qwe45")
	fmt.Println(Unpack("qwe\\45") == "qwe44444")
	fmt.Println(Unpack("qwe\\\\5") == "qwe\\\\\\\\\\")
}
