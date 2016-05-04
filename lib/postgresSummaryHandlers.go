package trademe

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    //"time"
)

const (
    DB_USER     = "pguser"
    DB_PASSWORD = "pguser"
    DB_NAME     = "trademe"
    DB_HOST     = "192.168.99.100"
)


type PostgresSummaryHandler struct {
	//summariesMap map[string]Summary

}

func TestVisibility() {
  fmt.Println("OK")
}

func (summaryHandler PostgresSummaryHandler) handleSummary(key string, summary Summary) {
	//summaryHandler.summariesMap[key]=summary
  dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  defer db.Close()

  sql:="INSERT INTO dailysummaries(cat,day,mean,min,max,rangevalue,stdev,numberoftransaction,total,p25,p50,p75) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);"
  queryStmt, err := db.Prepare(sql)
  checkErr(err)

  _, err = queryStmt.Exec(summary.cat, summary.day, summary.mean, summary.min, summary.max, summary.rangeValue, summary.stDev, summary.numberOfTransaction, summary.total,summary.percentiles.p25, summary.percentiles.p50, summary.percentiles.p75)
  checkErr(err)

}
