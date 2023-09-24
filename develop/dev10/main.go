package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//india.colorado.edu 13 (Get the time)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan bool)

	// Отправка ввода пользователя в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			_, err := fmt.Fprintf(conn, input+"\n")
			if err != nil {
				fmt.Println("Error sending data:", err)
				break
			}
		}
		done <- true
	}()

	// Чтение данных из сокета и вывод их на STDOUT
	go func() {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Connection closed by server.")
				break
			}
			fmt.Print(line)
		}
		done <- true
	}()

	// Обработка завершения программы
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
	case <-signals:
		fmt.Println("Exiting...")
	}
}
