// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.
package spconfig

import (
	"github.com/charmbracelet/log"
	"github.com/sailpoint-oss/sailpoint-cli/internal/config"
	"github.com/sailpoint-oss/sailpoint-cli/internal/spconfig"
	"github.com/spf13/cobra"
)

func newDownloadCmd() *cobra.Command {
	var importIDs []string
	var exportIDs []string
	var folderPath string
	cmd := &cobra.Command{
		Use:     "download",
		Short:   "Download the results of import or export job from IdentityNow",
		Long:    "\nDownload the results of import or export job from IdentityNow\n\n",
		Example: "sail spconfig download -export <export job id> -import <import job id>",
		Aliases: []string{"down"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			apiClient, err := config.InitAPIClient()
			if err != nil {
				return err
			}

			if len(importIDs) > 0 {
				for i := 0; i < len(importIDs); i++ {
					jobId := importIDs[i]
					log.Info("Checking Import Job", "JobID", jobId)
					err := spconfig.DownloadImport(*apiClient, jobId, "spconfig-import-"+jobId+".json", folderPath)
					if err != nil {
						return err
					}
				}
			} else {
				log.Info("No Import Job IDs provided")
			}

			if len(exportIDs) > 0 {
				for i := 0; i < len(exportIDs); i++ {
					jobId := exportIDs[i]
					log.Info("Checking Export Job", "JobID", jobId)
					err := spconfig.DownloadExport(*apiClient, jobId, "spconfig-export-"+jobId+".json", folderPath)
					if err != nil {
						return err
					}
				}
			} else {
				log.Info("No Export Job IDs provided")
			}

			return nil
		},
	}

	cmd.Flags().StringArrayVarP(&importIDs, "import", "i", []string{}, "specify the IDs of the import jobs to download results for")
	cmd.Flags().StringArrayVarP(&exportIDs, "export", "e", []string{}, "specify the IDs of the export jobs to download results for")
	cmd.Flags().StringVarP(&folderPath, "folderPath", "f", "spconfig-exports", "folder path to save the search results in. If the directory doesn't exist, then it will be automatically created. (default is the current working directory)")

	return cmd
}
