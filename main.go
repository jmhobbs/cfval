package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/jagregory/cfval/reporting"
)

var forgiving = flag.Bool("forgiving", false, "Ignore unrecognised resources")

type ByContext reporting.Failures

func (a ByContext) Len() int           { return len(a) }
func (a ByContext) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByContext) Less(i, j int) bool { return a[i].ContextReadable < a[j].ContextReadable }

func printFailures(failures reporting.Failures) {
	sort.Sort(ByContext(failures))

	maxLength := 0
	for _, failure := range failures {
		context := failure.ContextReadable
		if len(context) > maxLength {
			maxLength = len(context)
		}
	}

	for _, failure := range failures {
		context := failure.ContextReadable

		fmt.Print(context)
		fmt.Print(" ")
		for i := 0; i < maxLength-len(context); i++ {
			fmt.Print(".")
		}
		fmt.Print("... ")
		fmt.Printf("%s\n", failure.Message)
	}
}

func printSummary(failures reporting.Failures) {
	fmt.Printf("%d failures\n", len(failures))
}

func main() {
	app := cli.NewApp()
	app.Name = "cfval"
	app.Usage = "CloudFormation template validator"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "James Gregory",
			Email: "james@jagregory.com",
		},
	}
	app.Version = "0.1.0"
	app.Action = func(c *cli.Context) {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading JSON from Stdin")
			return
		}

		template, err := parseTemplateJSON(bytes, *forgiving)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		if ok, errors := template.Validate(); ok {
			fmt.Println("Pass, no errors found.")
		} else {
			printFailures(errors)
			fmt.Println()
			printSummary(errors)
		}
	}

	app.Run(os.Args)
}
