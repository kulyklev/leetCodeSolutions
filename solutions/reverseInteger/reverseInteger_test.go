package reverseInteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseInteger(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				num: 123,
			},
			want: 321,
		},
		{
			name: "Case 2",
			args: args{
				num: 120,
			},
			want: 21,
		},
		{
			name: "Case 3",
			args: args{
				num: 102,
			},
			want: 201,
		},
		{
			name: "Case 4",
			args: args{
				num: -123,
			},
			want: -321,
		},
		{
			name: "Case 5",
			args: args{
				num: -123,
			},
			want: -321,
		},
		{
			name: "Case 6",
			args: args{
				num: -2147483648,
			},
			want: 0,
		},
		{
			name: "Case 7",
			args: args{
				num: 2147483647,
			},
			want: 0,
		},
		{
			name: "Case 8",
			args: args{
				num: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseInteger(tt.args.num))
		})
	}
}
