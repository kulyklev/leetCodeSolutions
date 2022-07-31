package containerWithMostWater

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_maxArea(t *testing.T) {
	f, err := os.Open("data.txt")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	height, err := readInts(f)

	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		height []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "Case 2",
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "Case 3",
			args: args{
				height: height,
			},
			want: 705634720,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.height))
		})
	}
}

func Benchmark_maxArea(b *testing.B) {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	height, err := readInts(f)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	var text string

	for scanner.Scan() {
		text = scanner.Text()
		text = strings.Replace(text, ",", "", -1)
		x, err := strconv.Atoi(text)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, scanner.Err()
}
