package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Объявляем флаги для утилиты
	after := flag.Int("A", 0, "Печатать +N строк после совпадения")
	before := flag.Int("B", 0, "Печатать +N строк до совпадения")
	context := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "Подсчитать количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invert := flag.Bool("v", false, "Исключать совпадения")
	fixed := flag.Bool("F", false, "Точное совпадение со строкой")
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
	lineNumber := 0
	outputBuffer := []string{}
	inContext := false
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Применяем фильтры
		match := false
		if *ignoreCase {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else {
			match = strings.Contains(line, pattern)
		}

		if (*invert && !match) || (!*invert && match) {
			if *count {
				matchCount++
			} else if *lineNum {
				fmt.Printf("%d:%s\n", lineNumber, line)
			} else {
				outputBuffer = append(outputBuffer, line)
				inContext = true
			}
		} else if inContext {
			// Печатаем строки вокруг совпадения
			if *before > 0 {
				start := len(outputBuffer) - *before
				if start < 0 {
					start = 0
				}
				for _, contextLine := range outputBuffer[start:] {
					fmt.Println(contextLine)
				}
			} else {
				for _, contextLine := range outputBuffer {
					fmt.Println(contextLine)
				}
			}

			// Печатаем строки после совпадения
			if *after > 0 {
				for i := 1; i <= *after; i++ {
					if scanner.Scan() {
						fmt.Println(scanner.Text())
					}
				}
			}

			outputBuffer = []string{}
			inContext = false
		}
	}

	// Печатаем оставшиеся строки в буфере, если есть
	if inContext {
		for _, contextLine := range outputBuffer {
			fmt.Println(contextLine)
		}
	}

	// Печатаем количество совпадений, если указан флаг -c
	if *count {
		fmt.Printf("Количество совпадений: %d\n", matchCount)
	}
}
