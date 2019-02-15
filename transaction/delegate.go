package transaction

import (
	"encoding/json"
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-extender/models"
)

type DelegateDataResource struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Stake  string `json:"stake"`
}

func (DelegateDataResource) Transform(txData resource.ItemInterface) resource.Interface {
	data := txData.(models.DelegateData)

	return DelegateDataResource{
		PubKey: data.PubKey,
		Coin:   data.Coin,
		Stake:  helpers.PipStr2Bip(data.Stake),
	}
}

func (resource DelegateDataResource) TransformFromJsonRaw(raw json.RawMessage) resource.Interface {
	var data models.DelegateData
	err := json.Unmarshal(raw, &data)
	helpers.CheckErr(err)

	return resource.Transform(data)
}