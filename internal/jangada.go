package cli

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
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
	AppName    string
	ModuleName string
	Database   string

	DefaultHost string
	DefaultPort string

	dirBase       string
	DirectoryPath string

	TransportInfo   TransportInfo
	ApplicationInfo ApplicationInfo
}

type TransportInfo struct {
	FlagTransportLayer string
	FlagMethodName     string
}

type ApplicationInfo struct {
	FlagService string
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

// SetAppName ...
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
func SetModuleName(mod string) {
	if mod == "" {
		cfg.ModuleName = cfg.AppName
		return
	}
	cfg.ModuleName = mod
}

// SetDefaultHost ...
func SetDefaultHost(host string) {
	if host == "" {
		return
	}
	cfg.DefaultHost = host
}

// SetDefaultPort set the default port
func SetDefaultPort(port string) {
	if port == "" {
		return
	}
	cfg.DefaultPort = OnlyNumbers(port)
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

	funcs := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}

	var t *template.Template
	if t, err = template.New(cfg.dirBase).Funcs(funcs).Parse(tmpl); err != nil {
		panic(err)
	}

	if err = t.Execute(file, cfg); err != nil {
		panic(err)
	}

	log.Add(color.FgHiGreen, color.Bold).Print("\tcreate\t")
	log.Add(color.FgHiWhite, color.Reset).Printf("%s\n", path)
}

// SetFlagTransportLayer set the transport layer
func SetFlagTransportLayer(layer string) { cfg.TransportInfo.FlagTransportLayer = layer }

// SetTransportFlagMethodName set the transport layer
func SetTransportFlagMethodName(method string) { cfg.TransportInfo.FlagMethodName = method }

// SetApplicationFlagService set the transport layer
func SetApplicationFlagService(service string) { cfg.ApplicationInfo.FlagService = service }

// OnlyNumbers returns only numbers
func OnlyNumbers(n string) string {
	return strings.Join(regexp.MustCompile(`\d`).FindAllString(n, -1), "")
}
