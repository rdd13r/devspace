package list

import (
	"github.com/covexo/devspace/pkg/devspace/config/configutil"
	"github.com/covexo/devspace/pkg/util/log"
	"github.com/spf13/cobra"
)

type syncCmd struct{}

func newSyncCmd() *cobra.Command {
	cmd := &syncCmd{}

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Lists sync configuration",
		Long: `
	#######################################################
	################# devspace list sync ##################
	#######################################################
	Lists the sync configuration
	#######################################################
	`,
		Args: cobra.NoArgs,
		Run:  cmd.RunListSync,
	}

	return syncCmd
}

// RunListSync runs the list sync command logic
func (cmd *syncCmd) RunListSync(cobraCmd *cobra.Command, args []string) {
	config := configutil.GetConfig()

	if config.DevSpace.Sync == nil || len(*config.DevSpace.Sync) == 0 {
		log.Info("No sync paths are configured. Run `devspace add sync` to add new sync path\n")
		return
	}

	headerColumnNames := []string{
		"Service",
		"Selector",
		"Local Path",
		"Container Path",
		"Excluded Paths",
	}

	syncPaths := make([][]string, 0, len(*config.DevSpace.Sync))

	// Transform values into string arrays
	for _, value := range *config.DevSpace.Sync {
		service := ""
		selector := ""

		if value.Service != nil {
			service = *value.Service
		} else {
			for k, v := range *value.LabelSelector {
				if len(selector) > 0 {
					selector += ", "
				}

				selector += k + "=" + *v
			}
		}
		excludedPaths := ""

		if value.ExcludePaths != nil {
			for _, v := range *value.ExcludePaths {
				if len(excludedPaths) > 0 {
					excludedPaths += ", "
				}

				excludedPaths += v
			}
		}

		syncPaths = append(syncPaths, []string{
			service,
			selector,
			*value.LocalSubPath,
			*value.ContainerPath,
			excludedPaths,
		})
	}

	log.PrintTable(headerColumnNames, syncPaths)
}
