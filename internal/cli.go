package cli

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

type Jangada struct {
	AppName         string
	ModuleName      string
	DefaultAddrPort string
	Database        string

	dirBase       string
	directoryPath string
}

func New() error {
	dirBase, _ := filepath.Abs("")
	jang := &Jangada{dirBase: dirBase}
	jang.setFullDirectoryPath()

	root := jang.Root()
	root.AddCommand(jang.commandNewProject())

	return root.Execute()
}

func (*Jangada) Root() *cobra.Command {
	return &cobra.Command{
		Use:     "jangada",
		Short:   "Jangada is a CLI tool for project scaffolding and code generation.",
		Example: "jangada new my-app",
	}
}

func (jang *Jangada) CreateCoreFiles()  {}
func (jang *Jangada) CreateTestsFiles() {}
