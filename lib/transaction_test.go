package trademelib

import (
	"fmt"
	"strings"
	"testing"
	//"strconv"
)

func TestTransaction(t *testing.T) {
	strData := "LatestListing [title=Swarovski Love Turtledoves New. Pay now., link=/pottery-glass/glass-crystal/ornaments/birds/auction-1018856592.htm, closingTimeText=closes in 57 secs, bidInfo=, priceInfo=];" +
		"LatestListing [title=Swarovski Pacifier Small Light Sapphire New. Pay now., link=/pottery-glass/glass-crystal/ornaments/other/auction-1018856724.htm, closingTimeText=closes in 1 min, bidInfo=, priceInfo=];" +
		"LatestListing [title=Swarovski Chinese Symbols Double Carps Ornament New. Pay now., link=/pottery-glass/glass-crystal/ornaments/other-animals/auction-1018856731.htm, closingTimeText=closes in 1 min, bidInfo=, priceInfo=];" +
		"LatestListing [title=Swarovski Pram New. Pay now., link=/pottery-glass/glass-crystal/ornaments/other/auction-1018856747.htm, closingTimeText=closes in 1 min, bidInfo=2 bid, priceInfo=$100.00];" +
		"LatestListing [title=Zibo Speckled Hen Ornament New. Pay now., link=/pottery-glass/glass-crystal/ornaments/birds/auction-1018856759.htm, closingTimeText=closes in 1 min, bidInfo=, priceInfo=$10.00];" +
		"LatestListing [title=Zibo Pied Sow New. Pay now., link=/pottery-glass/glass-crystal/ornaments/other-animals/auction-1018856764.htm, closingTimeText=closes in 1 min, bidInfo=, priceInfo=];" +
		"LatestListing [title=Zibo Horned Owl New. Pay now., link=/pottery-glass/glass-crystal/ornaments/birds/auction-1018857187.htm, closingTimeText=closes in 2 mins, bidInfo=1 bid, priceInfo=];"

	transactions, err := ParseLines(strData, "201601262142")
	if err != nil {
		t.Errorf("Error in parsing lines %s", err)
	}
	amountOfTransaction := len(transactions)
	if amountOfTransaction != 7 {
		t.Errorf("amountOfTransaction expected 7 was %d", amountOfTransaction)
	}
	expectedTransaction := strings.Split(strData, ";")
	for index, transactionStr := range expectedTransaction {
		if transactionStr != "" {
			testTransaction(t, transactionStr, transactions[index])
		}
	}
}

func testTransaction(t *testing.T, expectedTransactionStr string, transaction Transaction) int {
	//cleanUPStr := transactionStr[15 : len(transactionStr)-1]
	cleanUPStr := expectedTransactionStr[15 : len(expectedTransactionStr)-1]
	m := buildPropMap(cleanUPStr)
	for key, value := range m {
		if "title" == key {
			if transaction.title != value {
				t.Errorf("title expected[%s] but [%s]", transaction.title, value)
			}
		}
		if "priceInfo" == key {
			if transaction.price == 0 {
				if value != "" {
					t.Errorf("priceInfo expected[%s] but [%s]", transaction.price, value)
				}
			} else {
				priceAsFloat := float64(transaction.price) / 100
				priceAsString := fmt.Sprintf("$%.2f", priceAsFloat)
				//fmt.Println("PR_FL: ", priceAsString)
				if priceAsString != value {
					t.Errorf("priceInfo expected[%s] but [%s]", transaction.price, value)
				}
			}
		}
	}
	//fmt.Println(cleanUPStr)
	return 0
}

func buildPropMap(transactionStr string) map[string]string {
	m := make(map[string]string)
	tokens := strings.Split(transactionStr, ",")
	for _, token := range tokens {
		propPart := strings.Split(token, "=")
		key := strings.Trim(propPart[0], " ")
		value := strings.Trim(propPart[1], " ")
		m[key] = value
	}
	return m
}
