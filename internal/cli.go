package cli

import (
	"path/filepath"

	"github.com/isaqueveras/jangada/internal/sail"
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

// GetDirectoryPath returns the directory path.
func (j *Jangada) GetDirectoryPath() string {
	return j.directoryPath
}

func New() error {
	dirBase, _ := filepath.Abs("")
	jang := &Jangada{dirBase: dirBase}
	jang.setFullDirectoryPath()

	jangada := &cobra.Command{
		Use:     "jangada",
		Short:   "Jangada is a CLI tool for project scaffolding and code generation.",
		Example: "jangada new my-app --module github.com/username/my-app --database postgres",
	}

	jangada.AddCommand(jang.commandNewProject())
	jangada.AddCommand(sail.NewCommand().Execute(jang.GetDirectoryPath()))

	return jangada.Execute()
}
