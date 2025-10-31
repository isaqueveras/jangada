package sail

import (
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// Application is the handler for the 'jangada sail application' command.
func Application(cmd *cobra.Command, args []string) {
	cli.SetApplicationFlagService(cmd.Flag("service").Value.String())

	cfg := cli.GetConfig()
	if cfg.ApplicationInfo.FlagService != "" {
		cmd.PrintErrln("not implemented yet")
		return
	}

	folder, entity := getFolderAndEntityToTransport(args...)
	info := &info{Folder: folder, Entity: entity, Module: cli.GetModuleName()}
	if err := createFileForTemplate(info, applicationTemplate); err != nil {
		cmd.PrintErrln(err)
	}
}
