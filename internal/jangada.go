package cli

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/mod/modfile"
)

var cfg *Jangada

// GetConfig ....
func GetConfig() *Jangada {
	return cfg
}

// Jangada ...
type Jangada struct {
	AppName     string
	ModuleName  string
	DefaultHost string
	Database    string

	dirBase       string
	DirectoryPath string

	TransportInfo TransportInfo
}

type TransportInfo struct {
	FlagTransportLayer   string
	FlagCreateController bool
	FlagCreateMethod     string
}

// Init create a new instance
func Init(dirBase string) {
	cfg = &Jangada{
		dirBase:     dirBase,
		AppName:     "my-app",
		ModuleName:  "my-app",
		DefaultHost: "localhost:8080",
		Database:    "postgres",
	}
}

func SetAppName(name string) {
	if name == "" {
		return
	}
	cfg.AppName = name
}

// SetDatabase ...
func SetDatabase(db string) {
	if db == "" {
		return
	}
	cfg.Database = db
}

// SetFullDirectoryPath ...
func SetFullDirectoryPath() {
	cfg.DirectoryPath = cfg.dirBase + "/" + strings.ToLower(cfg.AppName)
}

// SetModuleName ...
func SetModuleName(module string) {
	if module == "" {
		cfg.ModuleName = cfg.AppName
		return
	}
	cfg.ModuleName = module
}

// SetDefaultHost ...
func SetDefaultHost(host string) {
	if host == "" {
		return
	}
	cfg.DefaultHost = host
}

func GetDirBase() string {
	return cfg.dirBase
}

// GetModuleName returns the module name
func GetModuleName() string {
	dirGoMod := GetDirBase() + "/go.mod"
	if _, err := os.Stat(dirGoMod); os.IsNotExist(err) {
		return cfg.ModuleName
	}

	data, err := os.ReadFile(dirGoMod)
	if err != nil {
		panic(err)
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	return modFile.Module.Mod.Path
}

// CreateFile ...
func CreateFile(path, tmpl string) {
	log := color.New()

	dir := fmt.Sprintf("%s/%s", cfg.DirectoryPath, filepath.Dir(path))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}

	pathFile := fmt.Sprintf("%s/%s", cfg.DirectoryPath, path)
	if _, err := os.Stat(pathFile); !os.IsNotExist(err) {
		log.Add(color.FgHiMagenta, color.Bold).Print("\texist\t")
		log.Add(color.Reset).Printf("%s\n", path)
		return
	}

	file, err := os.Create(pathFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var t *template.Template
	if t, err = template.New(cfg.dirBase).Parse(tmpl); err != nil {
		panic(err)
	}

	if err := t.Execute(file, cfg); err != nil {
		panic(err)
	}

	log.Add(color.FgHiGreen, color.Bold).Print("\tcreate\t")
	log.Add(color.FgHiWhite, color.Reset).Printf("%s\n", path)
}

// SetFlagTransportLayer set the transport layer
func SetFlagTransportLayer(layer string) { cfg.TransportInfo.FlagTransportLayer = layer }

// SetTransportFlagCreateController set the transport layer
func SetTransportFlagCreateController(crud bool) { cfg.TransportInfo.FlagCreateController = crud }

// SetTransportFlagCreateMethod set the transport layer
func SetTransportFlagCreateMethod(method string) { cfg.TransportInfo.FlagCreateMethod = method }
