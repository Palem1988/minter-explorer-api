package data_resources

import (
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-tools/models"
)

type SetCandidate struct {
	PubKey string `json:"pub_key"`
}

func (SetCandidate) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.SetCandidateTxData)

	return SetCandidate{
		PubKey: data.PubKey,
	}
}
