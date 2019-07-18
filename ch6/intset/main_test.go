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
				t.Errorf("got: %v want: %v", got, tt.want)
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
				t.Errorf("Want %v , got %v ", tt.want, tt.s)
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
