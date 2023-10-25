package main

import "encoding/json"

type Account struct {
	Address    string
	PrivateKey string
}

type TokenConfig struct {
	Symbol   string
	Address  string
	Pair     string
	Group    string
	Website  string
	Telegram string
	Twitter  string
}

// --------------------------
type HoldersTaxes struct {
	Holders *string
	Buy     string
	Sell    string
}

func UnmarshalTokenInfoDexscreener(data []byte) (TokenInfoDexscreener, error) {
	var r TokenInfoDexscreener
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokenInfoDexscreener) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokenInfoDexscreener struct {
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	Pairs         []Pair  `json:"pairs,omitempty"`
}

type Pair struct {
	ChainID       *string       `json:"chainId,omitempty"`
	DexID         *string       `json:"dexId,omitempty"`
	URL           *string       `json:"url,omitempty"`
	PairAddress   *string       `json:"pairAddress,omitempty"`
	Labels        []string      `json:"labels,omitempty"`
	BaseToken     *EToken       `json:"baseToken,omitempty"`
	QuoteToken    *EToken       `json:"quoteToken,omitempty"`
	PriceNative   *string       `json:"priceNative,omitempty"`
	PriceUsd      *string       `json:"priceUsd,omitempty"`
	Volume        *VolumeChange `json:"volume,omitempty"`
	Fdv           *float64        `json:"fdv,omitempty"`
	PairCreatedAt *int64        `json:"pairCreatedAt,omitempty"`
}

type EToken struct {
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
	Symbol  *string `json:"symbol,omitempty"`
}

type VolumeChange struct {
	M5  *float64 `json:"m5,omitempty"`
	H1  *float64 `json:"h1,omitempty"`
	H6  *float64 `json:"h6,omitempty"`
	H24 *float64 `json:"h24,omitempty"`
}

// ----------------------------------------------

func UnmarshalTokenInfoEthExplorer(data []byte) (TokenInfoEthExplorer, error) {
	var r TokenInfoEthExplorer
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokenInfoEthExplorer) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokenInfoEthExplorer struct {
	SimulationResult *SimulationResult `json:"simulationResult,omitempty"`
	HolderAnalysis   *HolderAnalysis   `json:"holderAnalysis,omitempty"`
}

type HolderAnalysis struct {
	Holders *string `json:"holders,omitempty"`
}

type TaxDistribution struct {
	Tax   *int64 `json:"tax,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

type SimulationResult struct {
	BuyTax  *float64 `json:"buyTax,omitempty"`
	SellTax *float64 `json:"sellTax,omitempty"`
}
