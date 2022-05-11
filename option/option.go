package option

import (
	"fmt"
	"net/http"
	. "github.com/samtjro/go-tda/utils"
)

var endpoint_option string = "https://api.tdameritrade.com/v1/marketdata/chains"

type CHAIN struct {
	SYMBOL		string
	STATUS		string
	VOLATILITY	string
	STRIKE		int
}

// Single returns a string; containing a SINGLE option chain of your desired strike, type, etc., 
// it takes four parameters:
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
func Single(ticker,contractType,strikeCount,includeQuotes,strikeRange string) string {
	req,_ := http.NewRequest("GET",endpoint_option,nil)
	q := req.URL.Query()
	q.Add("symbol",ticker)
	q.Add("contractType",contractType)
	q.Add("strikeCount",strikeCount)
	q.Add("includeQuotes",includeQuotes)
	q.add("range",strikeRange)
	req.URL.RawQuery = q.Encode()
	body := Handler(req)

	return body
}

// func Analytical() string {}
// func Covered() string {}
// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Butterfly() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}
