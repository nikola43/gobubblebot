package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
func GetTokenInfo(tokenAddress string) (TokenInfoDexTools, error) {
	url := "https://www.dextools.io/shared/search/pair?query=" + tokenAddress
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfoDexTools(body)
	if err != nil {
		panic(err)
	}

	return tokenInfo, nil
}
*/

/*
func GetTokenInfo(tokenAddress string) (TokenInfo, error) {
	url := "https://api.ethplorer.io/getTokenInfo/" + tokenAddress + "?apiKey=freekey"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfo(body)
	if err != nil {
		panic(err)
	}

	return tokenInfo, nil
}
*/

func GetTokenInfoDex(tokenAddress string) (Pair, error) {
	url := "https://api.dexscreener.com/latest/dex/tokens/" + tokenAddress

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfoDexscreener(body)
	if err != nil {
		panic(err)
	}

	var p Pair

	for i := 0; i < len(tokenInfo.Pairs); i++ {
		pair := tokenInfo.Pairs[i]
		labels := pair.Labels

		for j := 0; j < len(labels); j++ {
			label := pair.Labels[j]

			if *pair.ChainID == "ethereum" && *pair.DexID == "uniswap" && label == "v2" {
				//fmt.Println(pair)
				return pair, nil
			}
		}
	}

	return p, nil
}

func GetPairAddress(tokenAddress string) (string, error) {

	url := "https://api.dexscreener.com/latest/dex/tokens/" + tokenAddress

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfoDexscreener(body)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(tokenInfo.Pairs); i++ {
		pair := tokenInfo.Pairs[i]
		labels := pair.Labels

		for j := 0; j < len(labels); j++ {
			label := pair.Labels[j]

			if *pair.ChainID == "ethereum" && *pair.DexID == "uniswap" && label == "v2" {
				//fmt.Println(pair)
				return *pair.PairAddress, nil
			}
		}
	}
	return "", nil
}

func GetTokenHoldersAndTaxes(tokenAddress string) (HoldersTaxes, error) {
	// Define the API URL
	apiURL := "https://api.honeypot.is/v2/IsHoneypot?address=" + tokenAddress

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfoEthExplorer(body)
	if err != nil {
		panic(err)
	}

	holdersTaxes := HoldersTaxes{
		Holders: tokenInfo.HolderAnalysis.Holders,
		Buy:     fmt.Sprintf("%f", *tokenInfo.SimulationResult.BuyTax),
		Sell:    fmt.Sprintf("%f", *tokenInfo.SimulationResult.SellTax),
	}

	return holdersTaxes, nil
}
