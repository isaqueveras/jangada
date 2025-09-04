// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// SailTransport defines the Sail transport structure.
type SailTransport struct {
	pathDir, folder, entity string
}

// Execute is the handler for the 'sail transport' command.
func (s *SailTransport) Execute(_ *cobra.Command, args []string) {
	folder, entity, layer := newSailTransportValidate(args...)
	mapperCreateLayerTransport[layer](&SailTransport{
		folder: folder, entity: entity, pathDir: s.pathDir,
	})
}

// createWebTransport generates the web transport layer structure.
func createWebTransport(app *SailTransport) {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("Creating web transport layer structure...\n\n")

	data := map[string]any{
		"folder": app.folder,
		"entity": app.entity,
		"layer":  "web",
	}

	for _, in := range TemplateTransport {
		time.Sleep(time.Second / 10)

		pathFile, err := createPath(in.Path, data)
		if err != nil {
			panic(err)
		}

		pathFullFile := fmt.Sprintf("%s%s", app.pathDir, pathFile)
		if err := createDir(pathFullFile); err != nil {
			panic(err)
		}

		if err = createFile(pathFullFile, in.Content, data); err != nil {
			panic(err)
		}

		log := color.New()
		log.Add(color.Reset, color.FgHiGreen, color.Bold).Print("\tcreate\t")
		log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
	}

	log.Add(color.Reset, color.Bold, color.FgHiBlue).Print("\nWeb transport layer structure created successfully!\n\n")
}

func createPath(key string, data map[string]any) (content string, err error) {
	templ, err := template.New("path").Parse(key)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := templ.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func createDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0755)
}

func createFile(path, content string, data map[string]any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var t *template.Template
	if t, err = template.New(path).Parse(content); err != nil {
		return err
	}

	if err := t.Execute(file, data); err != nil {
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
