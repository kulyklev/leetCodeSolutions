package integerToRoman

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/logrusorgru/aurora/v3"
)

type IntToRomanCmd struct {
	description string
	fs          *flag.FlagSet
	helpFlag    bool
}

func NewIntToRomanCmd() *IntToRomanCmd {
	containerWithMostWaterCmd := &IntToRomanCmd{
		description: "https://leetcode.com/problems/integer-to-roman/",
		fs:          flag.NewFlagSet("intToRomanCmd", flag.ContinueOnError),
	}

	// Here we specify flags
	containerWithMostWaterCmd.fs.BoolVar(&containerWithMostWaterCmd.helpFlag, "h", false, "Display help for the given command. When no command is given display help for the list command")

	//Custom usage
	containerWithMostWaterCmd.fs.Usage = func() {

		// Options section
		optionsTable := simpletable.New()

		containerWithMostWaterCmd.fs.VisitAll(func(f *flag.Flag) {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("-%s", f.Name)).String()},
				{Align: simpletable.AlignLeft, Text: f.Usage},
			}

			optionsTable.Body.Cells = append(optionsTable.Body.Cells, r)
		})

		optionsTable.SetStyle(simpletable.StyleCompactLite)

		// Printing help
		fmt.Println(aurora.BrightRed("Description:"))
		fmt.Println("\t", containerWithMostWaterCmd.description)
		fmt.Println(aurora.BrightRed("Usage:"))
		fmt.Println("\tintToRomanCmd [options]")
		fmt.Println(aurora.BrightRed("Options:"))
		fmt.Println(optionsTable.String())
	}

	return containerWithMostWaterCmd
}

func (i IntToRomanCmd) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i IntToRomanCmd) Run() error {
	if i.helpFlag {
		i.fs.Usage()
		os.Exit(0)
	}

	var intToConvert int
	fmt.Println("Enter the integer to convert to Roman: ")
	_, err := fmt.Scanf("%d", &intToConvert)
	if err != nil {
		return err
	}

	return callIntToRoman(intToConvert)
}

func (i IntToRomanCmd) Name() string {
	return i.fs.Name()
}

func (i IntToRomanCmd) Description() string {
	return i.description
}

func callIntToRoman(num int) error {
	res := intToRoman(num)
	fmt.Println("Result: ")
	fmt.Println(res)

	return nil
}

func intToRoman(num int) string {
	digit := 1
	var pop int
	var romeNums []string
	var sb strings.Builder
	mapTable := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	for num != 0 {
		powerOfTen := int(math.Pow10(digit))
		pop = num % 10
		num /= 10
		first := powerOfTen / 10
		second := powerOfTen / 2
		third := powerOfTen

		subMapTable := map[int]string{
			1:  mapTable[first],
			2:  mapTable[first] + mapTable[first],
			3:  mapTable[first] + mapTable[first] + mapTable[first],
			4:  mapTable[first] + mapTable[second],
			5:  mapTable[second],
			6:  mapTable[second] + mapTable[first],
			7:  mapTable[second] + mapTable[first] + mapTable[first],
			8:  mapTable[second] + mapTable[first] + mapTable[first] + mapTable[first],
			9:  mapTable[first] + mapTable[third],
			10: mapTable[third],
		}

		romeNums = append(romeNums, subMapTable[pop])
		digit++
	}

	reverse(romeNums)
	for _, p := range romeNums {
		sb.WriteString(p)
	}

	return sb.String()
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
