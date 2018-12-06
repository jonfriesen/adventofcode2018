package main

import (
	"testing"
)

func Test_runesAreSameChar(t *testing.T) {
	type args struct {
		a rune
		b rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"check a", args{'a', 'A'}, true},
		{"check z", args{'Z', 'z'}, true},
		{"check a", args{'a', 'Z'}, false},
		{"check a", args{'e', 'e'}, true},
		{"check a", args{'A', 'c'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runesAreSameChar(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("runesAreSameChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unitReacts(t *testing.T) {
	type args struct {
		a rune
		b rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"check a", args{'a', 'A'}, true},
		{"check z", args{'Z', 'z'}, true},
		{"check a", args{'a', 'Z'}, false},
		{"check a", args{'e', 'e'}, false},
		{"check Ac", args{'A', 'c'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unitsReact(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("unitReacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeRangeFromString(t *testing.T) {
	type args struct {
		s  string
		st int
		e  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"happy path", args{"abcdefg", 2, 3}, "abefg"},
		{"end of path", args{"abcdefg", 5, 6}, "abcde"},
		{"beginning of path", args{"abcdefg", 0, 2}, "defg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeRangeFromString(tt.args.s, tt.args.st, tt.args.e); got != tt.want {
				t.Errorf("removeRangeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processPolymer(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"happy path", args{"abcCdefg"}, "abdefg"},
		{"happy path", args{"aDcCdefg"}, "aefg"},
		{"sample input", args{"dabAcCaCBAcCcaDA"}, "dabCBAcaDA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processPolymer(tt.args.p); got != tt.want {
				t.Errorf("processPolymer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLower(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"test A", args{'A'}, 97},
		{"test a", args{'a'}, 97},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLower(tt.args.r); got != tt.want {
				t.Errorf("toLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
