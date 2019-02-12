package remove

import (
	"github.com/covexo/devspace/pkg/devspace/configure"
	"github.com/covexo/devspace/pkg/util/log"
	"github.com/spf13/cobra"
)

type imageCmd struct {
	RemoveAll bool
}

func newImageCmd() *cobra.Command {
	cmd := &imageCmd{}

	imageCmd := &cobra.Command{
		Use:   "image",
		Short: "Removes one or all images from the devspace",
		Long: `
	#######################################################
	############ devspace remove image ####################
	#######################################################
	Removes one or all images from a devspace:
	devspace remove image default
	devspace remove image --all
	#######################################################
	`,
		Args: cobra.MaximumNArgs(1),
		Run:  cmd.RunRemoveImage,
	}

	imageCmd.Flags().BoolVar(&cmd.RemoveAll, "all", false, "Remove all images")

	return imageCmd
}

// RunRemoveImage executes the remove image command logic
func (cmd *imageCmd) RunRemoveImage(cobraCmd *cobra.Command, args []string) {
	err := configure.RemoveImage(cmd.RemoveAll, args)
	if err != nil {
		log.Fatal(err)
	}
}
