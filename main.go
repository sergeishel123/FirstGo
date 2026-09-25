package main

import (
	"fmt"
	"time"
)

// Возвращает кол-во дней до нового года (к тому же так как функция публичная, то с большой буквы)
func DaysToNewYear(t time.Time) int {
	t2 := time.Date(t.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	delta := t2.Sub(t)
	return int(delta.Hours() / 24)
}

func main() {

	t := time.Now()

	days := DaysToNewYear(t)

	fmt.Printf("До Нового года осталось %d дней\n", days)
}
