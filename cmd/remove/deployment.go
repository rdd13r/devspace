package remove

import (
	"github.com/covexo/devspace/pkg/devspace/configure"
	"github.com/covexo/devspace/pkg/util/log"
	"github.com/spf13/cobra"
)

type deploymentCmd struct {
	RemoveAll bool
}

func newDeploymentCmd() *cobra.Command {
	cmd := &deploymentCmd{}

	deploymentCmd := &cobra.Command{
		Use:   "deployment",
		Short: "Removes one or all deployments from the devspace",
		Long: `
	#######################################################
	############ devspace remove deployment ###############
	#######################################################
	Removes one or all deployments from a devspace:
	devspace remove deployment devspace-default
	devspace remove deployment --all
	#######################################################
	`,
		Args: cobra.MaximumNArgs(1),
		Run:  cmd.RunRemoveDeployment,
	}

	deploymentCmd.Flags().BoolVar(&cmd.RemoveAll, "all", false, "Remove all deployments")

	return deploymentCmd
}

// RunRemoveDeployment executes the specified deployment
func (cmd *deploymentCmd) RunRemoveDeployment(cobraCmd *cobra.Command, args []string) {
	name := ""
	if len(args) > 0 {
		name = args[0]
	}

	err := configure.RemoveDeployment(cmd.RemoveAll, name)
	if err != nil {
		log.Fatal(err)
	}
}
