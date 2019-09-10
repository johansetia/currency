# currency
[![GoDoc](https://godoc.org/github.com/johansetia/currency?status.svg)](https://godoc.org/github.com/johansetia/currency)
[![Go Report Card](https://goreportcard.com/badge/github.com/johansetia/currency)](https://goreportcard.com/report/github.com/johansetia/currency)


currency is a simple package that can be used to find a real time currency from https://fixer.io/

### Instalation
```bash
$ go get https://github.com/johansetia/currency
```
### Usage
Keep it simple to use. before using this library, you must registed on this https://fixer.io to gat an API Key.

```go
package main

import (
	"fmt"
	"github.com/johansetia/currency"
	"github.com/johansetia/currency/code"
)

var key = "YOUR-SECRET-ACCESS-KEY"

func main() {
	c := currency.New(key)
	symbol, _ := c.NewSymbol()

	fmt.Println(symbol.Data)   // it will return data symbol
	fmt.Println(symbol.Status) // it will return a bolean type

	history, _ := c.History("2019-09-10", code.EUR, code.IDR)
	fmt.Println(history.Rates) // to get a rates data

	base, _ := c.SetBase(code.EUR)
	fmt.Println(base.Rates) // to get a rates data

	latest, _ := c.LatestRates(code.EUR, code.USD, code.IDR)
	fmt.Println(latest.Rates) // to get a rates data

	conv, _ := c.Convert(code.EUR, code.IDR, 1)
	fmt.Println(conv.Query.Amount)
	fmt.Println(conv.Info.Rate) // to get a rate data
	fmt.Println(conv.Result)    // to get a result data after converted

	fluc, _ := c.Time(currency.Fluctuation, "2019-09-08", "2019-09-09")
	fmt.Println(fluc.Rates) // to get a fluctuation data

	timeseties, _ := c.Time(currency.TimeSeries, "2019-09-08", "2019-09-09")
	fmt.Println(timeseties.Rates) // to get a timeseties data

}

```
### Reference
 Fixer.io Documentation :  https://fixer.io/documentation
### Structure

```bash
.
├── code
│   └── code.go
├── currency.go
├── go.mod
├── LICENSE
├── model.go
└── README.md

1 directory, 6 files

```


made with :purple_heart: by johan setiawan