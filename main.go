package main

import (
	"fmt"
	"time"
)

var bruno person.Person

func init() {
	bruno = person.Person{
		Name:       "Bruno",
		BirthDate:  time.Date(2001, time.Month(10), 17, 0, 0, 0, 0, &time.Location{}),
		Profession: "web developer",
	}
}

func main() {
	fmt.Print(bruno)
}
