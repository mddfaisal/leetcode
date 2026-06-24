package problem_1

import (
	"reflect"
	"testing"
)

func Test_Main(t *testing.T) {
	tests := []struct {
		name   string
		target int
		input  []int
		output []int
	}{
		{
			name:   "one",
			target: 9,
			input:  []int{2, 7, 11, 15},
			output: []int{0, 1},
		},
		{
			name:   "two",
			target: 6,
			input:  []int{3, 2, 4},
			output: []int{1, 2},
		},
		{
			name:   "three",
			target: 6,
			input:  []int{3, 3},
			output: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Main(tt.input, tt.target); !reflect.DeepEqual(got, tt.output) {
				t.Errorf("Got: %v, want: %v", got, tt.output)
			}
		})
	}
}
