// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

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
	for path, content := range TemplateMapInterface {
		data := map[string]any{
			"folder": app.folder,
			"entity": app.entity,
		}

		pathFile, err := createPath(path, data)
		if err != nil {
			log.Fatal("Error create path for web interface layer files")
		}

		pathFile = fmt.Sprintf("%s%s", app.pathDir, pathFile)
		if err := createDir(pathFile); err != nil {
			log.Fatal("Error creating directory for web interface layer files: ", err)
		}

		if err = createFile(pathFile, content, data); err != nil {
			log.Fatal("Error creating web interface layer files: ", err)
		}
	}
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
