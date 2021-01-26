package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/user/nameservice/x/nameservice/types"
)

// GetNameCount get the total number of Name
func (k Keeper) GetNameCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.NameCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetNameCount set the total number of Name
func (k Keeper) SetNameCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.NameCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateName creates a Name
func (k Keeper) CreateName(ctx sdk.Context, msg types.MsgBuyName) {
	// Create the Name
	count := k.GetNameCount(ctx)
	var Name = types.Name{
		Owner: msg.Owner,
		ID:    strconv.FormatInt(count, 10),
		Value: msg.Value,
		Price: msg.Price,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.NamePrefix + Name.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(Name)
	store.Set(key, value)

	// Update Name count
	k.SetNameCount(ctx, count+1)
}

// GetName returns the Name information
func (k Keeper) GetName(ctx sdk.Context, key string) (types.Name, error) {
	store := ctx.KVStore(k.storeKey)
	var Name types.Name
	byteKey := []byte(types.NamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &Name)
	if err != nil {
		return Name, err
	}
	return Name, nil
}

// SetName sets a Name
func (k Keeper) SetName(ctx sdk.Context, Name types.Name) {
	NameKey := Name.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(Name)
	key := []byte(types.NamePrefix + NameKey)
	store.Set(key, bz)
}

// DeleteName deletes a Name
func (k Keeper) DeleteName(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.NamePrefix + key))
}

//
// Functions used by querier
//

func listName(ctx sdk.Context, k Keeper) ([]byte, error) {
	var NameList []types.Name
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.NamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var Name types.Name
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &Name)
		NameList = append(NameList, Name)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, NameList)
	return res, nil
}

func getName(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	Name, err := k.GetName(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, Name)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetNameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	Name, err := k.GetName(ctx, key)
	if err != nil {
		return nil
	}
	return Name.Owner
}

// Check if the key exists in the store
func (k Keeper) NameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.NamePrefix + key))
}
