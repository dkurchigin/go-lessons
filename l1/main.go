//Домашнее задание
//Hello now()
//Завести Go репозиторий на GitHub, написать программу печатающую текущее время / точное время с использованием библиотеки NTP.
//Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода.

package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	time_, err := ntp.Time("ntp0.NL.net")
	currentTime := time.Now()

	if err != nil {
		fmt.Println("Something fu**ed up")
	} else {
		fmt.Println("Время компьютера и нтп-время")
		fmt.Printf("%s | %s", currentTime, time_)
	}
}