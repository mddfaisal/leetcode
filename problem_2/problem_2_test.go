package problem2

import (
	"reflect"
	"testing"
)

func getList(nums []int) *ListNode {
	res := &ListNode{}
	head := res
	for i := 0; i < len(nums); i++ {
		head.Next = &ListNode{Val: nums[i]}
		head = head.Next
	}
	return res.Next
}

func Test_AddTwoNumbers(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		output []int
	}{
		{
			name:   "one",
			list1:  []int{2, 4, 3},
			list2:  []int{5, 6, 4},
			output: []int{7, 0, 8},
		},
		{
			name:   "two",
			list1:  []int{0},
			list2:  []int{0},
			output: []int{0},
		},
		{
			name:   "three",
			list1:  []int{9, 9, 9, 9, 9, 9, 9},
			list2:  []int{9, 9, 9, 9},
			output: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := getList(tt.list1)
			if got := AddTwoNumbers(l1, getList(tt.list2)); !reflect.DeepEqual(got, getList(tt.output)) {
				t.Errorf("Error in test: %v", tt.name)
			}
		})
	}
}
