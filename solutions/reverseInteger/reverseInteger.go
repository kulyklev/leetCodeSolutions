package reverseInteger

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/logrusorgru/aurora/v3"
	errors "github.com/pkg/errors"
)

var ErrOverflow = errors.New("integer overflow")

type ReverseIntCmd struct {
	description string
	fs          *flag.FlagSet
	helpFlag    bool
}

func NewReverseIntCmd() *ReverseIntCmd {
	reverseIntegerCmd := &ReverseIntCmd{
		description: "https://leetcode.com/problems/reverse-integer/",
		fs:          flag.NewFlagSet("reverseInt", flag.ContinueOnError),
	}

	// Here we specify flags
	reverseIntegerCmd.fs.BoolVar(&reverseIntegerCmd.helpFlag, "h", false, "Display help for the given command. When no command is given display help for the list command")

	//Custom usage
	reverseIntegerCmd.fs.Usage = func() {

		// Options section
		optionsTable := simpletable.New()

		reverseIntegerCmd.fs.VisitAll(func(f *flag.Flag) {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("-%s", f.Name)).String()},
				{Align: simpletable.AlignLeft, Text: f.Usage},
			}

			optionsTable.Body.Cells = append(optionsTable.Body.Cells, r)
		})

		optionsTable.SetStyle(simpletable.StyleCompactLite)

		// Printing help
		fmt.Println(aurora.BrightRed("Description:"))
		fmt.Println("\t", reverseIntegerCmd.description)
		fmt.Println(aurora.BrightRed("Usage:"))
		fmt.Println("\treverseInt [options]")
		fmt.Println(aurora.BrightRed("Options:"))
		fmt.Println(optionsTable.String())
	}

	return reverseIntegerCmd
}

func (r ReverseIntCmd) Init(args []string) error {
	return r.fs.Parse(args)
}

func (r ReverseIntCmd) Run() error {
	if r.helpFlag {
		r.fs.Usage()
		os.Exit(0)
	}

	var s string
	fmt.Println("Enter string to convert it to int: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		return err
	}

	return callReverseInteger(s)
}

func (r ReverseIntCmd) Name() string {
	return r.fs.Name()
}

func (r ReverseIntCmd) Description() string {
	return r.description
}

func callReverseInteger(s string) error {
	num, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	res := reverseInteger(num)
	fmt.Println("Result: ")
	fmt.Println(res)

	return nil
}

func reverseInteger(x int) int {

	var err error
	var t []int
	res := 0
	for x != 0 {
		t = append(t, x%10)
		x /= 10
	}

	for i, elem := range t {
		res, err = add32(res, elem*int(math.Pow10(len(t)-i-1)))

		if err != nil {
			return 0
		}
	}

	return res
}

func add32(left, right int) (int, error) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, ErrOverflow
		}
	} else {
		if left < math.MinInt32-right {
			return 0, ErrOverflow
		}
	}
	return left + right, nil
}
