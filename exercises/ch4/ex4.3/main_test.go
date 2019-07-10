package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Reverse, even number of elements",
			args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			"Reverse two elements ",
			args{[]int{0, 1}},
			[]int{1, 0},
		},
		{
			"Reverse single element",
			args{[]int{1}},
			[]int{1},
		},
		{
			"Reverse empty arry",
			args{[]int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(&tt.args.s)
			for i, v := range tt.want {
				if v != tt.args.s[i] {
					t.Errorf("want[%d]: %d != %d", i, v, tt.args.s[i])
				}
			}
		})
	}
}
