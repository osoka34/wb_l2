package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// Получаем точное время с NTP-сервера
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Форматируем и выводим текущее время
	currentTime := time.Now()
	fmt.Println("Текущее время:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Точное время с NTP:", ntpTime.Format("2006-01-02 15:04:05"))
}
