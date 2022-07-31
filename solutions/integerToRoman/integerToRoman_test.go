package integerToRoman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{num: 1},
			want: "I",
		},
		{
			name: "Case 2",
			args: args{num: 2},
			want: "II",
		},
		{
			name: "Case 3",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "Case 4",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "Case 5",
			args: args{num: 5},
			want: "V",
		},
		{
			name: "Case 6",
			args: args{num: 6},
			want: "VI",
		},
		{
			name: "Case 7",
			args: args{num: 7},
			want: "VII",
		},
		{
			name: "Case 8",
			args: args{num: 8},
			want: "VIII",
		},
		{
			name: "Case 9",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "Case 10",
			args: args{num: 10},
			want: "X",
		},
		{
			name: "Case 11",
			args: args{num: 11},
			want: "XI",
		},
		{
			name: "Case 12",
			args: args{num: 19},
			want: "XIX",
		},
		{
			name: "Case 13",
			args: args{num: 20},
			want: "XX",
		},
		{
			name: "Case 14",
			args: args{num: 40},
			want: "XL",
		},
		{
			name: "Case 15",
			args: args{num: 90},
			want: "XC",
		},
		{
			name: "Case 16",
			args: args{num: 400},
			want: "CD",
		},
		{
			name: "Case 17",
			args: args{num: 900},
			want: "CM",
		},
		{
			name: "Case 18",
			args: args{num: 901},
			want: "CMI",
		},
		{
			name: "Case 19",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "Case 20",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
		{
			name: "Case 21",
			args: args{num: 3999},
			want: "MMMCMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, intToRoman(tt.args.num))
		})
	}
}

func Benchmark_intToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(3999)
	}
}
