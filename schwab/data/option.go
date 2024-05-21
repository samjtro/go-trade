package option 

// Anatomy of the TDA option return:
//{"symbol":"AAPL","status":"SUCCESS","underlying":null,"strategy":"SINGLE","interval":0.0,"isDelayed":true,"isIndex":false,"interestRate":0.1,"underlyingPrice":152.625,"volatility":29.0,"daysToExpiration":0.0,"numberOfContracts":270,"putExpDateMap":{},"callExpDateMap":{"2022-05-13:2":{
//"145.0":[{"putCall":"CALL","symbol":"AAPL_051322C145","description":"AAPL May 13 2022 145 Call (Weekly)","exchangeName":"OPR","bid":9.5,"ask":10.95,"last":9.8,"mark":10.23,"bidSize":50,"askSize":51,"bidAskSize":"50X51","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":9.87,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652212791092,"quoteTimeInLong":1652212798993,"netChange":-0.07,"volatility":49.752,"delta":0.903,"gamma":0.022,"theta":-0.181,"vega":0.027,"rho":0.013,"openInterest":907,"timeValue":0.29,"theoreticalOptionValue":9.874,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":145.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":-0.75,"markChange":0.35,"markPercentChange":3.55,"intrinsicValue":9.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],
//"146.0":[{"putCall":"CALL","symbol":"AAPL_051322C146","description":"AAPL May 13 2022 146 Call(Weekly)","exchangeName":"OPR","bid":8.85,"ask":9.75,"last":10.35,"mark":9.3,"bidSize":21,"askSize":50,"bidAskSize":"21X50","lastSize":0,"highPrice":0.0,"lowPrice":0.0,"openPrice":0.0,"closePrice":8.98,"totalVolume":0,"tradeDate":null,"tradeTimeInLong":1652209972509,"quoteTimeInLong":1652212799988,"netChange":1.37,"volatility":49.456,"delta":0.878,"gamma":0.026,"theta":-0.212,"vega":0.031,"rho":0.013,"openInterest":320,"timeValue":1.84,"theoreticalOptionValue":8.985,"theoreticalVolatility":29.0,"optionDeliverablesList":null,"strikePrice":146.0,"expirationDate":1652472000000,"daysToExpiration":2,"expirationType":"S","lastTradingDay":1652486400000,"multiplier":100.0,"settlementType":" ","deliverableNote":"","isIndexOption":null,"percentChange":15.2,"markChange":0.32,"markPercentChange":3.52,"intrinsicValue":8.51,"inTheMoney":true,"mini":false,"pennyPilot":true,"nonStandard":false}],

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/samjtro/go-trade/schwab/utils"
	"github.com/samjtro/go-trade/schwab"
)

var endpoint_option string = fmt.Sprintf(schwab.Endpoint + "/chains")

//type UNDERLYING struct {}

type CONTRACT struct {
	TYPE                   string
	SYMBOL                 string
	STRIKE                 float64
	EXCHANGE               string
	EXPIRATION             float64
	DAYS2EXPIRATION        float64
	BID                    float64
	ASK                    float64
	LAST                   float64
	MARK                   float64
	BIDASK_SIZE            string
	VOLATILITY             float64
	DELTA                  float64
	GAMMA                  float64
	THETA                  float64
	VEGA                   float64
	RHO                    float64
	OPEN_INTEREST          float64
	TIME_VALUE             float64
	THEORETICAL_VALUE      float64
	THEORETICAL_VOLATILITY float64
	PERCENT_CHANGE         float64
	MARK_CHANGE            float64
	MARK_PERCENT_CHANGE    float64
	INTRINSIC_VALUE        float64
	IN_THE_MONEY           bool //bool
}

