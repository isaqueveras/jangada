// Package sail provides commands to create layers for a bounded context.
package sail

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	cli "github.com/isaqueveras/jangada/internal"
	templateSail "github.com/isaqueveras/jangada/internal/sail/template"
)

// SailTransport defines the Sail transport structure.
type SailTransport struct {
	pathDir, folder, entity, module, layer string
}

// Execute is the handler for the 'sail transport' command.
func Execute(cmd *cobra.Command, args []string) {
	cli.SetFlagTransportLayer(cmd.Flag("layer").Value.String())
	cli.SetTransportFlagMethodName(cmd.Flag("name").Value.String())
	if cmd.Flag("name").Value.String() == "" {
		cli.SetTransportFlagCreateController(cmd.Flag("controller").Value.String() == "true")
	}

	folder, entity := getFolderAndEntityToTransport(args...)
	cfg := cli.GetConfig()
	st := &SailTransport{
		folder:  folder,
		entity:  entity,
		pathDir: cfg.DirectoryPath,
		module:  cfg.ModuleName,
		layer:   cfg.TransportInfo.FlagTransportLayer,
	}

	mapperCreateLayerTransport[cfg.TransportInfo.FlagTransportLayer](st)
}

type webTransportTemplateData struct {
	Layer, Folder, Entity, Module, Method string
}

// createTransport generates the transport layer structure.
func createTransport(st *SailTransport) {
	var (
		cfg  = cli.GetConfig()
		log  = color.New()
		err  error
		data = &webTransportTemplateData{
			Folder: st.folder,
			Entity: strings.ToLower(st.entity),
			Module: cli.GetModuleName(),
			Layer:  st.layer,
			Method: cfg.TransportInfo.FlagMethodName,
		}
	)

	if !cfg.TransportInfo.FlagCreateController && cfg.TransportInfo.FlagMethodName == "" {
		log.Add(color.Reset, color.FgHiRed, color.Bold).Println("You must use --controller or --method flag.")
		return
	}

	log.Add(color.Bold, color.FgHiGreen).
		Print("Creating transport layer structure...\n")

	switch {
	case cfg.TransportInfo.FlagCreateController && st.layer == RestTransportLayer:
		err = createFileTransport(data, transportTemplateRest)
	case cfg.TransportInfo.FlagCreateController && st.layer == WebTransportLayer:
		err = errors.New("transport layer (web) not implemented")
	case !cfg.TransportInfo.FlagCreateController && cfg.TransportInfo.FlagMethodName != "":
		err = createFileTransport(data, transportTemplateRest)
	}

	if err != nil {
		log.Add(color.Reset, color.FgHiRed).Println("Error: " + err.Error())
		return
	}

	color.New().Add(color.Reset, color.Bold, color.FgHiGreen).
		Print("Transport layer structure created successfully!\n")
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
				log.Add(color.Reset, color.FgHiMagenta, color.Bold).Print("- [exist]\t")
				log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
				continue
			}

			log.Add(color.Reset, color.FgHiYellow, color.Bold).Print("- [update]\t")
			log.Add(color.Reset, color.FgHiWhite).Printf("%s\n", pathFile)
			return updateFile(pathFile, in, data)
		}

		if err = createFile(pathFile, in.Content, data); err != nil {
			return err
		}

		log.Add(color.Reset, color.FgHiGreen, color.Bold).Print("- [created]\t")
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
	if strings.Contains(in.Path, "controller.go") && data.Method != "" {
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}

		updated := (string(content) + "\n" + templateSail.ControllerMethod)
		if err = createTemplateParser(path, updated, file, data); err != nil {
			return err
		}

		return file.Close()
	}

	return nil
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
