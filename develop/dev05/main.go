package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Объявляем флаги для утилиты
	after := flag.Int("A", 0, "Печатать +N строк после совпадения")
	before := flag.Int("B", 0, "Печатать +N строк до совпадения")
	context := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "Подсчитать количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invert := flag.Bool("v", false, "Исключать совпадения")
	fixed := flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "Напечатать номер строки")

	// Парсим флаги командной строки
	flag.Parse()

	// Получаем паттерн для поиска
	pattern := flag.Arg(0)

	// Проверяем, есть ли паттерн
	if pattern == "" {
		fmt.Println("Укажите паттерн для поиска")
		return
	}

	// Проверяем, использовать ли флаг -F для точного совпадения

	// Создаем регулярное выражение на основе паттерна
	var regex *regexp.Regexp

	switch {
	case *fixed:
		//полное совпадение, игнорирование всех спецсимволов
		regex = regexp.MustCompile(regexp.QuoteMeta(pattern))
	case !*fixed && *ignoreCase:
		//игнорирование регистра
		regex = regexp.MustCompile("(?i)" + pattern)
	default:
		regex = regexp.MustCompile(pattern)
	}

	// Открываем файл или используем stdin
	var input *os.File
	if flag.NArg() > 1 {
		fileName := flag.Arg(1)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Ошибка при открытии файла: %v\n", err)
			return
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	// Читаем строки из входного потока и применяем фильтры
	scanner := bufio.NewScanner(input)

	matchLines := []int{}
	outputBuffer := []string{}

	for scanner.Scan() {

		line := scanner.Text()
		match := regex.MatchString(line)

		outputBuffer = append(outputBuffer, line)

		if match {
			matchLines = append(matchLines, len(outputBuffer))
		}

	}

	switch {
	case *after > 0 && *before > 0 || *context > 0:
		flag := false
		for _, num := range matchLines {
			if flag {
				fmt.Println("--")
				flag = false
			}
			for i := num - *before; i <= num+*after; i++ {
				if i > 0 && i < len(outputBuffer) {
					if *lineNum {
						fmt.Printf("%d: %s\n", i, outputBuffer[i-1])
					} else {
						fmt.Println(outputBuffer[i-1])
					}
				}
			}
			flag = true
		}
	case *after > 0:
		for _, num := range matchLines {
			for i := num; i <= num+*after; i++ {
				if i > 0 && i < len(outputBuffer) {
					if *lineNum {
						fmt.Printf("%d: %s\n", i, outputBuffer[i-1])
					} else {
						fmt.Println(outputBuffer[i-1])
					}
				}
			}
		}
	case *before > 0:
		for _, num := range matchLines {
			for i := num - *before; i <= num; i++ {
				if i > 0 && i < len(outputBuffer) {
					if *lineNum {
						fmt.Printf("%d: %s\n", i, outputBuffer[i-1])
					} else {
						fmt.Println(outputBuffer[i-1])
					}
				}
			}
		}
	case *invert:
		for i, line := range outputBuffer {
			if isInArray(i+1, matchLines) {
				continue
			}
			if *lineNum {
				fmt.Printf("%d: %s\n", i+1, line)
			} else {
				fmt.Println(line)
			}
		}

	}

	// Печатаем количество совпадений, если указан флаг -c
	if *count {
		fmt.Printf("Количество совпадений: %d\n", len(matchLines))
	}
}

func isInArray(target int, arr []int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}
