package containerWithMostWater

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/logrusorgru/aurora/v3"
)

type CwmwCmd struct {
	description string
	fs          *flag.FlagSet
	helpFlag    bool
}

func NewCwmwCmd() *CwmwCmd {
	containerWithMostWaterCmd := &CwmwCmd{
		description: "https://leetcode.com/problems/container-with-most-water/",
		fs:          flag.NewFlagSet("cwmw", flag.ContinueOnError),
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
		fmt.Println("\tcwmw [options]")
		fmt.Println(aurora.BrightRed("Options:"))
		fmt.Println(optionsTable.String())
	}

	return containerWithMostWaterCmd
}

func (c CwmwCmd) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c CwmwCmd) Run() error {
	if c.helpFlag {
		c.fs.Usage()
		os.Exit(0)
	}

	var intsLen int
	fmt.Println("Enter the length of slice: ")
	fmt.Scanf("%d", &intsLen)

	height := make([]int, intsLen)
	for i := 0; i < intsLen; i++ {
		fmt.Scanf("%d", &height[i])
	}

	return callMaxArea(height)
}

func (c CwmwCmd) Name() string {
	return c.fs.Name()
}

func (c CwmwCmd) Description() string {
	return c.description
}

func callMaxArea(height []int) error {
	res := maxArea(height)
	fmt.Println("Result: ")
	fmt.Println(res)

	return nil
}

func maxArea(height []int) int {
	var maxArea int
	var width int
	var h int
	var tmp int
	var left int
	right := len(height) - 1

	for left < right {
		width = right - left

		if height[left] > height[right] {
			h = height[right]
		} else {
			h = height[left]
		}

		tmp = h * width

		if tmp > maxArea {
			maxArea = tmp
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func maxArea1(height []int) int {
	res := 0
	tmp := 0
	var operand1 int
	var operand2 int

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				operand1 = height[j]
			} else {
				operand1 = height[i]
			}

			operand2 = j - i

			tmp = operand1 * operand2

			if tmp > res {
				res = tmp
			}
		}
	}

	return res
}
