package newapp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	cli "github.com/isaqueveras/jangada/internal"
	temp "github.com/isaqueveras/jangada/internal/template"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "new [name]",
		Short: "Create a new app",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}
}

func execute(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	cli.SetAppName(args[0])

	setFlagsNewProject(cmd)
	createRootDir()
	createRootFiles()
	gitInit()

	generateTemplate()
	execGoModTidy()
	copyStaticFiles()

	log := color.New()
	log.Add(color.Reset, color.Bold, color.BgGreen, color.FgWhite).
		Print("\n\t🛶 Jangada (v0.1.0-beta)")
	log.Add(color.Reset, color.BgHiBlack, color.FgHiWhite).
		Print("\n\tLightning-fast web framework for Golang focused on simplicity and productivity.")
	log.Println("\n\tIt helps developers quickly scaffold, build, and manage modern web applications with minimal configuration.")
	fmt.Println("")
}

func setFlagsNewProject(cmd *cobra.Command) {
	defaultAddrPort, _ := cmd.Flags().GetString("port")
	cli.SetDefaultAddrPort(defaultAddrPort)

	moduleName, _ := cmd.Flags().GetString("module")
	cli.SetModuleName(moduleName)

	database, _ := cmd.Flags().GetString("database")
	cli.SetDatabase(database)
}

func createRootDir() {
	cli.SetFullDirectoryPath()
	fmt.Printf("cli.GetDirectoryPath(): %v\n", cli.GetDirectoryPath())
	if err := os.MkdirAll(cli.GetDirectoryPath(), 0755); err != nil {
		panic(err)
	}
}

func createRootFiles() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInstalling the project...\n\n")

	for path, tmpl := range temp.GetTemplateForNewApp() {
		cli.CreateFile(path, tmpl)
	}
}

// gitInit ...
func gitInit() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInitializing git repository\n\n")
	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Println("git init")
}

func generateTemplate() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nGenerating template...\n\n")
	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Println(`go run github.com/a-h/templ/cmd/templ@latest generate`)

	command := exec.Command("bash", "-c", "go run github.com/a-h/templ/cmd/templ@latest generate")
	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Run(); err != nil {
		panic(err)
	}
}

func execGoModTidy() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInstalling dependencies...\n\n")

	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Print("go mod tidy\n\n")

	cmd := fmt.Sprintf("cd %s && go mod tidy", cli.GetDirectoryPath())
	command := exec.Command("bash", "-c", cmd)
	command.Stdout, command.Stderr = os.Stdout, os.Stderr

	if err := command.Run(); err != nil {
		panic(err)
	}
}

func copyStaticFiles() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nCopying static files\n\n")

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Add(color.FgHiRed, color.Bold).Print("\tError: ")
		log.Add(color.Reset).Printf("could not get file path")
		return
	}

	src := fmt.Sprintf("%s/static/background.png", filepath.Dir(cli.RemoveLastSegment(filename)))
	dst := fmt.Sprintf("%s/public/background.png", cli.GetDirectoryPath())

	if err := cli.CopyFile(src, dst); err != nil {
		log.Add(color.FgHiRed, color.Bold).Print("\tError: ")
		log.Add(color.Reset).Printf("copying static file: %s\n", err)
		return
	}

	log.Add(color.FgHiGreen, color.Bold).Print("\tCopied\t")
	log.Add(color.FgHiWhite, color.Reset).Printf("%s\n", dst)
}