// Single returns a []CONTRACT; containing a SINGLE option chain of your desired strike, type, etc.,
// it takes four parameters:
// ticker = "AAPL", etc.
// contractType = "CALL", "PUT", "ALL";
// strikeRange = returns option chains for a given range:
// ITM = in da money
// NTM = near da money
// OTM = out da money
// SAK = strikes above market
// SBK = strikes below market
// SNK = strikes near market
// ALL* = default, all strikes;
// strikeCount = The number of strikes to return above and below the at-the-money price;
// toDate = Only return expirations before this date. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
// Lets examine a sample call of Single: Single("AAPL","CALL","ALL","5","2022-07-01").
// This returns 5 AAPL CALL contracts both above and below the at the money price, with no preference as to the status of the contract ("ALL"), expiring before 2022-07-01
func Single(ticker, contractType, strikeRange, strikeCount, toDate string) ([]CONTRACT, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return []CONTRACT{}, err
	}

	var chain []CONTRACT
	var Type, symbol, exchange, bidAskSize string
	var strikePrice, exp, d2e, bid, ask, last, mark, volatility, delta, gamma, theta, vega, rho, openInterest, timeValue, theoreticalValue, theoreticalVolatility, percentChange, markChange, markPercentChange, intrinsicValue float64
	var inTheMoney bool
	split := strings.Split(body, "}],")

	for _, x := range split {
		split2 := strings.Split(x, "\"")

		for i, x := range split2 {
			switch x {
			case "putCall":
				Type = split2[i+2]
			case "symbol":
				symbol = split2[i+2]
			case "exchangeName":
				exchange = split2[i+2]
			case "strikePrice":
				strikePrice1 := utils.TrimFL(split2[i+1])

				strikePrice, err = strconv.ParseFloat(strikePrice1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "expirationDate":
				exp1 := utils.TrimFL(split2[i+1])

				exp, err = strconv.ParseFloat(exp1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "daysToExpiration":
				d2e1 := utils.TrimFL(split2[i+1])

				d2e, err = strconv.ParseFloat(d2e1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "bid":
				bid1 := utils.TrimFL(split2[i+1])

				bid, err = strconv.ParseFloat(bid1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "ask":
				ask1 := utils.TrimFL(split2[i+1])

				ask, err = strconv.ParseFloat(ask1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "last":
				last1 := utils.TrimFL(split2[i+1])

				last, err = strconv.ParseFloat(last1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "mark":
				mark1 := utils.TrimFL(split2[i+1])

				mark, err = strconv.ParseFloat(mark1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "bidAskSize":
				bidAskSize = split2[i+2]
			case "volatility":
				volatility1 := utils.TrimFL(split2[i+1])

				volatility, err = strconv.ParseFloat(volatility1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "delta":
				delta1 := utils.TrimFL(split2[i+1])

				delta, err = strconv.ParseFloat(delta1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "gamma":
				gamma1 := utils.TrimFL(split2[i+1])

				gamma, err = strconv.ParseFloat(gamma1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theta":
				theta1 := utils.TrimFL(split2[i+1])

				theta, err = strconv.ParseFloat(theta1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "vega":
				vega1 := utils.TrimFL(split2[i+1])

				vega, err = strconv.ParseFloat(vega1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "rho":
				rho1 := utils.TrimFL(split2[i+1])

				rho, err = strconv.ParseFloat(rho1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "openInterest":
				openInterest1 := utils.TrimFL(split2[i+1])

				openInterest, err = strconv.ParseFloat(openInterest1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "timeValue":
				timeValue1 := utils.TrimFL(split2[i+1])

				timeValue, err = strconv.ParseFloat(timeValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theoreticalOptionValue":
				theoreticalValue1 := utils.TrimFL(split2[i+1])

				theoreticalValue, err = strconv.ParseFloat(theoreticalValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "theoreticalVolatility":
				theoreticalVolatility1 := utils.TrimFL(split2[i+1])

				theoreticalVolatility, err = strconv.ParseFloat(theoreticalVolatility1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "percentChange":
				percentChange1 := utils.TrimFL(split2[i+1])

				percentChange, err = strconv.ParseFloat(percentChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "markChange":
				markChange1 := utils.TrimFL(split2[i+1])

				markChange, err = strconv.ParseFloat(markChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "markPercentChange":
				markPercentChange1 := utils.TrimFL(split2[i+1])

				markPercentChange, err = strconv.ParseFloat(markPercentChange1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "intrinsicValue":
				intrinsicValue1 := utils.TrimFL(split2[i+1])

				intrinsicValue, err = strconv.ParseFloat(intrinsicValue1, 64)

				if err != nil {
					log.Fatalf(err.Error())
				}
			case "inTheMoney":
				inTheMoney, err = strconv.ParseBool(utils.TrimFL(split2[i+1]))

				if err != nil {
					log.Fatalf(err.Error())
				}
			}
		}

		contract := CONTRACT{
			TYPE:                   Type,
			SYMBOL:                 symbol,
			STRIKE:                 strikePrice,
			EXCHANGE:               exchange,
			EXPIRATION:             exp,
			DAYS2EXPIRATION:        d2e,
			BID:                    bid,
			ASK:                    ask,
			LAST:                   last,
			MARK:                   mark,
			BIDASK_SIZE:            bidAskSize,
			VOLATILITY:             volatility,
			DELTA:                  delta,
			GAMMA:                  gamma,
			THETA:                  theta,
			VEGA:                   vega,
			RHO:                    rho,
			OPEN_INTEREST:          openInterest,
			TIME_VALUE:             timeValue,
			THEORETICAL_VALUE:      theoreticalValue,
			THEORETICAL_VOLATILITY: theoreticalVolatility,
			PERCENT_CHANGE:         percentChange,
			MARK_CHANGE:            markChange,
			MARK_PERCENT_CHANGE:    markPercentChange,
			INTRINSIC_VALUE:        intrinsicValue,
			IN_THE_MONEY:           inTheMoney,
		}

		chain = append(chain, contract)
	}

	return chain, nil
}

// Covered returns a string; containing covered option calls.
// Not functional ATM.
func Covered(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "COVERED")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// Butterfly returns a string; containing Butterfly spread option calls.
// Not functional ATM.
func Butterfly(ticker, contractType, strikeRange, strikeCount, toDate string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "BUTTERFLY")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// ANALYTICAL returns a string; allows you to control additional parameters for theoretical value calculations:
// It takes nine parameters:
// Not functional ATM.
func Analytical(ticker, contractType, strikeRange, strikeCount, toDate, volatility, underlyingPrice, interestRate, daysToExpiration string) (string, error) {
	req, _ := http.NewRequest("GET", endpoint_option, nil)
	q := req.URL.Query()
	q.Add("strategy", "ANALYTICAL")
	q.Add("symbol", ticker)
	q.Add("contractType", contractType)
	q.Add("range", strikeRange)
	q.Add("strikeCount", strikeCount)
	q.Add("toDate", toDate)
	q.Add("volatility", volatility)
	q.Add("underlyingPrice", underlyingPrice)
	q.Add("interestRate", interestRate)
	q.Add("daysToExpiration", underlyingPrice)
	req.URL.RawQuery = q.Encode()
	body, err := utils.Handler(req)

	if err != nil {
		return "", err
	}

	return body, nil
}

// func Vertical() string {}
// func Calendar() string {}
// func Strangle() string {}
// func Straddle() string {}
// func Condor() string {}
// func Diagonal() string {}
// func Collar() string {}
// func Roll() string {}