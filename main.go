package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//fmt.Println("Первое задание")
	//FirstTask(10)
	//fmt.Println("Второе задание")
	//SecondTask(10000, 2)
	//fmt.Println("Третье задание")
	//ThirdTask()
	//fmt.Println("Четвертое задание")
	//FourthTask()
	fmt.Println("Пятое задание")
	FifthTask(10, 4)
}

func FirstTask(n int) {
	t := time.Now()
	fmt.Println("Время начала", t)
	vector := make([]int, n)

	for i := 0; i < n; i++ {
		vector[i] = rand.Intn(n)
	}

	//fmt.Println("искоходный вектор: ", vector)

	multiplier := 2

	for i := 0; i < len(vector); i++ {
		vector[i] *= multiplier
	}

	//fmt.Println("обработанный вектор: ", vector)
	elapsedTime := time.Since(t)
	fmt.Println("Прошло времени: ", elapsedTime)
}

func SecondTask(n, m int) {
	t := time.Now()
	fmt.Println(t)
	multiplier := 2
	vector := make([]int, n)

	for i := 0; i < n; i++ {
		vector[i] = rand.Intn(n)
	}

	//fmt.Println("Исходный вектор: ", vector)

	var wg sync.WaitGroup

	partSize := n / m
	for i := 0; i < m; i++ {
		wg.Add(1)
		start := i * partSize
		end := (i + 1) * partSize
		if i == m-1 {
			end = n
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				vector[j] *= multiplier
			}
		}(start, end)
	}

	wg.Wait()

	//fmt.Println("Обработанный вектор:", vector)

	elapsedTime := time.Since(t)
	fmt.Println("Прошло времени: ", elapsedTime)
}

func processVector(vector []int, multiplier, m int, t time.Time) time.Duration {
	n := len(vector)
	partSize := n / m
	var wg sync.WaitGroup

	for i := 0; i < m; i++ {
		wg.Add(1)
		start := i * partSize
		end := (i + 1) * partSize
		if i == m-1 {
			end = n
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				vector[j] *= multiplier
			}
		}(start, end)
	}

	wg.Wait()
	elapsedTime := time.Since(t)
	return time.Duration(elapsedTime.Microseconds())
}

func ThirdTask() {
	startTime := time.Now()
	NValues := []int{10, 100, 1000, 100000}
	MValues := []int{2, 3, 4, 5, 10}

	for _, n := range NValues {
		for _, m := range MValues {
			vector := make([]int, n)

			for i := 0; i < n; i++ {
				vector[i] = i + 1
			}

			elapsedTime := processVector(vector, 2, m, startTime)

			fmt.Printf("%d\t%d\t%s\n", n, m, elapsedTime)
		}
	}
}

func FourthTask() {
	startTime := time.Now()
	NValues := []int{10, 100, 1000, 100000}
	MValues := []int{2, 3, 4, 5, 10}

	for _, n := range NValues {
		for _, m := range MValues {
			vector := make([]int, n)

			for i := 0; i < n; i++ {
				vector[i] = int(math.Sin(float64((i + 1) * (i + 1))))
			}

			elapsedTime := processVector(vector, 2, m, startTime)

			fmt.Printf("%d\t%d\t%s\n", n, m, elapsedTime)
		}
	}
}

func FifthTask(n, m int) {
	t := time.Now()
	fmt.Println(t)
	multiplier := 2
	vector := make([]int, n)

	for i := 0; i < n; i++ {
		vector[i] = rand.Intn(n)
	}

	fmt.Println("Исходный вектор: ", vector)

	var wg sync.WaitGroup

	partSize := 1
	i := 0 // Индекс текущего элемента вектора

	for i < n {
		wg.Add(1)
		start := i
		end := i + partSize
		if end > n {
			end = n
		}
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				vector[j] *= multiplier
			}
		}(start, end)

		// Увеличиваем partSize для следующей горутины
		partSize++
		i = end
	}

	wg.Wait()

	fmt.Println("Обработанный вектор:", vector)

	elapsedTime := time.Since(t)
	fmt.Println("Прошло времени: ", elapsedTime)
}
