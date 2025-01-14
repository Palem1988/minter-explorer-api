package data_resources

import (
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-tools/models"
)

type Send struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

func (Send) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.SendTxData)

	return Send{
		Coin:  data.Coin,
		To:    data.To,
		Value: helpers.PipStr2Bip(data.Value),
	}
}
