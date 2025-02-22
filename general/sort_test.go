package general

import (
	"math/rand"
	"testing"
)

func generateSlice(b *testing.B, size int) []int {
	b.Helper()
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999999)
	}
	return slice
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := generateSlice(b, 1000)
		b.StartTimer()
		BubbleSort(input)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := generateSlice(b, 1000)
		b.StartTimer()
		SelectionSort(input)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := generateSlice(b, 1000)
		b.StartTimer()
		MergeSort(input)
	}
}
