package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	// Тестовые входные данные
	tests := []struct {
		words        []string
		expectedSets *map[string]*[]string
	}{
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			&map[string]*[]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"hello", "world", "olleh", "dlrow"},
			&map[string]*[]string{
				"hello": {"hello", "world", "olleh", "dlrow"},
			},
		},
		{
			[]string{"apple", "banana", "cherry"},
			&map[string]*[]string{},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := FindAnagramSets(&test.words)

			if !reflect.DeepEqual(result, test.expectedSets) {
				t.Errorf("Ожидалось %v, но получено %v", test.expectedSets, result)
			}
		})
	}
}
