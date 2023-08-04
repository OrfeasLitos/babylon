package keeper

import (
	"fmt"

	"github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBTCValidator adds the given BTC validator to KVStore
func (k Keeper) SetBTCValidator(ctx sdk.Context, btcVal *types.BTCValidator) {
	store := k.btcValidatorStore(ctx)
	btcValBytes := k.cdc.MustMarshal(btcVal)
	store.Set(btcVal.BtcPk.MustMarshal(), btcValBytes)
}

// HasBTCValidator checks if the BTC validator exists
func (k Keeper) HasBTCValidator(ctx sdk.Context, valBTCPK []byte) bool {
	store := k.btcValidatorStore(ctx)
	return store.Has(valBTCPK)
}

// GetBTCValidator gets the BTC validator with the given validator Bitcoin PK
func (k Keeper) GetBTCValidator(ctx sdk.Context, valBTCPK []byte) (*types.BTCValidator, error) {
	store := k.btcValidatorStore(ctx)
	if !k.HasBTCValidator(ctx, valBTCPK) {
		return nil, types.ErrBTCValNotFound
	}
	btcValBytes := store.Get(valBTCPK)
	var btcVal types.BTCValidator
	k.cdc.MustUnmarshal(btcValBytes, &btcVal)
	return &btcVal, nil
}

// SlashBTCValidator slashes a BTC validator with the given PK
// A slashed BTC validator will not have voting power
func (k Keeper) SlashBTCValidator(ctx sdk.Context, valBTCPK []byte) error {
	btcVal, err := k.GetBTCValidator(ctx, valBTCPK)
	if err != nil {
		return err
	}
	if btcVal.IsSlashed() {
		return types.ErrBTCValAlreadySlashed
	}
	btcVal.SlashedBabylonHeight = uint64(ctx.BlockHeight())
	btcTip := k.btclcKeeper.GetTipInfo(ctx)
	if btcTip == nil {
		panic(fmt.Errorf("failed to get current BTC tip"))
	}
	btcVal.SlashedBtcHeight = btcTip.Height
	k.SetBTCValidator(ctx, btcVal)
	return nil
}

// btcValidatorStore returns the KVStore of the BTC validator set
// prefix: BTCValidatorKey
// key: Bitcoin secp256k1 PK
// value: BTCValidator object
func (k Keeper) btcValidatorStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.BTCValidatorKey)
}
