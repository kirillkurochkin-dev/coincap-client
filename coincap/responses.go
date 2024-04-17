package coincap

import "fmt"

type assetsResponse struct {
	Data      []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type assetResponse struct {
	Data      AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}

type AssetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24H string `json:"volumeUSD24H"`
	PriceUSD     string `json:"priceUSD"`
}

func (a AssetData) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s",
		a.ID, a.Rank, a.Symbol, a.Name, a.PriceUSD)
}
