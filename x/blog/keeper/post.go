package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	_, _ = ctx, post

	count := k.GetPostCount(ctx)
	post.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(byteKey, appendedValue)

	k.SetPostCount(ctx, count+1)
	return 1
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))

	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
