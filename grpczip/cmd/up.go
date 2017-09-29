package cmd

import (
	"github.com/lileio/lile"
	"github.com/sago35/grpczip/subscribers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC and pubub subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logrus.Fatal(lile.Serve())

		}()

		lile.Subscribe(&subscribers.GrpczipServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
