package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func FindAnagramSets(words *[]string) *map[string]*[]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		word = strings.ToLower(word)

		key := sortString(word)

		anagramSets[key] = append(anagramSets[key], word)
	}

	for key, words := range anagramSets {
		if len(words) <= 1 {
			delete(anagramSets, key)
		}
	}

	firstFindWordKeysMap := make(map[string]*[]string, len(anagramSets))
	for _, listValue := range anagramSets {
		listValue := listValue
		firstFindWordKeysMap[listValue[0]] = &listValue
	}

	return &firstFindWordKeysMap
}

func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	anagramSets := FindAnagramSets(&words)

	for key, words := range *anagramSets {
		fmt.Printf("Множество анаграмм для ключа '%s': %v\n", key, words)
	}
}
