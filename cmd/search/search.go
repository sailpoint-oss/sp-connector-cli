// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.
package search

import (
	"fmt"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
	"github.com/spf13/cobra"
)

func NewSearchCmd(apiClient *sailpoint.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search",
		Short:   "perform search in identitynow with a search string",
		Long:    "Search IdentityNow with a provided search string",
		Example: "sail search \"\"",
		Aliases: []string{"se"},
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprint(cmd.OutOrStdout(), cmd.UsageString())
		},
	}

	cmd.AddCommand(
		newQueryCmd(apiClient),
		newTemplateCmd(apiClient),
	)

	return cmd

}
