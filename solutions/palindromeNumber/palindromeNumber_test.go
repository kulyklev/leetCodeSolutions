package palindromeNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				num: 121,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				num: -121,
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				num: 1234,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.num))
		})
	}
}
