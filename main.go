// This program performs calls to algorithms solutions
package main

import (
	"errors"
	"fmt"
	"github.com/kulyklev/leetCodeSolutions/solutions/regularExpressionMatching"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/kulyklev/leetCodeSolutions/solutions"
	"github.com/kulyklev/leetCodeSolutions/solutions/atoi"
	"github.com/logrusorgru/aurora/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {

	// Construct the application logger.
	log, err := newLogger("ADMIN")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		if !errors.Is(err, solutions.ErrHelp) {
			log.Errorw("startup", "ERROR", err)
		}
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// =================================================================================================================
	// Registering commands

	cmds := []solutions.Runner{
		atoi.NewMyAtoiCmd(),
		regularExpressionMatching.NewRegularExpressionMatchingCmd(),
	}

	// =================================================================================================================
	// Help

	if len(os.Args) < 2 {
		return displayHelp(cmds)
	}

	// =================================================================================================================
	// Command execution

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[1:])
			return cmd.Run()
		}
	}

	return solutions.ErrFailedToRun
}

func newLogger(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}

func displayHelp(cmds []solutions.Runner) error {
	fmt.Println("\n")

	// =================================================================================================================
	// Options sections

	optionsData := [][]interface{}{
		{"-h, --help", "Display help for the given command. When no command is given display help for the list command"},
	}

	optionsTable := simpletable.New()

	for _, row := range optionsData {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("%s", row[0].(string))).String()},
			{Align: simpletable.AlignLeft, Text: row[1].(string)},
		}

		optionsTable.Body.Cells = append(optionsTable.Body.Cells, r)
	}

	optionsTable.SetStyle(simpletable.StyleCompactLite)

	// =================================================================================================================
	// Available commands sections

	availableCommandsData := map[string]string{}
	for _, cmd := range cmds {
		availableCommandsData[cmd.Name()] = cmd.Description()
	}

	availableCommandsTable := simpletable.New()

	for cmdName, cmdDesc := range availableCommandsData {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: aurora.Green(fmt.Sprintf("%s", cmdName)).String()},
			{Align: simpletable.AlignLeft, Text: cmdDesc},
		}

		availableCommandsTable.Body.Cells = append(availableCommandsTable.Body.Cells, r)
	}

	availableCommandsTable.SetStyle(simpletable.StyleCompactLite)

	// =================================================================================================================
	// Display everything

	fmt.Println("App name")
	fmt.Println("Version: ", aurora.Green(build).Bold())
	fmt.Println(aurora.BrightRed("Usage:"))
	fmt.Println("\tcommand [options] [arguments]")
	fmt.Println(aurora.BrightRed("Options:"))
	fmt.Println(optionsTable.String())
	fmt.Println(aurora.BrightRed("Available commands:"))
	fmt.Println(availableCommandsTable.String())

	return solutions.ErrHelp
}
