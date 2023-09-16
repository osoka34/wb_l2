package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortOptions struct {
	keyColumn   int
	numericSort bool
	reverseSort bool
	uniqueSort  bool
}

type Line struct {
	Text   string
	Fields []string
}

// go run main.go -i input.txt -o output.txt -u -r -n
func main() {
	inputFileName := flag.String("i", "", "Имя входного файла")
	outputFileName := flag.String("o", "", "Имя выходного файла")
	keyColumn := flag.Int("k", -1, "Номер колонки для сортировки (по умолчанию -1 для сортировки всей строки)")
	numericSort := flag.Bool("n", false, "Сортировать по числовому значению")
	reverseSort := flag.Bool("r", false, "Сортировать в обратном порядке")
	uniqueSort := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	if *inputFileName == "" {
		fmt.Println("Необходимо указать имя входного файла с флагом -i")
		os.Exit(1)
	}

	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Println("Ошибка при открытии входного файла:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var lines []Line
	for scanner.Scan() {
		lineText := scanner.Text()
		lineFields := strings.Fields(lineText)
		lines = append(lines, Line{Text: lineText, Fields: lineFields})
	}

	if scanner.Err() != nil {
		fmt.Println("Ошибка при чтении входного файла:", scanner.Err())
		os.Exit(1)
	}

	//sortOptions := SortOptions{*keyColumn, *numericSort, *reverseSort, *uniqueSort}

	sortLines(lines, *keyColumn-1, *numericSort)

	if *reverseSort {
		reverse(lines)
	}

	if *uniqueSort {
		lines = removeDuplicates(lines, *keyColumn-1)
	}

	if *outputFileName == "" {
		for _, line := range lines {
			fmt.Println(line.Text)
		}
	} else {
		outputFile, err := os.Create(*outputFileName)
		if err != nil {
			fmt.Println("Ошибка при создании выходного файла:", err)
			os.Exit(1)
		}
		defer outputFile.Close()

		writer := bufio.NewWriter(outputFile)
		for _, line := range lines {
			_, err := writer.WriteString(line.Text + "\n")
			if err != nil {
				fmt.Println("Ошибка при записи в выходной файл:", err)
				os.Exit(1)
			}
		}

		writer.Flush()
	}
}

func removeDuplicates(lines []Line, keyColumn int) []Line {
	seen := make(map[string]bool)
	var uniqueLines []Line

	switch {
	case keyColumn > 0:
		for _, line := range lines {
			if !seen[line.Fields[keyColumn]] {
				uniqueLines = append(uniqueLines, line)
				seen[line.Fields[keyColumn]] = true
			}
		}
	default:
		for _, line := range lines {
			if !seen[line.Text] {
				uniqueLines = append(uniqueLines, line)
				seen[line.Text] = true
			}
		}

	}

	return uniqueLines
}

func reverse(lines []Line) {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func numericSorting(lines []Line, keyColumn int) {
	if keyColumn > 0 {
		sort.SliceStable(lines, func(i, j int) bool {
			value1, err1 := strconv.ParseFloat(lines[i].Fields[keyColumn], 64)
			value2, err2 := strconv.ParseFloat(lines[j].Fields[keyColumn], 64)
			if err1 == nil && err2 == nil {
				return value1 < value2
			} else {
				return strings.Compare(lines[i].Fields[keyColumn], lines[j].Fields[keyColumn]) < 0
			}
		})
	} else {
		sort.SliceStable(lines, func(i, j int) bool {
			value1, err1 := strconv.ParseFloat(lines[i].Fields[1], 64)
			value2, err2 := strconv.ParseFloat(lines[j].Fields[1], 64)
			if err1 == nil && err2 == nil {
				return value1 < value2
			} else {
				return strings.Compare(lines[i].Text, lines[j].Text) < 0
			}
		})
	}
}

func sortLines(lines []Line, keyColumn int, numericSort bool) {
	switch {
	case numericSort:
		numericSorting(lines, keyColumn)
	case !numericSort && keyColumn > 0:
		sort.SliceStable(lines, func(i, j int) bool {
			return strings.Compare(lines[i].Fields[keyColumn], lines[j].Fields[keyColumn]) < 0
		})
	default:
		sort.SliceStable(lines, func(i, j int) bool {
			return strings.Compare(lines[i].Text, lines[j].Text) < 0
		})
	}
}
