package keeper

import (
	"encoding/binary"
	"github.com/cheqd/cheqd-node/x/cheqd/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// GetDidCount get the total number of did
func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidCountKey))
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDidCount set the total number of did
func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidCountKey))
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDid appends a did in the store with a new id and update the count
func (k Keeper) AppendDid(
	ctx sdk.Context,
	creator string,
	verkey string,
	alias string,
) uint64 {
	// Create the did
	count := k.GetDidCount(ctx)
	var did = types.Did{
		Creator: creator,
		Id:      count,
		Verkey:  verkey,
		Alias:   alias,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	value := k.cdc.MustMarshalBinaryBare(&did)
	store.Set(GetDidIDBytes(did.Id), value)

	// Update did count
	k.SetDidCount(ctx, count+1)

	return count
}

// SetDid set a specific did in the store
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshalBinaryBare(&did)
	store.Set(GetDidIDBytes(did.Id), b)
}

// GetDid returns a did from its id
func (k Keeper) GetDid(ctx sdk.Context, id uint64) types.Did {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	var did types.Did
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDidIDBytes(id)), &did)
	return did
}

// HasDid checks if the did exists in the store
func (k Keeper) HasDid(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	return store.Has(GetDidIDBytes(id))
}

// GetDidOwner returns the creator of the did
func (k Keeper) GetDidOwner(ctx sdk.Context, id uint64) string {
	return k.GetDid(ctx, id).Creator
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveDid(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	store.Delete(GetDidIDBytes(id))
}

// GetAllDid returns all did
func (k Keeper) GetAllDid(ctx sdk.Context) (list []types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Did
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDidIDBytes returns the byte representation of the ID
func GetDidIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDidIDFromBytes returns ID in uint64 format from a byte array
func GetDidIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}