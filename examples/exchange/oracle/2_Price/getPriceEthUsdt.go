package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
  getPrice("ETH", "USDT", "mainnet")
  getPrice("ETH", "USDT", "testnet")
}

func getPrice(_baseSymbol string, _quoteSymbol string, _network string) {
	network := common.LoadNetwork(_network, "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	baseSymbol := _baseSymbol
	quoteSymbol := _quoteSymbol
	oracleType := "BandIBC"
	oracleScaleFactor := uint32(6)
	res, err := exchangeClient.GetPrice(ctx, baseSymbol, quoteSymbol, oracleType, oracleScaleFactor)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print("price ", _baseSymbol, "/",  _quoteSymbol, "(", _network, "): ", string(str))
}
