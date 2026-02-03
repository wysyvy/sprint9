package main

import "testing"

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("ERROR, not empty")
	}

	result = generateRandomElements(-10)
	if len(result) != 0 {
		t.Errorf("ERROR, -10")
	}

	result = generateRandomElements(10)
	if len(result) != 10 {
		t.Errorf("ERROR, 10")
	}

	for _, v := range result {
		if v <= 0 {
			t.Errorf("ERROR, elem < 0")
			break
		}
	}
}

func TestMaximum(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("ERROR, not empty")
	}

	result = maximum([]int{10})
	if result != 10 {
		t.Errorf("ERROR, [10]")
	}

	result = maximum([]int{1, 2, 3, 4, 5})
	if result != 5 {
		t.Errorf("ERROR, [1 2 3 4 5]")
	}

	result = maximum([]int{100, 1, 2})
	if result != 100 {
		t.Errorf("ERROR, [100 1 2]")
	}
}

func TestMaxChunks(t *testing.T) {
	result := maxChunks([]int{})
	if result != 0 {
		t.Errorf("ERROR, not empty")
	}

	result = maxChunks([]int{10})
	if result != 10 {
		t.Errorf("ERROR, [10]")
	}

	result = maxChunks([]int{1, 2, 3, 4, 5, 6, 7, 8})
	if result != 8 {
		t.Errorf("ERROR, [1 2 3 4 5 6 7 8]")
	}

	result = maxChunks([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if result != 9 {
		t.Errorf("ERROR, [1 2 3 4 5 6 7 8 9]")
	}
}
