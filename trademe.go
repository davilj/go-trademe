package main

import (
	"os"
	"fmt"
	"github.com/davilj/trademelib"
)


func main() {
	p:=fmt.Println
	p("Starting trademe")
	argsWithoutProgram := os.Args[1:]

	dirForDailySummary:=argsWithoutProgram[0]
	p("Running daily summary on:", dirForDailySummary )

	var pgSummaryHandler trademelib.PostgresSummaryHandler;
	trademelib.CalcDailySums(dirForDailySummary, pgSummaryHandler)

}
