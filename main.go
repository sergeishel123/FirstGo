package main

import (
	"encoding/json" // для JSON ответов
	"net/http"      // HTTP -  сервер
	"time"
)

// Возвращает кол-во дней до нового года (к тому же так как функция публичная, то с большой буквы)
func DaysToNewYear(t time.Time) int {
	t2 := time.Date(t.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	delta := t2.Sub(t)
	return int(delta.Hours() / 24)
}

// Новый обработчик HTTP-запросов (Шаг 5 задания)
func apiHandler(Response http.ResponseWriter, Request *http.Request) {

	// Дату выцепляю как параметр Get-запроса date
	dateInQuery := Request.URL.Query().Get("date")
	var date time.Time
	var err error

	//	Определяем дату для расчета, если пустая, то текущую берём
	if dateInQuery == "" {
		date = time.Now()
	} else {

		// Пытаюсь распарсить дату в классическом виде

		date, err = time.Parse("2026-09-25", dateInQuery)
		if err != nil { // Я наконец понял, что функции в Go зачастую возвращают вторым результатом ошибку. Если она не пустая(не nil), то плохо

			// Если формат неверный, возвращаем и JSON с ошибкой, статус response 400
			Response.Header().Set("Content-Type", "application/json")
			Response.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(Response).Encode(map[string]string{"Ошибка": "Неправильный формат даты, используй формат  YYYY-MM-DD"})
			return

		}
	}

	days := DaysToNewYear(date)

	// Если без ошибочно, тол json - ом вовзращаем ответ со статусом 200
	Response.Header().Set("Content-Type", "application/json")
	Response.WriteHeader(http.StatusOK)

	MyResponse := map[string]interface{}{"Количество дней до Нового Года": days}
	if dateInQuery != "" {
		MyResponse["date"] = dateInQuery
	}

	json.NewEncoder(Response).Encode(MyResponse)
}

// Endpoint по идее
func main() {

	// Обработчик Get - запросов сделал по пути SergeyAndDays
	http.HandleFunc("/SergeyAndDays", apiHandler)

	// Теперь по пути http://localhost:8080/SergeyAndDays по идее можно обращаться для расчёта кол-ва дней

	// Запуск сервера локально  (порт 8080)
	http.ListenAndServe(":8080", nil)
}
