// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// SailTransport defines the Sail transport structure.
type SailTransport struct {
	pathDir, folder, entity, module string
}

// transportCommand ...
func transportCommand() *cobra.Command {
	cmd := &SailTransport{
		pathDir: cli.GetDirectoryPath(),
		module:  cli.GetModuleName(),
	}

	return &cobra.Command{
		Use:     "transport",
		Short:   "Create transport layer",
		Args:    cobra.RangeArgs(1, 2),
		Example: exampleCreateTransportText,
		Run:     cmd.Execute,
	}
}

// Execute is the handler for the 'sail transport' command.
func (s *SailTransport) Execute(_ *cobra.Command, args []string) {
	folder, entity, layer := newSailTransportValidate(args...)
	mapperCreateLayerTransport[layer](&SailTransport{
		folder:  folder,
		entity:  entity,
		pathDir: s.pathDir,
		module:  s.module,
	})
}

type webTransportTemplateData struct {
	Folder string
	Entity string
	Module string
	Layer  TransportLayer
}

// createWebTransport generates the web transport layer structure.
func createWebTransport(app *SailTransport) {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("Creating web transport layer structure...\n\n")

	data := &webTransportTemplateData{
		Folder: app.folder,
		Entity: app.entity,
		Module: cli.GetModuleName(),
		Layer:  WebTransportLayer,
	}

	for _, in := range WebTransportTemplate {
		pathFile, err := createPath(in.Path, data)
		if err != nil {
			panic(err)
		}

		if err := createDir(pathFile); err != nil {
			panic(err)
		}

		data.Entity = cli.Capitalize(app.entity)
		if err = createFile(pathFile, in.Content, data); err != nil {
			panic(err)
		}

		log := color.New()
		log.Add(color.Reset, color.FgHiGreen, color.Bold).Print("\tcreate\t")
		log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
	}

	log.Add(color.Reset, color.Bold, color.FgHiBlue).Print("\nWeb transport layer structure created successfully!\n\n")
}

func createPath(key string, data *webTransportTemplateData) (content string, err error) {
	templ, err := template.New("path").Parse(key)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := templ.Execute(&buf, &data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func createDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0755)
}

func createFile(path, content string, data *webTransportTemplateData) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var t *template.Template
	if t, err = template.New(path).Parse(content); err != nil {
		return err
	}

	if err := t.Execute(file, &data); err != nil {
		return err
	}

	return nil
}

// createRestTransport generates the rest transport layer structure.
func createRestTransport(app *SailTransport) {
	log.Fatal("rest transport layer not implemented yet")
}

func createGRPCTransport(app *SailTransport) {
	log.Fatal("gRPC transport layer not implemented yet")
}

func createGraphQLTransport(app *SailTransport) {
	log.Fatal("GraphQL transport layer not implemented yet")
}

func createWebhookTransport(app *SailTransport) {
	log.Fatal("Webhook transport layer not implemented yet")
}

func createAllTransport(app *SailTransport) {
	log.Fatal("All transport layer not implemented yet")
}
