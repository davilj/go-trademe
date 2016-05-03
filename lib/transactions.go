package trademelib

import (
	//"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	title       string
	id          int64
	cat         string
	closingTime int64
	bids        int
	price       int
}

func ExtractTransaction(fileName string) ([]Transaction, error) {
	transactions := make([]Transaction, 0)
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return transactions, err
	}
	tokens := strings.Split(fileName, "/")
	nameOfFile := tokens[len(tokens)-1]
	dateOfFile := (strings.Split(nameOfFile, "."))[0]
	//fmt.Println(dateOfFile)
	return ParseLines(string(dat), dateOfFile)

}

func ParseLines(data string, dateStr string) ([]Transaction, error) {
	allTransactions := make([]Transaction, 0)
	layout := "200601021504"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return allTransactions, err
	}
	unixTime := t.Unix()

	lines := strings.Split(data, ";")
	for _, line := range lines {
		if line != "" && len(line) > 15 {
			//fmt.Println("Parsing [%s]", line)
			transaction := ParseTransaction(line, unixTime)
			allTransactions = append(allTransactions, transaction)
		}
	}
	return allTransactions, nil
}

func convert2Int(aString string) (int, error) {
	if "" == aString {
		return 0, nil
	} else {
		i, err := strconv.Atoi(aString)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
}

func convert2Price(aString string) int {
	numberRegx, _ := regexp.Compile("[0-9]+.[0-9]{2}")
	money, _ := strconv.ParseFloat(numberRegx.FindString(aString), 32)
	return int(money * 100)
}

func extractTimeInMin(aTime string) int64 {
	hourRegx, _ := regexp.Compile("[0-9]+ hr")
	minRegx, _ := regexp.Compile("[0-9]+ min")
	hour := extractNumber(hourRegx.FindString(aTime))
	min := extractNumber(minRegx.FindString(aTime))
	return hour*60 + min
}

func extractNumber(aTime string) int64 {
	numberRegx, _ := regexp.Compile("[0-9]+")
	time, _ := strconv.Atoi(numberRegx.FindString(aTime))
	return int64(time)
}

func ParseTransaction(line string, unixTime int64) Transaction {
	//fmt.Println("Parsing: [" + line + "], " + strconv.Itoa(len(line)))
	data := line[15 : len(line)-1]
	tokens := strings.Split(data, ",")
	transaction := Transaction{"no title", 0, "no cat", 0, 0, 0}
	for _, token := range tokens {
		parts := strings.Split(token, "=")
		key := strings.Trim(parts[0], " ")
		value := ""
		if len(parts) == 2 {
			value = strings.Trim(parts[1], " ")
		}

		if key == "title" {
			transaction.title = value
		}
		if key == "link" {
			tokens := strings.Split(value, "/")
			lastIndex := len(tokens) - 1
			transaction.cat = strings.Join(tokens[:lastIndex], ":")
			transaction.id = extractNumber(tokens[lastIndex])
		}
		if key == "closingTimeText" {
			minutes := extractTimeInMin(value)
			transaction.closingTime = unixTime + minutes*60
		}
		if key == "bidInfo" {
			transaction.bids = int(extractNumber(value))
		}
		if key == "priceInfo" {
			transaction.price = convert2Price(value)
		}
	}
	return transaction
}
