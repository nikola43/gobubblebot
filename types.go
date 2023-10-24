package main

import "encoding/json"

type Account struct {
	Address    string
	PrivateKey string
}

type TokenConfig struct {
	Address  string
	Pair     string
	Group    string
	Website  string
	Telegram string
	Twitter  string
}

func UnmarshalTokenInfo(data []byte) (TokenInfo, error) {
	var r TokenInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokenInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokenInfo struct {
	Address           string `json:"address"`
	Decimals          string `json:"decimals"`
	Name              string `json:"name"`
	Owner             string `json:"owner"`
	Price             Price  `json:"price"`
	Symbol            string `json:"symbol"`
	TotalSupply       string `json:"totalSupply"`
	TransfersCount    int64  `json:"transfersCount"`
	TxsCount          int64  `json:"txsCount"`
	IssuancesCount    int64  `json:"issuancesCount"`
	LastUpdated       int64  `json:"lastUpdated"`
	HoldersCount      int64  `json:"holdersCount"`
	Website           string `json:"website"`
	EthTransfersCount int64  `json:"ethTransfersCount"`
	CountOps          int64  `json:"countOps"`
}

type Price struct {
	Rate            float64 `json:"rate"`
	Diff            float64 `json:"diff"`
	Diff7D          float64 `json:"diff7d"`
	Ts              int64   `json:"ts"`
	MarketCapUsd    float64 `json:"marketCapUsd"`
	AvailableSupply int64   `json:"availableSupply"`
	Volume24H       float64 `json:"volume24h"`
	VolDiff1        float64 `json:"volDiff1"`
	VolDiff7        float64 `json:"volDiff7"`
	VolDiff30       float64 `json:"volDiff30"`
	Diff30D         float64 `json:"diff30d"`
	Currency        string  `json:"currency"`
}

///----------------

func UnmarshalTokenInfoDex(data []byte) (TokenInfoDex, error) {
	var r TokenInfoDex
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokenInfoDex) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokenInfoDex struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         []Pair `json:"pairs"`
}

type Pair struct {
	ChainID       ChainID     `json:"chainId"`
	DexID         DexID       `json:"dexId"`
	URL           string      `json:"url"`
	PairAddress   string      `json:"pairAddress"`
	Labels        []Label     `json:"labels,omitempty"`
	BaseToken     EToken      `json:"baseToken"`
	QuoteToken    EToken      `json:"quoteToken"`
	PriceNative   string      `json:"priceNative"`
	PriceUsd      string      `json:"priceUsd"`
	Txns          Txns        `json:"txns"`
	Volume        PriceChange `json:"volume"`
	PriceChange   PriceChange `json:"priceChange"`
	Liquidity     Liquidity   `json:"liquidity"`
	Fdv           int64       `json:"fdv"`
	PairCreatedAt int64       `json:"pairCreatedAt"`
}

type EToken struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Liquidity struct {
	Usd   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

type PriceChange struct {
	H24 float64 `json:"h24"`
	H6  float64 `json:"h6"`
	H1  float64 `json:"h1"`
	M5  float64 `json:"m5"`
}

type Txns struct {
	M5  H1 `json:"m5"`
	H1  H1 `json:"h1"`
	H6  H1 `json:"h6"`
	H24 H1 `json:"h24"`
}

type H1 struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

type ChainID string

const (
	Ethereum   ChainID = "ethereum"
	Pulsechain ChainID = "pulsechain"
)

type DexID string

const (
	Pancakeswap  DexID = "pancakeswap"
	Pulsex       DexID = "pulsex"
	Safemoonswap DexID = "safemoonswap"
	Uniswap      DexID = "uniswap"
)

type Label string

const (
	V1 Label = "v1"
	V2 Label = "v2"
	V3 Label = "v3"
)
