package keeper

import (
	"context"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Comments(goCtx context.Context, req *types.QueryCommentsRequest) (*types.QueryCommentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	var comments []*types.Comment

	store := ctx.KVStore(k.storeKey)
	commentStore := prefix.NewStore(store, []byte(types.CommentKey))

	post, _ := k.GetPost(ctx, req.Id)
	postID := post.Id

	pageRes, err := query.Paginate(commentStore, req.Pagination, func(key []byte, value []byte) error {
		var comment types.Comment
		if error := k.cdc.Unmarshal(value, &comment); error != nil {
			return error
		}
		if comment.PostID == postID {
			comments = append(comments, &comment)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryCommentsResponse{Post: &post, Comment: comments, Pagination: pageRes}, nil
}
