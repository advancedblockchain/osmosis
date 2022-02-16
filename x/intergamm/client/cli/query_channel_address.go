package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/osmosis-labs/osmosis/x/intergamm/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdChannelAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "channel-address [port] [channel]",
		Short: "Return the channel&#39;s account address if channel exists",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPort := args[0]
			reqChannel := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryChannelAddressRequest{

				Port:    reqPort,
				Channel: reqChannel,
			}

			res, err := queryClient.ChannelAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
