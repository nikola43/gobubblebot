package main

import (
	"io"
	"log"
	"net/http"
)

func GetTokenInfo(tokenAddress string) (TokenInfo, error) {
	url := "https://api.ethplorer.io/getTokenInfo/" + tokenAddress + "?apiKey=freekey"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfo(body)
	if err != nil {
		log.Fatalln(err)
	}

	return tokenInfo, nil
}

func GetPairAddress(tokenAddress string) (string, error) {

	url := "https://api.dexscreener.com/latest/dex/tokens/" + tokenAddress

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	tokenInfo, err := UnmarshalTokenInfoDex(body)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(tokenInfo.Pairs); i++ {
		pair := tokenInfo.Pairs[i]
		labels := pair.Labels

		for j := 0; j < len(labels); j++ {
			label := pair.Labels[j]

			if pair.ChainID == "ethereum" && pair.DexID == "uniswap" && label == "v2" {
				//fmt.Println(pair)
				return pair.PairAddress, nil
			}
		}
	}
	return "", nil
}
