package atoi

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/alexeyco/simpletable"
	"github.com/logrusorgru/aurora/v3"
)

type MyAtoiCmd struct {
	description string
	fs          *flag.FlagSet
	helpFlag    bool
}

func NewMyAtoiCmd() *MyAtoiCmd {
	myAtoiCmd := &MyAtoiCmd{
		description: "https://leetcode.com/problems/string-to-integer-atoi/",
		fs:          flag.NewFlagSet("atoi", flag.ContinueOnError),
	}

	// Here we specify flags
	myAtoiCmd.fs.BoolVar(&myAtoiCmd.helpFlag, "h", false, "Display help for the given command. When no command is given display help for the list command")

	//Custom usage
	myAtoiCmd.fs.Usage = func() {

		// Options section
		optionsTable := simpletable.New()

		myAtoiCmd.fs.VisitAll(func(f *flag.Flag) {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("-%s", f.Name)).String()},
				{Align: simpletable.AlignLeft, Text: f.Usage},
			}

			optionsTable.Body.Cells = append(optionsTable.Body.Cells, r)
		})

		optionsTable.SetStyle(simpletable.StyleCompactLite)

		// Printing help
		fmt.Println(aurora.BrightRed("Description:"))
		fmt.Println("\t", myAtoiCmd.description)
		fmt.Println(aurora.BrightRed("Usage:"))
		fmt.Println("\tmigrate [options]")
		fmt.Println(aurora.BrightRed("Options:"))
		fmt.Println(optionsTable.String())
	}

	return myAtoiCmd
}

func (m MyAtoiCmd) Init(args []string) error {
	return m.fs.Parse(args)
}

func (m MyAtoiCmd) Run() error {
	if m.helpFlag {
		m.fs.Usage()
		os.Exit(0)
	}

	//TODO ask for input

	return callAtoi("1234")
}

func (m MyAtoiCmd) Name() string {
	return m.fs.Name()
}

func (m MyAtoiCmd) Description() string {
	return m.description
}

func callAtoi(string2 string) error {
	//TODO implement me
	return nil
}

func myAtoi(s string) int {
	tmpStr := strings.TrimSpace(s)
	var sb strings.Builder

	for i, rn := range tmpStr {

		if i == 0 && (rn == '+' || rn == '-') {
			sb.WriteRune(rn)
			continue
		}

		if unicode.IsDigit(rn) {
			sb.WriteRune(rn)
		} else {
			break
		}

	}

	res, err := strconv.Atoi(sb.String())

	if err != nil {
		fmt.Println("Convert string to int")
		fmt.Println(err)
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
