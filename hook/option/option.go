package options

import (
	"fmt"
	"net/http"
	"github.com/samtjro/go-tda/utils"
)

type CHAIN struct {
	SYMBOL		string
	STATUS		string
	VOLATILITY	string
	STRIKE		int
}

// simpleOption takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL"
// strikeCount = number of strikes to return above and below the at-the-money price
// includeQuotes = "TRUE", "FALSE"
// strike = desired strike price
// strikeRange = returns option chains for a given range:
// ITM = in da money
// NTM = near da money
// OTM = out da money
// SAK = strikes above market
// SBK = strikes below market
// SNK = strikes near market
// ALL* = default, all strikes
func simple(ticker,contractType,strikeCount,includeQuotes,strike,strikeRange string) string {
	req,_ := http.NewRequest("GET",endpoint_option,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("contractType",contractType)
	q.Add("strikeCount",strikeCount)
	q.Add("includeQuotes",includeQuotes)
	q.add("strike",strike)
	q.add("range",strikeRange)
	req.URL.RawQuery = q.Encode()
	body := utils.handler(req)

	return body
}

//func strategy() string {}
