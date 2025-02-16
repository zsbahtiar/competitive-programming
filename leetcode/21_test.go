package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	var cases = []struct {
		name  string
		input struct {
			list1 *ListNode
			list2 *ListNode
		}
		want *ListNode
	}{
		{
			name: "Example 1: Two non-empty lists",
			input: struct {
				list1 *ListNode
				list2 *ListNode
			}{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Example 2: Both empty lists",
			input: struct {
				list1 *ListNode
				list2 *ListNode
			}{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "Example 3: One empty list",
			input: struct {
				list1 *ListNode
				list2 *ListNode
			}{
				list1: nil,
				list2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := MergeTwoSortedLists(c.input.list1, c.input.list2)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%s: got different result than expected", c.name)
			}
		})
	}
}
