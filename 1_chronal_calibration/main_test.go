package main

import (
	"testing"
)

func TestFrequencyValue_stringToInt32(t *testing.T) {
	tests := []struct {
		name string
		iv   FrequencyValue
		want int32
	}{
		{
			name: "positive num",
			iv:   FrequencyValue("1"),
			want: 1,
		},
		{
			name: "negative num",
			iv:   FrequencyValue("-1"),
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.iv.asInt32(); got != tt.want {
				t.Errorf("FrequencyValue.stringToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequencyCollection_findFrequency(t *testing.T) {
	type fields struct {
		starting int32
		values   []FrequencyValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "basic test",
			fields: fields{
				values:   []FrequencyValue{"1", "2", "-1", "3"},
				starting: 0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FrequencyCollection{
				starting: tt.fields.starting,
				values:   tt.fields.values,
			}
			if got := fc.findFrequency(); got != tt.want {
				t.Errorf("FrequencyCollection.findFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequencyCollection_findDuplicateFrequency(t *testing.T) {
	type fields struct {
		starting int32
		values   []FrequencyValue
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "basic test",
			fields: fields{
				values:   []FrequencyValue{"1", "-1", "1", "-1"},
				starting: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FrequencyCollection{
				starting: tt.fields.starting,
				values:   tt.fields.values,
			}
			if got := fc.findDuplicateFrequency(); got != tt.want {
				t.Errorf("FrequencyCollection.findDuplicateFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
