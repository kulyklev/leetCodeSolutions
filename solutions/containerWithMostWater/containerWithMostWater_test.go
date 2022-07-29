package containerWithMostWater

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}

	testID := 0

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
				height: []int{},
			},
			want: 705634720,
		},
	}

	t.Logf("\t\tTest %d:\tWhen finding container with most water.", testID)
	{
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := maxArea1(tt.args.height); got != tt.want {
					t.Errorf("\t%s\tTest %d:\tmaxArea() =: %v, want: %v", failed, testID, got, tt.want)
				} else {
					t.Logf("\t%s\tTest %d:\tFound container with most water.", success, testID)
				}
			})
		}
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
