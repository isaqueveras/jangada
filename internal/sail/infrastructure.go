package sail

import (
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// Infrastructure is the handler for the 'sail infrastructure' command.
func Infrastructure(cmd *cobra.Command, args []string) {
	folder, entity := getFolderAndEntityToTransport(args...)
	info := &info{Folder: folder, Entity: entity, Module: cli.GetModuleName()}
	if err := createFileForTemplate(info, infrastructureTemplate); err != nil {
		cmd.PrintErrln(err)
	}
}
