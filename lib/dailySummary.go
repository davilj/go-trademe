package trademe

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
	"time"
	"fmt"
)

type SummaryHandler interface {
	handleSummary(key string, summary Summary)
}

type Percentiles struct {
	p25 float32
	p50 float32
	p75 float32
}

type Summary struct {
	cat                 string
	day                 string
	mean                float32
	min                 int
	max                 int
	rangeValue          int
	stDev               float32
	numberOfTransaction int
	total               int
	percentiles         Percentiles
}

func CalcDailySums(parentDir string, summaryHandler SummaryHandler) {
	//allSummaries := make(map[string]Summary)
	mapDayAndCat2Files, err := ParseDailyDirs(parentDir)
	if err != nil {

	}
	for key, value := range mapDayAndCat2Files {
		for _, fileName := range value {
			fmt.Println("--",fileName)
			transactions, err := ExtractTransaction(key + "/" + fileName)
			if err != nil {

			}
			transactionMap := make(map[string][]int)
			for _, transaction := range transactions {
				if transaction.bids > 0 {
					tm := time.Unix(transaction.closingTime, 0)

					key := transaction.cat + "|" + tm.Format("2006-01-02")
					transactionArr := transactionMap[key]
					newArr := append(transactionArr, transaction.price)
					transactionMap[key] = newArr
				}
			}

			summariesMap := extractSummaries(transactionMap)
			for key, value := range summariesMap {
				//allSummaries[key] = value
				summaryHandler.handleSummary(key, value)
			}
		}
	}
}

func extractSummaries(transactionMap map[string][]int) map[string]Summary {
	//numberOfSummaries := len(transactionMap);
	summaryMap := make(map[string]Summary)
	for key, value := range transactionMap {
		summary := summarise(key, value)
		summaryMap[key] = summary
	}
	return summaryMap
}

func summarise(catAndDate string, transactions []int) Summary {
	tokens := strings.Split(catAndDate, "|")
	cat := tokens[0]
	time := tokens[1]
	sort.Ints(transactions)
	numberOfTransaction := len(transactions)
	min := transactions[0]
	max := transactions[numberOfTransaction-1]
	rangeValue := max - min
	mean := calcMean(transactions)
	percentiles := calcPercentiles(transactions)
	total := calcTotal(transactions)
	stdDev := calcStandardDeviation(float64(mean), transactions)
	//time := "2006-01-02"

	summary := Summary{cat, time, mean, min, max, rangeValue, stdDev, numberOfTransaction, total, percentiles}
	return summary
}

func calcTotal(values []int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}

func calcStandardDeviation(mean float64, transactions []int) float32 {
	total := 0.0
	for _, value := range transactions {
		diff := float64(value) - mean
		total = total + (diff * diff)
	}
	return float32(math.Sqrt(total / float64(len(transactions))))
}

func calcMean(values []int) float32 {
	sum := 0
	for _, value := range values {
		sum = sum + value
	}
	return float32(sum) / float32(len(values))
}

func calcPercentiles(values []int) Percentiles {
	if len(values) == 1 {
		return Percentiles{float32(values[0]), float32(values[0]), float32(values[0])}
	}
	numberOfValues := len(values)
	percentiles := Percentiles{0.0, 0.0, 0.0}
	index := numberOfValues / 2
	if math.Mod(float64(numberOfValues), 2.0) == 0.0 {
		percentiles.p50 = float32(values[index-1]+values[index]) / 2.0
	} else {
		percentiles.p50 = float32(values[numberOfValues/2])
	}
	index4 := index / 2
	percentiles.p25 = float32(values[index4])
	percentiles.p75 = float32(values[numberOfValues-1-index4])
	return percentiles
}

func ParseDailyDirs(parentDir string) (map[string][]string, error) {
	mapDayAndCar2Files := make(map[string][]string)
	days, err := ioutil.ReadDir(parentDir)
	if err != nil {
		return mapDayAndCar2Files, err
	}

	for _, day := range days {
		dayDir := parentDir + "/" + day.Name()
		catInDays, err := ioutil.ReadDir(dayDir)
		if err != nil {
			return mapDayAndCar2Files, err
		}
		for _, cat := range catInDays {
			catDir := dayDir + "/" + cat.Name()
			files, err := ioutil.ReadDir(catDir)
			if err != nil {
				return mapDayAndCar2Files, err
			}
			filesName := make([]string, len(files))
			for index, file := range files {
				filesName[index] = file.Name()
			}
			mapDayAndCar2Files[catDir] = filesName
		}
	}
	return mapDayAndCar2Files, nil
}
