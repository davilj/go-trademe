package trademelib

import (
	//"strings"
	"fmt"
	"testing"
	//"strconv"
)

var handle TestSummaryHandler

type TestSummaryHandler struct {
	summariesMap map[string]Summary

}

func init() {
	handle.summariesMap=make(map[string]Summary)
}

func (summaryHandler TestSummaryHandler) handleSummary(key string, summary Summary) {
	summaryHandler.summariesMap[key]=summary
}

func TestBuildDailySummaries(t *testing.T) {
	testDir := "testData"

	//handle.summariesMap=make(map[string]Summary)
	CalcDailySums(testDir, handle)
	result := handle.summariesMap
	//fmt.Println(result)
	if len(result) != 16 {
		t.Errorf("expecting [%d] summaries, but was [%d]", 16, len(result))
	}
	keysFor26, keysFor27 := buildForEachDayAndCat()
	for index := 0; index < len(keysFor26); index++ {
		key26 := keysFor26[index]
		key27 := keysFor27[index]
		//fmt.Println(key26)
		//fmt.Println(key27)
		value26 := result[key26]
		value27 := result[key27]
		//fmt.Println(value26)
		//fmt.Println(value27)
		if value26 == value27 {
			t.Errorf("expecting [%s], but was [%s]", value26, value27)
		}
	}

	//check if D is correctly calculated
	dValue := result[":computers:D|2016-01-26"]

	checkValueF(t, 500.0, dValue.mean)
	checkValueI(t, 100, dValue.min)
	checkValueI(t, 900, dValue.max)
	checkValueI(t, 800, dValue.rangeValue)
	checkValueF(t, 258.19888, dValue.stDev)
	checkValueI(t, 9, dValue.numberOfTransaction)
	checkValueI(t, 4500, dValue.total)
	percentiles := dValue.percentiles
	checkValueF(t, 500.0, percentiles.p50)
	checkValueF(t, 300.0, percentiles.p25)
	checkValueF(t, 700.0, percentiles.p75)

}

func checkValueF(t *testing.T, expected float32, actual float32) {
	if expected != actual {
		t.Errorf("expecting [%s], but was [%s]", expected, actual)
	}
}

func checkValueI(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("expecting [%s], but was [%s]", expected, actual)
	}
}

func buildForEachDayAndCat() ([8]string, [8]string) {
	days := []string{"2016-01-26", "2016-01-27"}
	catPart1 := []string{"computers", "001"}
	catPart2 := []string{"A", "B", "C", "D"}
	keysFor26 := [8]string{}
	keysFor27 := [8]string{}
	index26 := 0
	index27 := 0
	for _, day := range days {
		for _, part1 := range catPart1 {
			for _, part2 := range catPart2 {
				key := fmt.Sprintf(":%s:%s|%s", part1, part2, day)
				if day == "2016-01-26" {
					//fmt.Println(index26, key)
					keysFor26[index26] = key
					index26 = index26 + 1
				} else {
					//fmt.Println(index27, key)
					keysFor27[index27] = key
					index27 = index27 + 1
				}
			}
		}
	}
	return keysFor26, keysFor27
}
