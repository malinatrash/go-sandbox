package concurrency

import (
	"sync"
	"time"
)

func WaitGroupExample() {
	//withoutWait()
	//withWait()
	//writeWithoutConcurrent()
	//writeWithoutMutex()
	//writeWithMutex()
	readWithMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go println(i + 1)
	}
	println("exit")
}

func withWait() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			println(i + 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	println("exit")
}

func writeWithoutConcurrent() {
	start := time.Now()
	counter := 0

	for counter < 1000 {
		time.Sleep(time.Nanosecond)
		counter++
	}

	println(counter)
	println(time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	wg.Wait()

	println(counter)
	println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	println(counter)
	println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	println(counter)
	println(time.Now().Sub(start).Seconds())
}
