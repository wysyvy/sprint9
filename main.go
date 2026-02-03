package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}

	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	chunks := 8
	if len(data) < chunks {
		chunks = len(data)
	}

	chunkSize := len(data) / chunks
	maxValues := make([]int, chunks)
	var wg sync.WaitGroup

	for i := 0; i < chunks; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize

		if i == chunks-1 {
			end = len(data)
		}

		go func(i int) {
			defer wg.Done()
			chunk := data[start:end]
			maxValues[i] = maximum(chunk)
		}(i)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	rand.Seed(time.Now().UnixNano())
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mks\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mks\n", max, elapsed)
}
