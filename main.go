package main

import (
    "fmt"
    "time"
)

func main() {

    t1 := time.Now()
    t2 := time.Date(t1.Year() + 1, time.January, 1,0,0,0,0, time.UTC)

    delta := t2.Sub(t1)

    days := int(delta.Hours() / 24)

	fmt.Printf("До Нового года осталось %d дней", days)


}

