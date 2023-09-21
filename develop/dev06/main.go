package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//cat input.txt | main -d " " -f 2,3,4

func main() {
	// Определение флагов
	fieldsFlag := flag.String("f", "", "выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Разделитель
	delimiter := *delimiterFlag
	// Выбранные поля
	selectedFields := make(map[int]bool)

	// Обработка флага -f
	if *fieldsFlag != "" {
		fieldsList := strings.Split(*fieldsFlag, ",")
		for _, field := range fieldsList {
			fieldIndexes := parseFieldIndex(field)
			//if fieldIndex > 0 {
			//	selectedFields[fieldIndex] = true
			//}
			if fieldIndexes != nil {
				for _, fieldIndex := range fieldIndexes {
					selectedFields[fieldIndex] = true
				}
			}
		}
	}

	//fmt.Println(selectedFields)

	if *fieldsFlag == "" {
		return
	}

	// Чтение строк из STDIN и обработка
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, delimiter)
		containsDelimiter := strings.Contains(line, delimiter)
		switch {
		case *separatedFlag && containsDelimiter:
			outputParts := []string{}
			for i, part := range parts {
				// Если выбрано конкретное поле (колонка) или выбор не задан (выводим все поля)
				if _, ok := selectedFields[i+1]; !ok {
					continue
				}
				outputParts = append(outputParts, part)
			}
			outputLine := strings.Join(outputParts, delimiter)
			fmt.Println(outputLine)
		case *separatedFlag && !containsDelimiter:
			continue
		case !*separatedFlag && containsDelimiter:
			outputParts := []string{}
			for i, part := range parts {
				// Если выбрано конкретное поле (колонка) или выбор не задан (выводим все поля)
				if _, ok := selectedFields[i+1]; !ok {
					continue
				}
				outputParts = append(outputParts, part)
			}
			outputLine := strings.Join(outputParts, delimiter)
			fmt.Println(outputLine)
		case !*separatedFlag && !containsDelimiter:
			fmt.Println(line)
		}

		// Если у строки есть разделитель
		//if len(parts) > 1 {
		//	// Если утилита настроена на выводить только строки с разделителем (-s), пропускаем строки без разделителя
		//	if *separatedFlag && strings.Contains(line, delimiter) {
		//		//fmt.Println(line)
		//		outputParts := []string{}
		//		for i, part := range parts {
		//			// Если выбрано конкретное поле (колонка) или выбор не задан (выводим все поля)
		//			if _, ok := selectedFields[i+1]; !ok {
		//				continue
		//			}
		//			outputParts = append(outputParts, part)
		//		}
		//		outputLine := strings.Join(outputParts, delimiter)
		//		fmt.Println(outputLine)
		//	} else if *separatedFlag {
		//
		//	}
		//
		//
		//	outputParts := []string{}
		//	for i, part := range parts {
		//		// Если выбрано конкретное поле (колонка) или выбор не задан (выводим все поля)
		//		if len(selectedFields) == 0 || selectedFields[i+1] {
		//			outputParts = append(outputParts, part)
		//		}
		//	}
		//	outputLine := strings.Join(outputParts, delimiter)
		//	fmt.Println(outputLine)
		//}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения стандартного ввода:", err)
		os.Exit(1)
	}
}

// Парсинг индекса поля (колонки) из строки вида "N" или "N-M"
func parseFieldIndex(field string) []int {
	fieldParts := strings.Split(field, "-")
	//fmt.Println(fieldParts)
	if len(fieldParts) == 1 {
		// Если только одно поле
		return []int{parseInt(fieldParts[0])}
	} else if len(fieldParts) == 2 {
		// Если диапазон полей (например, "N-M")
		start := parseInt(fieldParts[0])
		end := parseInt(fieldParts[1])
		if start <= end {
			return []int{start, end}
		}
	}
	// В случае ошибки возвращаем -1
	return nil
}

// Преобразование строки в целое число с обработкой ошибок
func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return n
}
