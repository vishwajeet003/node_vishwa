package client

import (
	"errors"
	"fmt"
	"reflect"

	aclient "github.com/akash-network/akash-api/go/node/client"
	"github.com/akash-network/akash-api/go/node/client/v1beta2"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	cmtrpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

var (
	ErrInvalidClient = errors.New("invalid client")
)

func GetClientQueryContext(cmd *cobra.Command) (v1beta2.Query, error) {
	cctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, err
	}

	var cl v1beta2.Query
	err = aclient.DiscoverQueryClient(cmd.Context(), cctx, func(i interface{}) error {
		var valid bool

		if cl, valid = i.(v1beta2.Query); !valid {
			return fmt.Errorf("%w: expected %s, actual %s", ErrInvalidClient, reflect.TypeOf(cl), reflect.TypeOf(i))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return cl, nil
}

func RPCAkash(_ *cmtrpctypes.Context) (*aclient.Akash, error) {
	result := &aclient.Akash{
		ClientInfo: &aclient.ClientInfo{
			ApiVersion: "v1beta2",
		},
	}

	return result, nil
}
