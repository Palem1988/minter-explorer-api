package transaction

import (
	"github.com/MinterTeam/minter-explorer-api/blocks"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// TODO: replace string in StartBlock, EndBlock to int
type SelectFilter struct {
	Addresses       []string
	BlockId         *uint64
	StartBlock      *string
	EndBlock        *string
	ValidatorPubKey *string
}

func (f *SelectFilter) Filter(q *orm.Query) (*orm.Query, error) {
	if len(f.Addresses) > 0 {
		q = q.WhereIn("from_address.address IN (?)", pg.In(f.Addresses)).
			WhereOr("tx_output__to_address.address IN (?)", pg.In(f.Addresses))
	}

	if f.ValidatorPubKey != nil {
		q = q.Join("JOIN transaction_validator").
			JoinOn("transaction_validator.transaction_id = transaction.id").
			Join("JOIN validators").
			JoinOn("validators.public_key = ?", f.ValidatorPubKey)
	}

	if f.BlockId != nil {
		q = q.Where("transaction.block_id = ?", f.BlockId)
	}

	blocksRange := blocks.RangeSelectFilter{StartBlock: f.StartBlock, EndBlock: f.EndBlock}
	q = q.Apply(blocksRange.Filter)

	return q, nil
}