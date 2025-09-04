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

// SailInterface defines the Sail interface structure.
type SailInterface struct {
	pathDir, folder, entity string
}

// Execute is the handler for the 'sail interface' command.
func (s *SailInterface) Execute(_ *cobra.Command, args []string) {
	folder, entity, layer := newSailInterfaceValidate(args...)
	mapperCreateLayerInterface[layer](&SailInterface{
		folder: folder, entity: entity, pathDir: s.pathDir,
	})
}

// createWebInterface generates the web interface layer structure.
func createWebInterface(app *SailInterface) {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("Creating web interface layer structure...\n\n")

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

	log.Add(color.Reset, color.Bold, color.FgHiBlue).Print("\nWeb interface layer structure created successfully!\n\n")
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

// createRestInterface generates the rest interface layer structure.
func createRestInterface(app *SailInterface) {
	log.Fatal("rest interface layer not implemented yet")
}

func createGRPCInterface(app *SailInterface) {
	log.Fatal("gRPC interface layer not implemented yet")
}

func createGraphQLInterface(app *SailInterface) {
	log.Fatal("GraphQL interface layer not implemented yet")
}

func createWebhookInterface(app *SailInterface) {
	log.Fatal("Webhook interface layer not implemented yet")
}

func createAllInterface(app *SailInterface) {
	log.Fatal("All interface layer not implemented yet")
}

// color.New(color.Bold, color.FgHiBlue).Println("Creating web interface layer structure...")
// color.Black(" - " + folderName + "/interface/web/handler.go")
// color.Black(" - " + folderName + "/interface/web/middleware.go")
// color.Black(" - " + folderName + "/interface/web/router.go\n\n")
