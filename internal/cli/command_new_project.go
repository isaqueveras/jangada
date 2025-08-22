package cli

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/isaqueveras/jangada/internal/cli/tmpl"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func (jang *Jangada) commandNewProject() *cobra.Command {
	return &cobra.Command{
		Use:   "new [name]",
		Short: "Create a new project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			log := color.New()
			jang.AppName = args[0]

			jang.getFlagsNewProject(cmd)
			jang.validate()
			jang.createRootDir()
			jang.createRootFiles()
			jang.GitInit()
			jang.GenerateTemplate()
			jang.ExecGoModTidy()

			jang.CopyStaticFiles()

			// jang.BuildApp()

			log.Add(color.Reset, color.Bold, color.BgGreen, color.FgWhite).
				Print("\n\t🛶 Jangada (v0.1.0-beta)")
			log.Add(color.Reset, color.BgHiBlack, color.FgHiWhite).
				Print("\n\tLightning-fast web framework for Golang focused on simplicity and productivity.")
			log.Println("\n\tIt helps developers quickly scaffold, build, and manage modern web applications with minimal configuration.")
			fmt.Println("")
		},
	}
}

func (jang *Jangada) validate() {
	if jang.ModuleName == "" {
		jang.ModuleName = jang.AppName
	}

	if jang.DefaultAddrPort == "" {
		jang.DefaultAddrPort = ":8080"
	}
}

func (jang *Jangada) getFlagsNewProject(cmd *cobra.Command) {
	jang.DefaultAddrPort, _ = cmd.Flags().GetString("port")
	jang.ModuleName, _ = cmd.Flags().GetString("module")
	jang.Database, _ = cmd.Flags().GetString("database")
}

func (jang *Jangada) createRootDir() {
	jang.setFullDirectoryPath()
	if err := os.MkdirAll(jang.directoryPath, 0755); err != nil {
		panic(err)
	}
}

func (jang *Jangada) setFullDirectoryPath() {
	jang.directoryPath = jang.dirBase + "/" + strings.ToLower(jang.AppName)
}

func (jang *Jangada) createRootFiles() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInstalling the project...\n\n")

	for name, tmpl := range tmpl.GetTemplateForNewProject() {
		jang.createFile(name, tmpl.String())
	}
}

func (jang *Jangada) createFile(name, tmpl string) {
	dir := fmt.Sprintf("%s/%s", jang.directoryPath, filepath.Dir(name))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}

	log := color.New()

	pathFile := fmt.Sprintf("%s/%s", jang.directoryPath, name)
	if _, err := os.Stat(pathFile); !os.IsNotExist(err) {
		log.Add(color.FgHiMagenta, color.Bold).Print("\texist\t")
		log.Add(color.Reset).Printf("%s\n", name)
		return
	}

	file, err := os.Create(pathFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var t *template.Template
	if t, err = template.New(jang.dirBase).Parse(tmpl); err != nil {
		panic(err)
	}

	if err := t.Execute(file, jang); err != nil {
		panic(err)
	}

	log.Add(color.FgHiGreen, color.Bold).Print("\tcreate\t")
	log.Add(color.FgHiWhite, color.Reset).Printf("%s\n", name)
}

func (jang *Jangada) ExecGoModTidy() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInstalling dependencies...\n\n")

	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Print("go mod tidy\n\n")

	cmd := fmt.Sprintf("cd %s && go mod tidy", jang.directoryPath)
	command := exec.Command("bash", "-c", cmd)
	command.Stdout, command.Stderr = os.Stdout, os.Stderr

	if err := command.Run(); err != nil {
		panic(err)
	}
}

func (jang *Jangada) GitInit() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nInitializing git repository\n\n")
	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Println("git init")
}

func (jang *Jangada) GenerateTemplate() {
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

func (jang *Jangada) BuildApp() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nBuilding the project...\n\n")
	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Println("go build")

	cmd := fmt.Sprintf("go build -o ./bin/app-%s %s/%s/cmd/app", jang.AppName, jang.dirBase, jang.AppName)

	fmt.Printf("cmd: %v\n", cmd)

	command := exec.Command("bash", "-c", cmd)
	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Run(); err != nil {
		panic(err)
	}
}

func (jang *Jangada) CopyStaticFiles() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nCopying static files\n\n")

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Add(color.FgHiRed, color.Bold).Print("\tError: ")
		log.Add(color.Reset).Printf("could not get file path")
		return
	}

	src := fmt.Sprintf("%s/static/background.png", filepath.Dir(filename))
	dst := fmt.Sprintf("%s/public/background.png", jang.directoryPath)

	if err := copyFile(src, dst); err != nil {
		log.Add(color.FgHiRed, color.Bold).Print("\tError: ")
		log.Add(color.Reset).Printf("copying static file: %s\n", err)
		return
	}

	log.Add(color.FgHiGreen, color.Bold).Print("\tCopied\t")
	log.Add(color.FgHiWhite, color.Reset).Printf("%s\n", dst)
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(data)
	return err
}
