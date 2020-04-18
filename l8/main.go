//Домашнее задание
//Параллельное исполнение
//Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).
//
//Функция принимает на вход:
//- слайс с заданиями `[]func() error`;
//- число заданий которые можно выполнять параллельно (`N`);
//- максимальное число ошибок после которого нужно приостановить обработку.

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func runGoroutines(funcs []func() error, goroutininesCount int, errorCount int) {
	chErr := make(chan error, errorCount)
	parallelTasks := make(chan struct{}, goroutininesCount)

	for i, func_ := range funcs {
		go func(innerFunc func() error, i int) {
			parallelTasks <- struct{}{}
			err := innerFunc()
			<-parallelTasks
			fmt.Printf("%v %v\n", i, err)
			chErr <- err
		}(func_, i)
	}

	for i := 0; i < errorCount; i++ {
		<-chErr
	}
}

func randomThirteen() error {
	rand.Seed(time.Now().Unix() % int64(time.Now().Nanosecond()))
	for true {
		if rand.Intn(100000000) == 13 {
			return errors.New("have 13")
		}
	}
	return nil
}

func randomSixtySix() error {
	rand.Seed(time.Now().Unix() % int64(time.Now().Nanosecond()))
	for true {
		if rand.Intn(100000000) == 66 {
			return errors.New("have 66")
		}
	}
	return nil
}

func randomSixtyNine() error {
	rand.Seed(time.Now().Unix() % int64(time.Now().Nanosecond()))
	for true {
		if rand.Intn(100000000) == 69 {
			return errors.New("have 69")
		}
	}
	return nil
}
func randomNintyNine() error {
	rand.Seed(time.Now().Unix() % int64(time.Now().Nanosecond()))
	for true {
		if rand.Intn(100000000) == 99 {
			return errors.New("have 99")
		}
	}
	return nil
}

type SomeFuncType []func() error

func main() {
	var funcsForRun SomeFuncType
	funcsForRun = append(funcsForRun, randomThirteen, randomSixtyNine, randomSixtySix, randomNintyNine)
	runGoroutines(funcsForRun, 2, 4)
}