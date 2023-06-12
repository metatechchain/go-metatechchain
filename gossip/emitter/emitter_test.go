package emitter

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/metatechchain/meta-base/hash"
	"github.com/metatechchain/meta-base/inter/idx"
	"github.com/metatechchain/meta-base/inter/pos"
	"github.com/stretchr/testify/require"

	"github.com/metatechchain/go-metatechchain/gossip/emitter/mock"
	"github.com/metatechchain/go-metatechchain/integration/makegenesis"
	"github.com/metatechchain/go-metatechchain/inter"
	"github.com/metatechchain/go-metatechchain/metatechchain"
	"github.com/metatechchain/go-metatechchain/vecmt"
)

//go:generate go run github.com/golang/mock/mockgen -package=mock -destination=mock/world.go github.com/metatechchain/go-metatechchain/gossip/emitter External,TxPool,TxSigner,Signer

func TestEmitter(t *testing.T) {
	cfg := DefaultConfig()
	gValidators := makegenesis.GetFakeValidators(3)
	vv := pos.NewBuilder()
	for _, v := range gValidators {
		vv.Set(v.ID, pos.Weight(1))
	}
	validators := vv.Build()
	cfg.Validator.ID = gValidators[0].ID

	ctrl := gomock.NewController(t)
	external := mock.NewMockExternal(ctrl)
	txPool := mock.NewMockTxPool(ctrl)
	signer := mock.NewMockSigner(ctrl)
	txSigner := mock.NewMockTxSigner(ctrl)

	external.EXPECT().Lock().
		AnyTimes()
	external.EXPECT().Unlock().
		AnyTimes()
	external.EXPECT().DagIndex().
		Return((*vecmt.Index)(nil)).
		AnyTimes()
	external.EXPECT().IsSynced().
		Return(true).
		AnyTimes()
	external.EXPECT().PeersNum().
		Return(int(3)).
		AnyTimes()

	em := NewEmitter(cfg, World{
		External: external,
		TxPool:   txPool,
		Signer:   signer,
		TxSigner: txSigner,
	})

	t.Run("init", func(t *testing.T) {
		external.EXPECT().GetRules().
			Return(metatechchain.FakeNetRules()).
			AnyTimes()

		external.EXPECT().GetEpochValidators().
			Return(validators, idx.Epoch(1)).
			AnyTimes()

		external.EXPECT().GetLastEvent(idx.Epoch(1), cfg.Validator.ID).
			Return((*hash.Event)(nil)).
			AnyTimes()

		external.EXPECT().GetGenesisTime().
			Return(inter.Timestamp(uint64(time.Now().UnixNano()))).
			AnyTimes()

		em.init()
	})

	t.Run("memorizeTxTimes", func(t *testing.T) {
		require := require.New(t)
		tx := types.NewTransaction(1, common.Address{}, big.NewInt(1), 1, big.NewInt(1), nil)

		external.EXPECT().IsBusy().
			Return(true).
			AnyTimes()

		_, ok := em.txTime.Get(tx.Hash())
		require.False(ok)

		before := time.Now()
		em.memorizeTxTimes(types.Transactions{tx})
		after := time.Now()

		cached, ok := em.txTime.Get(tx.Hash())
		got := cached.(time.Time)
		require.True(ok)
		require.True(got.After(before))
		require.True(got.Before(after))
	})

	t.Run("tick", func(t *testing.T) {
		em.tick()
	})
}
