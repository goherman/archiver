package shannon_fano

import (
	"reflect"
	"testing"
)

func Test_bestDividerPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "Base name1",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "Base name2",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "Base name3",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "Base name4",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "Base name5",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "Base name",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "Base name2",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
		{
			name: "Base name3",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
