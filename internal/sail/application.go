package sail

import (
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// Application is the handler for the 'jangada sail application' command.
func Application(cmd *cobra.Command, args []string) {
	cli.SetApplicationFlagService(cmd.Flag("service").Value.String())

	cfg := cli.GetConfig()
	folder, entity := getFolderAndEntityToTransport(args...)

	if cfg.ApplicationInfo.FlagService != "" {
		cmd.PrintErrln("not implemented yet")
		return
	}

	info := &info{
		Folder: folder,
		Entity: entity,
		Module: cli.GetModuleName(),
	}

	if err := createAllApplication(info); err != nil {
		cmd.PrintErrln(err)
		return
	}
}

func createAllApplication(data *info) error {
	return createFileForTemplate(data, applicationTemplate)
}
