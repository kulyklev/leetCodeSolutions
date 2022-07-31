package palindromeNumber

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/logrusorgru/aurora/v3"
)

type PalindromeNumCmd struct {
	description string
	fs          *flag.FlagSet
	helpFlag    bool
}

func NewPalindromeNumCmd() *PalindromeNumCmd {
	palindromeNumCmd := &PalindromeNumCmd{
		description: "https://leetcode.com/problems/palindrome-number/",
		fs:          flag.NewFlagSet("isPalindromeNum", flag.ContinueOnError),
	}

	// Here we specify flags
	palindromeNumCmd.fs.BoolVar(&palindromeNumCmd.helpFlag, "h", false, "Display help for the given command. When no command is given display help for the list command")

	//Custom usage
	palindromeNumCmd.fs.Usage = func() {

		// Options section
		optionsTable := simpletable.New()

		palindromeNumCmd.fs.VisitAll(func(f *flag.Flag) {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("-%s", f.Name)).String()},
				{Align: simpletable.AlignLeft, Text: f.Usage},
			}

			optionsTable.Body.Cells = append(optionsTable.Body.Cells, r)
		})

		optionsTable.SetStyle(simpletable.StyleCompactLite)

		// Printing help
		fmt.Println(aurora.BrightRed("Description:"))
		fmt.Println("\t", palindromeNumCmd.description)
		fmt.Println(aurora.BrightRed("Usage:"))
		fmt.Println("\tisPalindromeNum [options]")
		fmt.Println(aurora.BrightRed("Options:"))
		fmt.Println(optionsTable.String())
	}

	return palindromeNumCmd
}

func (p PalindromeNumCmd) Init(args []string) error {
	return p.fs.Parse(args)
}

func (p PalindromeNumCmd) Run() error {
	if p.helpFlag {
		p.fs.Usage()
		os.Exit(0)
	}

	var s string
	fmt.Println("Enter number to check if it palindrome: ")
	_, err := fmt.Scanln(&s)
	if err != nil {
		return err
	}

	return callIsPalindrome(s)
}

func (p PalindromeNumCmd) Name() string {
	return p.fs.Name()
}

func (p PalindromeNumCmd) Description() string {
	return p.description
}

func callIsPalindrome(s string) error {
	num, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	res := isPalindrome(num)
	fmt.Println("Result: ")
	fmt.Println(res)

	return nil
}

func isPalindrome(x int) bool {
	input := x

	if x < 0 {
		return false
	}

	rev := 0
	pop := 0

	for x != 0 {
		pop = x % 10
		x /= 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			fmt.Println("Max int overflow")
		}

		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			fmt.Println("Min int overflow")
		}

		rev = rev*10 + pop
	}

	return input == rev
}
