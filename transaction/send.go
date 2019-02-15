package transaction

import (
	"encoding/json"
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-extender/models"
)

type SendDataResource struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

func (SendDataResource) Transform(txData resource.ItemInterface) resource.Interface {
	data := txData.(models.SendData)

	return SendDataResource{
		Coin:  data.Coin,
		To:    data.To,
		Value: helpers.PipStr2Bip(data.Value),
	}
}

func (resource SendDataResource) TransformFromJsonRaw(raw json.RawMessage) resource.Interface {
	var data models.SendData

	err := json.Unmarshal(raw, &data)
	helpers.CheckErr(err)

	return resource.Transform(data)
}
