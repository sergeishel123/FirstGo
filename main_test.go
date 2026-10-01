package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
			id:       1,
			date:     time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			id:       2,
			date:     time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			id:       3,
			date:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 366, // 2024 високосный
		},
		{
			id:       4,
			date:     time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			expected: 306,
		},
		{
			id:       5,
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

//  тестирование HTTP API

type apiSuccessResponse struct {
	Days int    `json:"Количество дней до Нового Года"`
	Date string `json:"date,omitempty"`
}

// для разбора JSON-ответа с ошибкой
type apiErrorResponse struct {
	Error string `json:"Ошибка"`
}

func TestHTTPHandler(t *testing.T) {

	tests := []struct {
		name           string
		query          string // сразу имею в виду аргументы GET - зароса
		expectedStatus int    // Ожидаемый HTTP-код
		expectedDays   int    // Ожидаемое количество дней
		expectError    bool
	}{
		{
			name:           "Вообще дату не передаем в запросе",
			query:          "",
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "Запрос с високосным годом",
			query:          "date=2024-01-01",
			expectedStatus: http.StatusOK,
			expectedDays:   366,
			expectError:    false,
		},
		{
			name:           "невалидной датой",
			query:          "date=hello",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			url := "/SergeyAndDays"
			if tt.query != "" {
				url += "?" + tt.query
			}
			request := httptest.NewRequest(http.MethodGet, url, nil)

			// регистратор ответа
			w := httptest.NewRecorder()

			apiHandler(w, request)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Ожидаемый статус %d,а  получен %d", tt.expectedStatus, res.StatusCode)
			}

			// Если ожидаеся ошибка, то по-хорошему надо, что в теле ответа есть поле Ошибка
			if tt.expectError {
				var errResp apiErrorResponse
				err := json.NewDecoder(res.Body).Decode(&errResp)
				if err != nil {

					t.Fatalf("Не удалось разобрать JSON ошибки: %v", err)
				}
				if errResp.Error == "" {
					t.Errorf("Ожидалось сообщение об ошибке, но поле Ошибка пустое")
				}
				return
			}

			//полученный JSON
			var resp apiSuccessResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			if err != nil {
				t.Fatalf("Не удалось разобрать JSON ответа: %v", err)
			}

			if resp.Days != tt.expectedDays {
				t.Errorf("Ожидаемое кол-во дней %d, а получено %d", tt.expectedDays, resp.Days)
			}
		})
	}
}
