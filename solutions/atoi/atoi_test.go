package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				str: "123",
			},
			want: 123,
		},
		{
			name: "Case 2",
			args: args{
				str: "-123",
			},
			want: -123,
		},
		{
			name: "Case 3",
			args: args{
				str: "    123",
			},
			want: 123,
		},
		{
			name: "Case 4",
			args: args{
				str: "   -123",
			},
			want: -123,
		},
		{
			name: "Case 5",
			args: args{
				str: "0032",
			},
			want: 32,
		},
		{
			name: "Case 6",
			args: args{
				str: "qwerty",
			},
			want: 0,
		},
		{
			name: "Case 7",
			args: args{
				str: "1234 qwerty",
			},
			want: 1234,
		},
		{
			name: "Case 8",
			args: args{
				str: "-1234 qwerty",
			},
			want: -1234,
		},
		{
			name: "Case 9",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "Case 10",
			args: args{
				str: "91283472332",
			},
			want: 2147483647,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, myAtoi(tt.args.str))
		})
	}
}
