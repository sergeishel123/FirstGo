package main

import (
	"fmt"
	"testing"
	"time"
)

// t — это имя переменной (используют t для тестов)
// * - классический указатель (хранит адрес переменной)

// Тестирую функцию  DaysToNewYear(t1 time.Time)

func TestDaysToNewYear(t *testing.T) {

	// В Golan [] - это срез(динамический массив или список), tests - список структур

	tests := []struct {
		id       int
		date     time.Time // аргумент t
		expected int       // Возвращаемое значение функции
	}{
		{
			id:       1, // 1 января
			date:     time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			id:       2, //  31 декабря
			date:     time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			id:       3, // 1 января
			date:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 366, // 2024 високосный
		},
		{
			id:       4, // 1 марта
			date:     time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			expected: 306,
		},
		{
			id:       5, // 1 июля
			date:     time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC),
			expected: 184,
		},
	}

	// Проходим по всем
	// с помощью range итерируемся по коллекциям - возвращает пару из индексов элементов и самих элементов
	// Причём была проблема явно - если не используешь вот индекс элемента далее в коде,
	// то выдает исключение, поэтому если не обращаешься к нему, то для переменной для индекса в цикле использовать _
	for _, cur_test := range tests {
		t.Run(fmt.Sprintf("Тест номер %d", cur_test.id), func(t *testing.T) {
			current_answer := DaysToNewYear(cur_test.date)
			if current_answer != cur_test.expected {
				t.Errorf("DaysToNewYear(%v) = %d; а ожидается %d", cur_test.date, current_answer, cur_test.expected)
				// % специфиатор v в принципе выводит любой объект
			}
		})
	}
}
