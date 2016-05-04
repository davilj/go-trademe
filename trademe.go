package main

import (
	"os"
	"fmt"
	"github.com/davilj/trademelib"
)


func main() {
	p:=fmt.Println
	p("Starting trademe")
	dataDrive:="/trademeData"
	if (len(os.Args)>1) {
		dataDrive=os.Args[1]
	}
	p("Running daily summary on:", dataDrive )

	var pgSummaryHandler trademelib.PostgresSummaryHandler;
	trademelib.CalcDailySums(dirForDailySummary, pgSummaryHandler)

}
