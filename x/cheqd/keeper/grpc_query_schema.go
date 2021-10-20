package keeper

import (
	"context"

	"github.com/cheqd/cheqd-node/x/cheqd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Schema(c context.Context, req *types.QueryGetSchemaRequest) (*types.QueryGetSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasSchema(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	state, err := k.GetSchema(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetSchemaResponse{Schema: state.GetSchema()}, nil
}