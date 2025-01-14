package data_resources

import (
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/MinterTeam/minter-explorer-tools/models"
)

type EditCandidate struct {
	PubKey        string `json:"pub_key"`
	RewardAddress string `json:"reward_address"`
	OwnerAddress  string `json:"owner_address"`
}

func (EditCandidate) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.EditCandidateTxData)

	return EditCandidate{
		PubKey:        data.PubKey,
		RewardAddress: data.RewardAddress,
		OwnerAddress:  data.OwnerAddress,
	}
}
