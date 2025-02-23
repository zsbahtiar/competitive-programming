package general

import (
	"fmt"
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	var tests = []struct {
		arg struct {
			nums   []int
			search int
		}
		want bool
	}{
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{32, 27, 43, 3, 9, 82, 10},
				search: 3,
			},
			want: true,
		},
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{6, 4, 2, -2, 10, 5, 0, 1, -5, 7},
				search: 5,
			},
			want: true,
		},
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{6, 4, 2, -2, 10, 5, 0, 1, -5, 7},
				search: 8,
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := SequentialSearch(test.arg.nums, test.arg.search)
		if got != test.want {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		arg struct {
			nums   []int
			search int
		}
		want bool
	}{
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{32, 27, 43, 3, 9, 82, 10},
				search: 3,
			},
			want: true,
		},
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{6, 4, 2, -2, 10, 5, 0, 1, -5, 7},
				search: 5,
			},
			want: true,
		},
		{
			arg: struct {
				nums   []int
				search int
			}{
				nums:   []int{6, 4, 2, -2, 10, 5, 0, 1, -5, 7},
				search: 8,
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := BinarySearch(test.arg.nums, test.arg.search)
		if got != test.want {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}

func BenchmarkSequentialSearch(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			nums := generateSlice(b, size)
			target := nums[size/2]

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				SequentialSearch(nums, target)
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			nums := generateSlice(b, size)
			target := nums[size/2]

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BinarySearch(nums, target)
			}
		})
	}
}
