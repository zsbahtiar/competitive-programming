package general

import (
	"math/rand"
	"reflect"
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

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{32, 27, 43, 3, 9, 82, 10},
			want: []int{3, 9, 10, 27, 32, 43, 82},
		},
		{
			arg:  []int{11, 6, 3, 24, 46, 22, 7},
			want: []int{3, 6, 7, 11, 22, 24, 46},
		},
	}

	for _, test := range tests {
		got := BubbleSort(test.arg)
		if !reflect.DeepEqual(&got, &test.want) {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{32, 27, 43, 3, 9, 82, 10},
			want: []int{3, 9, 10, 27, 32, 43, 82},
		},
		{
			arg:  []int{11, 6, 3, 24, 46, 22, 7},
			want: []int{3, 6, 7, 11, 22, 24, 46},
		},
	}

	for _, test := range tests {
		got := MergeSort(test.arg)
		if !reflect.DeepEqual(&got, &test.want) {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{32, 27, 43, 3, 9, 82, 10},
			want: []int{3, 9, 10, 27, 32, 43, 82},
		},
		{
			arg:  []int{11, 6, 3, 24, 46, 22, 7},
			want: []int{3, 6, 7, 11, 22, 24, 46},
		},
	}

	for _, test := range tests {
		got := SelectionSort(test.arg)
		if !reflect.DeepEqual(&got, &test.want) {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}
