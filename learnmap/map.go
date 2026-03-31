package main

import (
	"fmt"
)

type currencyStruct struct {
	name   string
	symbol string
}

func main() {
	currencyCode := make(map[string]string) // allocate memory for map and declare
	//alternatively m := map[string]string{} -> uses {"USD": "US Dollar"} if already has data
	// var currencyCode map[string]string -> nil map panic (only declare but not allocate mem - in case of maps and slices)

	currencyCode["USD"] = "US Dollar"
	currencyCode["EUR"] = "Euro"
	currencyCode["GBP"] = "GB Pound"
	fmt.Println(currencyCode)
	currency := "USD"
	currencyValue := currencyCode[currency]
	fmt.Println("currency code: ", currency, " currency value: ", currencyValue)

	notExist := currencyCode["nonExist"]
	fmt.Println("value return if key not exist: ", notExist)
	// -> return "" zero value string; for other types, return zero value of those type ( 0 for int, nil for pointer, etc.)
	// to check key exists: value, ok := map[key] -> ok return true false
	cyCode := "INR"
	if currencyName, ok := currencyCode[cyCode]; ok {
		fmt.Println("Currency name for currency code", cyCode, " is", currencyName)
		return
	}
	fmt.Println("Currency name for currency code", cyCode, "not found")

	//Iterate over map
	for k, v := range currencyCode {
		fmt.Println("Currency code: ", k, " is", v)
	}

	//delete a map element
	delete(currencyCode, "USD")

	//map of structs
	curUSD := currencyStruct{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currencyStruct{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currencyStruct{
		name:   "Euro",
		symbol: "€",
	}

	currencyCodeStruct := map[string]currencyStruct{
		"USD": curUSD,
		"EUR": curGBP,
		"GBP": curEUR,
	}

	for cyCode, cyInfo := range currencyCodeStruct {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n", cyCode, cyInfo.name, cyInfo.symbol)
	}

	fmt.Println("length of map using len(): ", len(currencyCode))

	// map is a reference type (similar to slice) -> variable name points to memory, meaning changing a variable affect the other.
	// map cannot be compared using == (only can check map == nil) -> check individual elements 1 by 1 or using reflection (check learnreflection)
}
