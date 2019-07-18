package main

import (
	"testing"
)

func TestIntSet_Has(t *testing.T) {
	var s IntSet
	type args struct {
		x int
	}
	tests := []struct {
		name string
		s    *IntSet
		args args
		want bool
	}{
		{"Test Empty Set for -1", &s, args{-1}, false},
		{"Test Empty Set for 1", &s, args{1}, false},
		{"Test Empty Set for 65", &s, args{65}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Has(tt.args.x); got != tt.want {
				t.Errorf("IntSet.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Add(t *testing.T) {
	var s IntSet

	type args struct {
		x int
	}
	tests := []struct {
		name string
		s    *IntSet
		args args
		want bool
	}{
		{"Test Empty Set for -1", &s, args{-1}, false},
		{"Test Empty Set for 1", &s, args{1}, true},
		{"Test Empty Set for 65", &s, args{65}, true},
		{"Test Empty Set for 144", &s, args{144}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.x)
			if got := tt.s.Has(tt.args.x); got != tt.want {
				t.Errorf("TestIntSet_Add() got: %v want: %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	var s1 IntSet
	var s2, s3 IntSet
	s1.Add(1)
	s1.Add(144)
	s2.Add(9)
	s3.Add(4)
	type args struct {
		t *IntSet
	}
	tests := []struct {
		name string
		s    *IntSet
		args args
		want string
	}{
		{"Union of s1 and s2", &s1, args{&s2}, "{1 9 144}"},
		{"Union of s2 and s1", &s2, args{&s1}, "{1 9 144}"},
		{"Union of s1 and s3", &s2, args{&s3}, "{1 4 9 144}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.UnionWith(tt.args.t)
			if tt.s.String() != tt.want {
				t.Errorf("TestIntSet_UnionWith() got %v, Want %v ", tt.s, tt.want)
			}
		})
	}
}

func TestIntSet_String(t *testing.T) {
	var s1, s2 IntSet
	s1.Add(1)
	s1.Add(2)
	s1.Add(4)
	s2.Add(1)
	s2.Add(2)
	s2.Add(4)
	s2.Add(16)
	s2.Add(32)

	tests := []struct {
		name string
		s    *IntSet
		want string
	}{
		{"Print Set with values 1,2, 4", &s1, "{1 2 4}"},
		{"Print Set with values 1,2, 4, 16, 32", &s2, "{1 2 4 16 32}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("IntSet.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Len(t *testing.T) {
	var s IntSet
	tests := []struct {
		name    string
		s       *IntSet
		wantLen int
	}{
		{"Len of empty set", &s, 0},
		{"Len of empty set", &s, 1},
		{"Len of empty set", &s, 2},
		{"Len of empty set", &s, 3},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLen := tt.s.Len(); gotLen != tt.wantLen {
				t.Errorf("IntSet_Len() = %v, want %v", gotLen, tt.wantLen)
			}
		})
		s.Add(i)
	}
}

func TestIntSet_Remove(t *testing.T) {
	var s IntSet
	s.Add(200)

	type args struct {
		x int
	}
	tests := []struct {
		name string
		s    *IntSet
		args args
		want string
	}{
		{"Remove non existent element", &s, args{1}, "{200}"},
		{"Remove on outside range element", &s, args{512}, "{200}"},
		{"Remove element", &s, args{200}, "{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.x)
			if got := tt.s.String(); got != tt.want {
				t.Errorf("IntSet_Remove(): %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Clear(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(100)

	tests := []struct {
		name string
		s    *IntSet
	}{
		{"Clear empty Set", &IntSet{}},
		{"Clear populate Set", &s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if got := tt.s.String(); got != "{}" {
				t.Error("IntSet_Clear, set not cleared")
			}
		})
	}
}

func TestIntSet_Copy(t *testing.T) {
	var s IntSet
	var s2 IntSet
	s2.Add(100)

	tests := []struct {
		name string
		s    *IntSet
		want *IntSet
	}{
		{"Copy Empty Set", &s, &IntSet{}},
		{"Copy filled Set", &s2, &s2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Copy(); got.String() != tt.s.String() || &got.words == &tt.s.words {
				t.Errorf("IntSet.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_AddAll(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		s    IntSet
		args args
		want string
	}{
		{"Adding 1-5 to an empty set", IntSet{}, args{[]int{1, 2, 3, 4, 5}}, "{1 2 3 4 5}"},
		{"Adding five and one large values to an empty set", IntSet{}, args{[]int{1, 2, 3, 4, 5, 222}}, "{1 2 3 4 5 222}"},
		{"Adding five and one x-large values to an empty set", IntSet{}, args{[]int{1, 2, 3, 4, 5, 684}}, "{1 2 3 4 5 684}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddAll(tt.args.vals...)
			if got := tt.s.String(); got != string(tt.want) {
				t.Errorf("IntSet.AddAll() = %v, want %v", got, string(tt.want))
			}
		})
	}
}
