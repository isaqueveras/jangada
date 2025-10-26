// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/fatih/color"
	cli "github.com/isaqueveras/jangada/internal"
	"github.com/spf13/cobra"
)

// SailTransport defines the Sail transport structure.
type SailTransport struct {
	pathDir, folder, entity, module, layer string
}

// Execute is the handler for the 'sail transport' command.
func Execute(cmd *cobra.Command, args []string) {
	cli.SetFlagTransportLayer(cmd.Flag("layer").Value.String())
	cli.SetFlagCRUD(cmd.Flag("crud").Value.String() == "true")
	cli.SetFlagMethod(cmd.Flag("method").Value.String())

	folder, entity := getFolderAndEntityToTransport(args...)

	cfg := cli.GetConfig()
	st := &SailTransport{
		folder:  folder,
		entity:  entity,
		pathDir: cfg.DirectoryPath,
		module:  cfg.ModuleName,
		layer:   cfg.TransportInfo.TransportLayer,
	}

	mapperCreateLayerTransport[TransportLayer(cfg.TransportInfo.TransportLayer)](st)
}

type webTransportTemplateData struct {
	Folder, Entity, Module, Method string
	Layer                          TransportLayer
}

// createTransport generates the transport layer structure.
func createTransport(st *SailTransport) {
	var (
		cfg  = cli.GetConfig()
		log  = color.New()
		data = &webTransportTemplateData{
			Folder: st.folder,
			Entity: strings.ToLower(st.entity),
			Module: cli.GetModuleName(),
			Layer:  TransportLayer(st.layer),
			Method: cfg.TransportInfo.FlagCreateMethod,
		}
	)

	if !cfg.TransportInfo.FlagCreateCRUD && cfg.TransportInfo.FlagCreateMethod == "" {
		log.Add(color.Reset, color.FgHiRed, color.Bold).Println("You must use --crud or --method flag.")
		return
	}

	// log.Add(color.Bold, color.FgHiBlue).
	// 	Print("Creating " + st.layer + " transport layer structure...\n\n")

	switch {
	case cfg.TransportInfo.FlagCreateCRUD:
		if err := createFileTransport(data, transportTemplateCRUD); err != nil {
			panic(err)
		}
	case cfg.TransportInfo.FlagCreateMethod != "":
		if err := createFileTransport(data, transportTemplateCreateMethod); err != nil {
			panic(err)
		}
	}

	// color.New().Add(color.Reset, color.Bold, color.FgHiBlue).
	// 	Print("\n" + st.layer + " transport layer structure created successfully!\n\n")
}

func createFileTransport(data *webTransportTemplateData, templates []Template) error {
	for _, in := range templates {
		pathFile, err := createPath(in.Path, data)
		if err != nil {
			return err
		}

		if err := createDir(pathFile); err != nil {
			return err
		}

		log := color.New()
		if _, err := os.Stat(pathFile); !os.IsNotExist(err) {
			if !in.CanModify {
				log.Add(color.Reset, color.FgHiMagenta, color.Bold).Print("\texist\t")
				log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
				continue
			}

			log.Add(color.Reset, color.FgHiYellow, color.Bold).Print("\tupdate\t")
			log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
			return updateFile(pathFile, in, data)
		}

		if err = createFile(pathFile, in.Content, data); err != nil {
			return err
		}

		log.Add(color.Reset, color.FgHiGreen, color.Bold).Print("\tcreated\t")
		log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
	}

	return nil
}

func createPath(key string, data *webTransportTemplateData) (content string, err error) {
	funcs := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	templ, err := template.New("path").Funcs(funcs).Parse(key)
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
	return os.MkdirAll(filepath.Dir(strings.ToLower(path)), 0755)
}

func updateFile(path string, in Template, data *webTransportTemplateData) (err error) {
	if !strings.Contains(in.Path, "_controller.go") {
		return nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tmpFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	updated := (string(content) + "\n" + in.Content)
	if err = createTemplateParser(path, updated, tmpFile, data); err != nil {
		return err
	}

	return tmpFile.Close()
}

func createFile(path, content string, data *webTransportTemplateData) error {
	file, err := os.Create(strings.ToLower(path))
	if err != nil {
		return err
	}
	defer file.Close()

	return createTemplateParser(path, content, file, data)
}

func createTemplateParser(path, content string, file *os.File, data *webTransportTemplateData) (err error) {
	funcs := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	var t *template.Template
	if t, err = template.New(path).Funcs(funcs).Parse(content); err != nil {
		return err
	}

	data.Entity = cli.Capitalize(data.Entity)
	return t.Execute(file, &data)
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

func getFolderAndEntityToTransport(args ...string) (string, string) {
	folder := args[0]
	if folder == "" {
		log.Fatal("You must provide a folder name for the bounded context.")
	}

	folderParts := strings.Split(folder, "/")
	if len(folderParts) == 1 {
		log.Fatal("You provide only the folder name. You must provide a entity name too.")
	} else if len(folderParts) < 2 {
		log.Fatal("You must provide a folder name for the bounded context.")
	}

	entity := folderParts[len(folderParts)-1]
	if entity == "" {
		log.Fatal("You must provide a entity name for the bounded context.")
	}

	url, ok := strings.CutSuffix(folder, entity)
	if !ok {
		log.Fatal("Error to parse folder and entity names.")
	}

	folder = strings.Trim(url, string(os.PathSeparator))
	return folder, entity
}
