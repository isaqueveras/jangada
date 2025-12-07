package sail

import (
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// Domain is the handler for the 'sail domain' command.
func Domain(cmd *cobra.Command, args []string) {
	folder, entity := getFolderAndEntityToTransport(args...)
	info := &info{Folder: folder, Entity: entity, Module: cli.GetModuleName()}
	if err := createFileForTemplate(info, domainTemplate, nil); err != nil {
		cmd.PrintErrln(err)
	}
}
