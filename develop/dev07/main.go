package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) chan interface{} {
	if len(channels) == 0 { // Если на входе каналов нет, вернем закрытый канал
		c := make(chan interface{})
		close(c)
		return c
	}

	c := make(chan interface{}) // Cоздадим канал для получения значений из входных каналов

	for _, ch := range channels { // Запустим подпрограмму для каждого входного канала для пересылки значений в выходной канал (output)
		go func(ch <-chan interface{}) {
			for v := range ch {
				c <- v
			}
		}(ch)
	}

	return c
}

func merge[T any](chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}

	for _, singleChan := range chans {
		wg.Add(1)
		go func(ch chan T) {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
		}(singleChan)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {

	sig := func(after time.Duration) chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()

	ch := merge(
		sig(2*time.Nanosecond),
		sig(5*time.Microsecond),
		sig(1*time.Second),
		sig(1*time.Microsecond),
		sig(1*time.Second),
	)
	fmt.Println(reflect.TypeOf(ch))
	fmt.Println(len(ch))

	fmt.Printf("fone after %v", time.Since(start))

	//for v := range ch {
	//	fmt.Println(v)
	//}
}
