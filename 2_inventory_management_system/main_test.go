package main

import (
	"reflect"
	"testing"
)

func TestBoxID_countDuplicates(t *testing.T) {
	tests := []struct {
		name string
		b    BoxID
		want BoxMetaCount
	}{
		{
			name: "happy path, unique 2 and 3",
			b:    BoxID("ababbc"),
			want: BoxMetaCount{hasTwo: true, hasThree: true},
		},
		{
			name: "happy path, just 2",
			b:    BoxID("abadc"),
			want: BoxMetaCount{hasTwo: true, hasThree: false},
		},
		{
			name: "happy path, just 3",
			b:    BoxID("bfbgb"),
			want: BoxMetaCount{hasTwo: false, hasThree: true},
		},
		{
			name: "happy path, no dupes",
			b:    BoxID("abcdef"),
			want: BoxMetaCount{hasTwo: false, hasThree: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.countDuplicates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoxID.countDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoxCollection_calculateChecksum(t *testing.T) {
	type fields struct {
		values []BoxID
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "happy path",
			fields: fields{
				values: []BoxID{"ababa", "bbasd", "rawr"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BoxCollection{
				values: tt.fields.values,
			}
			if got := bc.calculateChecksum(); got != tt.want {
				t.Errorf("BoxCollection.calculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoxID_checkDiffChars(t *testing.T) {
	type args struct {
		c BoxID
	}
	tests := []struct {
		name  string
		b     BoxID
		args  args
		want  bool
		want1 int
	}{
		{
			name:  "happy path - match",
			b:     BoxID("ababfb"),
			args:  args{BoxID("ababzb")},
			want:  true,
			want1: 4,
		},
		{
			name:  "happy path - no math",
			b:     BoxID("abzbfb"),
			args:  args{BoxID("ababzb")},
			want:  false,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.b.checkDiffChars(&tt.args.c)
			if got != tt.want {
				t.Errorf("BoxID.checkDiffChars() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BoxID.checkDiffChars() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBoxCollection_findSingleDiffBox(t *testing.T) {
	type fields struct {
		values []BoxID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy path",
			fields: fields{
				values: []BoxID{"aaaba", "aaala"},
			},
			want: "aaaa",
		},
		{
			name: "no match",
			fields: fields{
				values: []BoxID{"abcde", "fdsae"},
			},
			want: "ERROR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BoxCollection{
				values: tt.fields.values,
			}
			if got := bc.findSingleDiffBox(); got != tt.want {
				t.Errorf("BoxCollection.findSingleDiffBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
