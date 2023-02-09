package root

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newDebugCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "debug",
		Short:   "enable or disable debug mode",
		Long:    "Enable/Disable debug mode.",
		Example: "sail debug enable | disable",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			switch strings.ToLower(args[0]) {
			case "enable":
				viper.Set("debug", true)
			case "disable":
				viper.Set("debug", false)
			}

			return nil
		},
	}
	return cmd

}
