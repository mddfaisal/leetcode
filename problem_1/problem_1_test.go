package problem_1

import (
	"reflect"
	"testing"
)

func Test_Main(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Main(test.input); !reflect.DeepEqual(got, test.output) {
				t.Errorf("Got: %v, want: %v", got, test.output)
			}
		})
	}
}
