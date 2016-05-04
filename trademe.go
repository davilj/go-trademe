package main

import (
	"os"
	"fmt"
	trademelib "github.com/davilj/trademe/lib"
	)


func main() {
	p:=fmt.Println
	p("Starting trademe..1")
	dataDrive:="/trademeData"
	if (len(os.Args)>1) {
		p("configuring data drive")
		dataDrive=os.Args[1]
	}
	p("Running daily summary on:", dataDrive )

	var pgSummaryHandler trademelib.PostgresSummaryHandler;
	trademelib.CalcDailySums(dataDrive, pgSummaryHandler)
	p("Completed....")

}
